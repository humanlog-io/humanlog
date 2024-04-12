// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: svc/organization/v1/service.proto

package organizationv1

import (
	v1 "github.com/humanlogio/api/go/types/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrganizationId int64  `protobuf:"varint,1,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	AccountName    string `protobuf:"bytes,2,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	mi := &file_svc_organization_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_organization_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_svc_organization_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccountRequest) GetOrganizationId() int64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *CreateAccountRequest) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

type CreateAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *v1.Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *CreateAccountResponse) Reset() {
	*x = CreateAccountResponse{}
	mi := &file_svc_organization_v1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountResponse) ProtoMessage() {}

func (x *CreateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_organization_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountResponse.ProtoReflect.Descriptor instead.
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return file_svc_organization_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccountResponse) GetAccount() *v1.Account {
	if x != nil {
		return x.Account
	}
	return nil
}

type ListAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cursor         *v1.Cursor `protobuf:"bytes,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Limit          int32      `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	OrganizationId int64      `protobuf:"varint,3,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
}

func (x *ListAccountRequest) Reset() {
	*x = ListAccountRequest{}
	mi := &file_svc_organization_v1_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountRequest) ProtoMessage() {}

func (x *ListAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_organization_v1_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountRequest.ProtoReflect.Descriptor instead.
func (*ListAccountRequest) Descriptor() ([]byte, []int) {
	return file_svc_organization_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListAccountRequest) GetCursor() *v1.Cursor {
	if x != nil {
		return x.Cursor
	}
	return nil
}

func (x *ListAccountRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListAccountRequest) GetOrganizationId() int64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

type ListAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Next  *v1.Cursor                      `protobuf:"bytes,1,opt,name=next,proto3" json:"next,omitempty"`
	Items []*ListAccountResponse_ListItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListAccountResponse) Reset() {
	*x = ListAccountResponse{}
	mi := &file_svc_organization_v1_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountResponse) ProtoMessage() {}

func (x *ListAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_organization_v1_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountResponse.ProtoReflect.Descriptor instead.
func (*ListAccountResponse) Descriptor() ([]byte, []int) {
	return file_svc_organization_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListAccountResponse) GetNext() *v1.Cursor {
	if x != nil {
		return x.Next
	}
	return nil
}

func (x *ListAccountResponse) GetItems() []*ListAccountResponse_ListItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type ListUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cursor         *v1.Cursor `protobuf:"bytes,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Limit          int32      `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	OrganizationId int64      `protobuf:"varint,3,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
}

func (x *ListUserRequest) Reset() {
	*x = ListUserRequest{}
	mi := &file_svc_organization_v1_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserRequest) ProtoMessage() {}

func (x *ListUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_organization_v1_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserRequest.ProtoReflect.Descriptor instead.
func (*ListUserRequest) Descriptor() ([]byte, []int) {
	return file_svc_organization_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListUserRequest) GetCursor() *v1.Cursor {
	if x != nil {
		return x.Cursor
	}
	return nil
}

func (x *ListUserRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListUserRequest) GetOrganizationId() int64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

type ListUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Next  *v1.Cursor                   `protobuf:"bytes,1,opt,name=next,proto3" json:"next,omitempty"`
	Items []*ListUserResponse_ListItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListUserResponse) Reset() {
	*x = ListUserResponse{}
	mi := &file_svc_organization_v1_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserResponse) ProtoMessage() {}

func (x *ListUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_organization_v1_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserResponse.ProtoReflect.Descriptor instead.
func (*ListUserResponse) Descriptor() ([]byte, []int) {
	return file_svc_organization_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListUserResponse) GetNext() *v1.Cursor {
	if x != nil {
		return x.Next
	}
	return nil
}

func (x *ListUserResponse) GetItems() []*ListUserResponse_ListItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type InviteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrganizationId int64 `protobuf:"varint,1,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	UserEmail      int64 `protobuf:"varint,2,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
}

func (x *InviteUserRequest) Reset() {
	*x = InviteUserRequest{}
	mi := &file_svc_organization_v1_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InviteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteUserRequest) ProtoMessage() {}

func (x *InviteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_organization_v1_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteUserRequest.ProtoReflect.Descriptor instead.
func (*InviteUserRequest) Descriptor() ([]byte, []int) {
	return file_svc_organization_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *InviteUserRequest) GetOrganizationId() int64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *InviteUserRequest) GetUserEmail() int64 {
	if x != nil {
		return x.UserEmail
	}
	return 0
}

type InviteUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InviteUserResponse) Reset() {
	*x = InviteUserResponse{}
	mi := &file_svc_organization_v1_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InviteUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteUserResponse) ProtoMessage() {}

