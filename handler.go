package humanlog

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/kr/logfmt"
)

// Handler can recognize it's log lines, parse them and prettify them.
type Handler interface {
	CanHandle(line []byte) bool
	Prettify(skipUnchanged bool) []byte
	logfmt.Handler
}

var DefaultOptions = &HandlerOptions{
	SortLongest:    true,
	SkipUnchanged:  true,
	Truncates:      true,
	ColorFlag:      ColorModeAuto,
	LightBg:        false,
	TruncateLength: 15,
	TimeFormat:     time.Stamp,

	TimeFields:    []string{"time", "ts", "@timestamp", "timestamp"},
	MessageFields: []string{"message", "msg"},
	LevelFields:   []string{"level", "lvl", "loglevel", "severity"},

	KeyColor:              color.New(color.FgGreen),
	ValColor:              color.New(color.FgHiWhite),
	TimeLightBgColor:      color.New(color.FgBlack),
	TimeDarkBgColor:       color.New(color.FgWhite),
	MsgLightBgColor:       color.New(color.FgBlack),
	MsgAbsentLightBgColor: color.New(color.FgHiBlack),
	MsgDarkBgColor:        color.New(color.FgHiWhite),
	MsgAbsentDarkBgColor:  color.New(color.FgWhite),
	DebugLevelColor:       color.New(color.FgMagenta),
	InfoLevelColor:        color.New(color.FgCyan),
	WarnLevelColor:        color.New(color.FgYellow),
	ErrorLevelColor:       color.New(color.FgRed),
	PanicLevelColor:       color.New(color.BgRed),
	FatalLevelColor:       color.New(color.BgHiRed, color.FgHiWhite),
	UnknownLevelColor:     color.New(color.FgMagenta),
}

type ColorMode int

const (
	ColorModeOff ColorMode = iota
	ColorModeOn
	ColorModeAuto
)

func GrokColorMode(colorMode string) (ColorMode, error) {
	switch strings.ToLower(colorMode) {
	case "on", "always", "force", "true", "yes", "1":
		return ColorModeOn, nil
	case "off", "never", "false", "no", "0":
		return ColorModeOff, nil
	case "auto", "tty", "maybe", "":
		return ColorModeAuto, nil
	default:
		return ColorModeAuto, fmt.Errorf("'%s' is not a color mode (try 'on', 'off' or 'auto')", colorMode)
	}
}

func (colorMode ColorMode) Apply() {
	switch colorMode {
	case ColorModeOff:
		color.NoColor = true
	case ColorModeOn:
		color.NoColor = false
	default:
		// 'Auto' default is applied as a global variable initializer function, so nothing
		// to do here.
	}
}

type HandlerOptions struct {
	Skip map[string]struct{}
	Keep map[string]struct{}

	TimeFields    []string
	MessageFields []string
	LevelFields   []string

	SortLongest    bool
	SkipUnchanged  bool
	Truncates      bool
	LightBg        bool
	ColorFlag      ColorMode
	TruncateLength int
	TimeFormat     string

	KeyColor              *color.Color
	ValColor              *color.Color
	TimeLightBgColor      *color.Color
	TimeDarkBgColor       *color.Color
	MsgLightBgColor       *color.Color
	MsgAbsentLightBgColor *color.Color
	MsgDarkBgColor        *color.Color
	MsgAbsentDarkBgColor  *color.Color
	DebugLevelColor       *color.Color
	InfoLevelColor        *color.Color
	WarnLevelColor        *color.Color
	ErrorLevelColor       *color.Color
	PanicLevelColor       *color.Color
	FatalLevelColor       *color.Color
	UnknownLevelColor     *color.Color
}

func (h *HandlerOptions) shouldShowKey(key string) bool {
	if len(h.Keep) != 0 {
		if _, keep := h.Keep[key]; keep {
			return true
		}
	}
	if len(h.Skip) != 0 {
		if _, skip := h.Skip[key]; skip {
			return false
		}
	}
	return true
}

func (h *HandlerOptions) shouldShowUnchanged(key string) bool {
	if len(h.Keep) != 0 {
		if _, keep := h.Keep[key]; keep {
			return true
		}
	}
	return false
}

func (h *HandlerOptions) SetSkip(skip []string) {
	if h.Skip == nil {
		h.Skip = make(map[string]struct{})
	}
	for _, key := range skip {
		h.Skip[key] = struct{}{}
	}
}

func (h *HandlerOptions) SetKeep(keep []string) {
	if h.Keep == nil {
		h.Keep = make(map[string]struct{})
	}
	for _, key := range keep {
		h.Keep[key] = struct{}{}
	}
}
