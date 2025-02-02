package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/huh"
	"github.com/humanlogio/api/go/svc/auth/v1/authv1connect"
	"github.com/humanlogio/humanlog/internal/pkg/config"
	"github.com/humanlogio/humanlog/internal/pkg/state"
	"github.com/humanlogio/humanlog/pkg/auth"
	"github.com/urfave/cli"
)

const onboardingCmdName = "onboarding"

func onboardingCmd(
	getCtx func(cctx *cli.Context) context.Context,
	getLogger func(cctx *cli.Context) *slog.Logger,
	getCfg func(cctx *cli.Context) *config.Config,
	getState func(cctx *cli.Context) *state.State,
	getTokenSource func(cctx *cli.Context) *auth.UserRefreshableTokenSource,
	getAPIUrl func(cctx *cli.Context) string,
	getBaseSiteURL func(cctx *cli.Context) string,
	getHTTPClient func(*cli.Context, string) *http.Client,
) cli.Command {

	runsAsService := func(cfg *config.Config) bool {
		if cfg == nil {
			return false
		}
		if cfg.ExperimentalFeatures == nil {
			return false
		}
		if cfg.ExperimentalFeatures.ServeLocalhost != nil {
			return true
		}
		if cfg.ExperimentalFeatures.ShowInSystray != nil {
			return *cfg.ExperimentalFeatures.ShowInSystray
		}
		return false
	}

	ensureServiceEnabled := func(cctx *cli.Context) error {
		_, svc, err := prepareServiceCmd(cctx,
			getCtx,
			getLogger,
			getCfg,
			getState,
			getTokenSource,
			getAPIUrl,
			getBaseSiteURL,
			getHTTPClient,
		)
		if err != nil {
			return fmt.Errorf("failed to get humanlog service details: %v", err)
		}
		if err := svc.Stop(); err != nil {
			logwarn("failed to stop service (was it running?): %v", err)
		}
		if err := svc.Uninstall(); err != nil {
			logwarn("failed to uninstall service (was it installed?): %v", err)
		}
		if err := svc.Install(); err != nil {
			return fmt.Errorf("can't install service: %v", err)
		}
		if err := svc.Start(); err != nil {
			return fmt.Errorf("can't start service: %v", err)
		}
		return nil
	}

	var (
		forceNonInteractiveFlag = cli.BoolFlag{Name: "force-non-interactive"}
	)

	return cli.Command{
		Name:   onboardingCmdName,
		Usage:  "Onboarding humanlog after installs or updates",
		Hidden: true,
		Flags:  []cli.Flag{forceNonInteractiveFlag},
		Action: func(cctx *cli.Context) error {
			ctx := getCtx(cctx)
			cfg := getCfg(cctx)
			state := getState(cctx)
			ll := getLogger(cctx)
			tokenSource := getTokenSource(cctx)
			apiURL := getAPIUrl(cctx)
			httpClient := getHTTPClient(cctx, apiURL)
			// clOpts := connect.WithClientOptions(connect.WithInterceptors(auth.Interceptors(ll, tokenSource)...))
			// userSvc := userv1connect.NewUserServiceClient(httpClient, apiURL, clOpts)

			user, err := checkUserLoggedIn(ctx, ll, httpClient, apiURL, tokenSource)
			if err != nil {
				logwarn("unable to check if you're logged in: %v", err)
			}

			defer func() {
				if !runsAsService(cfg) {
					return
				}
				if err := ensureServiceEnabled(cctx); err != nil {
					logerror("unable to configure humanlog service: %v", err)
				} else {
					loginfo("humanlog service is configured")
				}
			}()

			if !isTerminal(os.Stdin) || cctx.Bool(forceNonInteractiveFlag.Name) {

				in := `# humanlog updates

Hey there!

Thanks for installing this version of humanlog. If this is your first time around, try this out:

` + "```bash" + `
humanlog onboarding
` + "```" + `

This will help you get started and learn everything that humanlog has to offer.

Bye! <3`

				out, err := glamour.Render(in, "dark")
				if err != nil {
					return err
				}
				fmt.Print(out)

				return nil
			}

			promptSignup := state.LastPromptedToSignupAt == nil && (user == nil)
			promptQueryEngine := state.LastPromptedToEnableLocalhostAt == nil && (cfg.ExperimentalFeatures == nil || cfg.ExperimentalFeatures.ServeLocalhost == nil)

			var (
				wantsSignup      = promptSignup && true
				wantsQueryEngine = promptQueryEngine && true
			)

			var fields []huh.Field
			// fields = append(fields, huh.NewNote().Title("Welcome to humanlog. New features are available and more are coming soon!"))
			if promptQueryEngine {
				fields = append(fields,
					huh.NewConfirm().
						Title("Humanlog now includes a log query engine, right here in your pocket.\n\n"+
							"You can use it to query your logs, plot graphs and do general log observability stuff. All on your machine!\n\n"+
							"To enable this feature, humanlog needs to run a background service.").
						Description("Do you want to enable the log query engine?").
						Affirmative("Yes!").Negative("No.").
						Value(&wantsQueryEngine),
				)
				state.LastPromptedToEnableLocalhostAt = ptr(time.Now())
				wantsSignup = user == nil
			}
			if promptSignup && !promptQueryEngine {
				fields = append(fields,
					huh.NewConfirm().
						Title("New features are coming soon. Sign in to learn more.").
						Description("Sign up to learn about upcoming releases?").
						Affirmative("Yes!").Negative("No").Value(&wantsSignup),
				)
				state.LastPromptedToSignupAt = ptr(time.Now())
			}
			if len(fields) > 0 {
				err := huh.NewForm(huh.NewGroup(fields...)).WithTheme(huhTheme).Run()
				if err != nil {
					return err
				}
				if err := state.WriteBack(); err != nil {
					logwarn("failed to record your answer: %v", err)
				}
			}

			if wantsSignup {
				loginfo("awesome, thanks for your interest!")

				authClient := authv1connect.NewAuthServiceClient(httpClient, apiURL)
				_, err := performLoginFlow(ctx, state, authClient, tokenSource)
				if err != nil {
					logerror("failed to sign up or sign in: %v", err)
				}
			}

			if wantsQueryEngine {
				dbpath := "~/.humanlog/data/db.humanlog"
				if err := os.MkdirAll(filepath.Dir(dbpath), 0755); err != nil {
					return fmt.Errorf("creating directory for DB: %v", err)
				}
				if cfg.ExperimentalFeatures == nil {
					cfg.ExperimentalFeatures = &config.Features{}
				}
				cfg.ExperimentalFeatures.ServeLocalhost = &config.ServeLocalhost{
					Port:   32764,
					Engine: "advanced",
					Cfg: map[string]interface{}{
						"path": dbpath,
					},
				}
				cfg.ExperimentalFeatures.ShowInSystray = ptr(true)
				if err := cfg.WriteBack(); err != nil {
					logerror("failed to update config file: %v", err)
				}
			}

			loginfo("keep an eye on `https://humanlog.io` for more updates!")

			return nil
		},
	}
}

func pjson(v any) string {
	out, err := json.MarshalIndent(v, "", "   ")
	if err != nil {
		panic(err)
	}
	return string(out)
}