func (x *InviteUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_organization_v1_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteUserResponse.ProtoReflect.Descriptor instead.
func (*InviteUserResponse) Descriptor() ([]byte, []int) {
	return file_svc_organization_v1_service_proto_rawDescGZIP(), []int{7}
}

type RevokeUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrganizationId int64 `protobuf:"varint,1,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	UserId         int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *RevokeUserRequest) Reset() {
	*x = RevokeUserRequest{}
	mi := &file_svc_organization_v1_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RevokeUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeUserRequest) ProtoMessage() {}

func (x *RevokeUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_organization_v1_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeUserRequest.ProtoReflect.Descriptor instead.
func (*RevokeUserRequest) Descriptor() ([]byte, []int) {
	return file_svc_organization_v1_service_proto_rawDescGZIP(), []int{8}
}

func (x *RevokeUserRequest) GetOrganizationId() int64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *RevokeUserRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type RevokeUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RevokeUserResponse) Reset() {
	*x = RevokeUserResponse{}
	mi := &file_svc_organization_v1_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RevokeUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeUserResponse) ProtoMessage() {}

func (x *RevokeUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_organization_v1_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeUserResponse.ProtoReflect.Descriptor instead.
func (*RevokeUserResponse) Descriptor() ([]byte, []int) {
	return file_svc_organization_v1_service_proto_rawDescGZIP(), []int{9}
}

type ListAccountResponse_ListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *v1.Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *ListAccountResponse_ListItem) Reset() {
	*x = ListAccountResponse_ListItem{}
	mi := &file_svc_organization_v1_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAccountResponse_ListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountResponse_ListItem) ProtoMessage() {}

func (x *ListAccountResponse_ListItem) ProtoReflect() protoreflect.Message {
	mi := &file_svc_organization_v1_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountResponse_ListItem.ProtoReflect.Descriptor instead.
func (*ListAccountResponse_ListItem) Descriptor() ([]byte, []int) {
	return file_svc_organization_v1_service_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ListAccountResponse_ListItem) GetAccount() *v1.Account {
	if x != nil {
		return x.Account
	}
	return nil
}

type ListUserResponse_ListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *v1.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ListUserResponse_ListItem) Reset() {
	*x = ListUserResponse_ListItem{}
	mi := &file_svc_organization_v1_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUserResponse_ListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserResponse_ListItem) ProtoMessage() {}

func (x *ListUserResponse_ListItem) ProtoReflect() protoreflect.Message {
	mi := &file_svc_organization_v1_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserResponse_ListItem.ProtoReflect.Descriptor instead.
func (*ListUserResponse_ListItem) Descriptor() ([]byte, []int) {
	return file_svc_organization_v1_service_proto_rawDescGZIP(), []int{5, 0}
}

func (x *ListUserResponse_ListItem) GetUser() *v1.User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_svc_organization_v1_service_proto protoreflect.FileDescriptor

var file_svc_organization_v1_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x76, 0x63, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x76, 0x63, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x16, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x72, 0x73, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x44, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x7d, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x63, 0x75, 0x72,
	0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x06, 0x63, 0x75, 0x72,
	0x73, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0xbd, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x65,
	0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74,
	0x12, 0x47, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x37, 0x0a, 0x08, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x2b, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x7a, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xae,
	0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72,
	0x73, 0x6f, 0x72, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x12, 0x44, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a,
	0x2e, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x22, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x5b, 0x0a, 0x11, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x14, 0x0a, 0x12,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x55, 0x0a, 0x11, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x52, 0x65, 0x76,
	0x6f, 0x6b, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0x80, 0x04, 0x0a, 0x13, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x62, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x76, 0x63, 0x2e,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x24, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x5f, 0x0a, 0x0a, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x26,
	0x2e, 0x73, 0x76, 0x63, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5f, 0x0a, 0x0a, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x26, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x76, 0x6f, 0x6b, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0xd6, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x6d, 0x61, 0x6e,
	0x6c, 0x6f, 0x67, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x76, 0x63,
	0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x3b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x53, 0x4f, 0x58, 0xaa, 0x02, 0x13, 0x53, 0x76, 0x63, 0x2e, 0x4f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x13, 0x53, 0x76,
	0x63, 0x5c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x1f, 0x53, 0x76, 0x63, 0x5c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x53, 0x76, 0x63, 0x3a, 0x3a, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_svc_organization_v1_service_proto_rawDescOnce sync.Once
	file_svc_organization_v1_service_proto_rawDescData = file_svc_organization_v1_service_proto_rawDesc
)

func file_svc_organization_v1_service_proto_rawDescGZIP() []byte {
	file_svc_organization_v1_service_proto_rawDescOnce.Do(func() {
		file_svc_organization_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_organization_v1_service_proto_rawDescData)
	})
	return file_svc_organization_v1_service_proto_rawDescData
}

var file_svc_organization_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_svc_organization_v1_service_proto_goTypes = []any{
	(*CreateAccountRequest)(nil),         // 0: svc.organization.v1.CreateAccountRequest
	(*CreateAccountResponse)(nil),        // 1: svc.organization.v1.CreateAccountResponse
	(*ListAccountRequest)(nil),           // 2: svc.organization.v1.ListAccountRequest
	(*ListAccountResponse)(nil),          // 3: svc.organization.v1.ListAccountResponse
	(*ListUserRequest)(nil),              // 4: svc.organization.v1.ListUserRequest
	(*ListUserResponse)(nil),             // 5: svc.organization.v1.ListUserResponse
	(*InviteUserRequest)(nil),            // 6: svc.organization.v1.InviteUserRequest
	(*InviteUserResponse)(nil),           // 7: svc.organization.v1.InviteUserResponse
	(*RevokeUserRequest)(nil),            // 8: svc.organization.v1.RevokeUserRequest
	(*RevokeUserResponse)(nil),           // 9: svc.organization.v1.RevokeUserResponse
	(*ListAccountResponse_ListItem)(nil), // 10: svc.organization.v1.ListAccountResponse.ListItem
	(*ListUserResponse_ListItem)(nil),    // 11: svc.organization.v1.ListUserResponse.ListItem
	(*v1.Account)(nil),                   // 12: types.v1.Account
	(*v1.Cursor)(nil),                    // 13: types.v1.Cursor
	(*v1.User)(nil),                      // 14: types.v1.User
}
var file_svc_organization_v1_service_proto_depIdxs = []int32{
	12, // 0: svc.organization.v1.CreateAccountResponse.account:type_name -> types.v1.Account
	13, // 1: svc.organization.v1.ListAccountRequest.cursor:type_name -> types.v1.Cursor
	13, // 2: svc.organization.v1.ListAccountResponse.next:type_name -> types.v1.Cursor
	10, // 3: svc.organization.v1.ListAccountResponse.items:type_name -> svc.organization.v1.ListAccountResponse.ListItem
	13, // 4: svc.organization.v1.ListUserRequest.cursor:type_name -> types.v1.Cursor
	13, // 5: svc.organization.v1.ListUserResponse.next:type_name -> types.v1.Cursor
	11, // 6: svc.organization.v1.ListUserResponse.items:type_name -> svc.organization.v1.ListUserResponse.ListItem
	12, // 7: svc.organization.v1.ListAccountResponse.ListItem.account:type_name -> types.v1.Account
	14, // 8: svc.organization.v1.ListUserResponse.ListItem.user:type_name -> types.v1.User
	0,  // 9: svc.organization.v1.OrganizationService.CreateAccount:input_type -> svc.organization.v1.CreateAccountRequest
	2,  // 10: svc.organization.v1.OrganizationService.ListAccount:input_type -> svc.organization.v1.ListAccountRequest
	4,  // 11: svc.organization.v1.OrganizationService.ListUser:input_type -> svc.organization.v1.ListUserRequest
	6,  // 12: svc.organization.v1.OrganizationService.InviteUser:input_type -> svc.organization.v1.InviteUserRequest
	8,  // 13: svc.organization.v1.OrganizationService.RevokeUser:input_type -> svc.organization.v1.RevokeUserRequest
	1,  // 14: svc.organization.v1.OrganizationService.CreateAccount:output_type -> svc.organization.v1.CreateAccountResponse
	3,  // 15: svc.organization.v1.OrganizationService.ListAccount:output_type -> svc.organization.v1.ListAccountResponse
	5,  // 16: svc.organization.v1.OrganizationService.ListUser:output_type -> svc.organization.v1.ListUserResponse
	7,  // 17: svc.organization.v1.OrganizationService.InviteUser:output_type -> svc.organization.v1.InviteUserResponse
	9,  // 18: svc.organization.v1.OrganizationService.RevokeUser:output_type -> svc.organization.v1.RevokeUserResponse
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_svc_organization_v1_service_proto_init() }
func file_svc_organization_v1_service_proto_init() {
	if File_svc_organization_v1_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_svc_organization_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_organization_v1_service_proto_goTypes,
		DependencyIndexes: file_svc_organization_v1_service_proto_depIdxs,
		MessageInfos:      file_svc_organization_v1_service_proto_msgTypes,
	}.Build()
	File_svc_organization_v1_service_proto = out.File
	file_svc_organization_v1_service_proto_rawDesc = nil
	file_svc_organization_v1_service_proto_goTypes = nil
	file_svc_organization_v1_service_proto_depIdxs = nil
}
