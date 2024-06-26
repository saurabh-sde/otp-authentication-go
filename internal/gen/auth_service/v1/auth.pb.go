// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: auth_service/v1/auth.proto

package auth_servicev1

import (
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

type SignupWithPhoneNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Mobile string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *SignupWithPhoneNumberRequest) Reset() {
	*x = SignupWithPhoneNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_v1_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupWithPhoneNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupWithPhoneNumberRequest) ProtoMessage() {}

func (x *SignupWithPhoneNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_v1_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupWithPhoneNumberRequest.ProtoReflect.Descriptor instead.
func (*SignupWithPhoneNumberRequest) Descriptor() ([]byte, []int) {
	return file_auth_service_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *SignupWithPhoneNumberRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SignupWithPhoneNumberRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type SignupWithPhoneNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SignupWithPhoneNumberResponse) Reset() {
	*x = SignupWithPhoneNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_v1_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupWithPhoneNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupWithPhoneNumberResponse) ProtoMessage() {}

func (x *SignupWithPhoneNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_v1_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupWithPhoneNumberResponse.ProtoReflect.Descriptor instead.
func (*SignupWithPhoneNumberResponse) Descriptor() ([]byte, []int) {
	return file_auth_service_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *SignupWithPhoneNumberResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type VerifyPhoneNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Code   string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *VerifyPhoneNumberRequest) Reset() {
	*x = VerifyPhoneNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_v1_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyPhoneNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyPhoneNumberRequest) ProtoMessage() {}

func (x *VerifyPhoneNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_v1_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyPhoneNumberRequest.ProtoReflect.Descriptor instead.
func (*VerifyPhoneNumberRequest) Descriptor() ([]byte, []int) {
	return file_auth_service_v1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyPhoneNumberRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *VerifyPhoneNumberRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type VerifyPhoneNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Valid   bool   `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *VerifyPhoneNumberResponse) Reset() {
	*x = VerifyPhoneNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_v1_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyPhoneNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyPhoneNumberResponse) ProtoMessage() {}

func (x *VerifyPhoneNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_v1_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyPhoneNumberResponse.ProtoReflect.Descriptor instead.
func (*VerifyPhoneNumberResponse) Descriptor() ([]byte, []int) {
	return file_auth_service_v1_auth_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyPhoneNumberResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *VerifyPhoneNumberResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

type LoginWithPhoneNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *LoginWithPhoneNumberRequest) Reset() {
	*x = LoginWithPhoneNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_v1_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginWithPhoneNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginWithPhoneNumberRequest) ProtoMessage() {}

func (x *LoginWithPhoneNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_v1_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginWithPhoneNumberRequest.ProtoReflect.Descriptor instead.
func (*LoginWithPhoneNumberRequest) Descriptor() ([]byte, []int) {
	return file_auth_service_v1_auth_proto_rawDescGZIP(), []int{4}
}

func (x *LoginWithPhoneNumberRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type LoginWithPhoneNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *LoginWithPhoneNumberResponse) Reset() {
	*x = LoginWithPhoneNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_v1_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginWithPhoneNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginWithPhoneNumberResponse) ProtoMessage() {}

func (x *LoginWithPhoneNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_v1_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginWithPhoneNumberResponse.ProtoReflect.Descriptor instead.
func (*LoginWithPhoneNumberResponse) Descriptor() ([]byte, []int) {
	return file_auth_service_v1_auth_proto_rawDescGZIP(), []int{5}
}

func (x *LoginWithPhoneNumberResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *LoginWithPhoneNumberResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ValidatePhoneNumberLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Code   string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *ValidatePhoneNumberLoginRequest) Reset() {
	*x = ValidatePhoneNumberLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_v1_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatePhoneNumberLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatePhoneNumberLoginRequest) ProtoMessage() {}

func (x *ValidatePhoneNumberLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_v1_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatePhoneNumberLoginRequest.ProtoReflect.Descriptor instead.
func (*ValidatePhoneNumberLoginRequest) Descriptor() ([]byte, []int) {
	return file_auth_service_v1_auth_proto_rawDescGZIP(), []int{6}
}

func (x *ValidatePhoneNumberLoginRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *ValidatePhoneNumberLoginRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type ValidatePhoneNumberLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Valid   bool   `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ValidatePhoneNumberLoginResponse) Reset() {
	*x = ValidatePhoneNumberLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_v1_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatePhoneNumberLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatePhoneNumberLoginResponse) ProtoMessage() {}

func (x *ValidatePhoneNumberLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_v1_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatePhoneNumberLoginResponse.ProtoReflect.Descriptor instead.
func (*ValidatePhoneNumberLoginResponse) Descriptor() ([]byte, []int) {
	return file_auth_service_v1_auth_proto_rawDescGZIP(), []int{7}
}

func (x *ValidatePhoneNumberLoginResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ValidatePhoneNumberLoginResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *ValidatePhoneNumberLoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetProfileRequest) Reset() {
	*x = GetProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_v1_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileRequest) ProtoMessage() {}

func (x *GetProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_v1_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileRequest.ProtoReflect.Descriptor instead.
func (*GetProfileRequest) Descriptor() ([]byte, []int) {
	return file_auth_service_v1_auth_proto_rawDescGZIP(), []int{8}
}

func (x *GetProfileRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Mobile string `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Status string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetProfileResponse) Reset() {
	*x = GetProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_v1_auth_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileResponse) ProtoMessage() {}

func (x *GetProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_v1_auth_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileResponse.ProtoReflect.Descriptor instead.
func (*GetProfileResponse) Descriptor() ([]byte, []int) {
	return file_auth_service_v1_auth_proto_rawDescGZIP(), []int{9}
}

func (x *GetProfileResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetProfileResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetProfileResponse) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *GetProfileResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_auth_service_v1_auth_proto protoreflect.FileDescriptor

var file_auth_service_v1_auth_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x4a, 0x0a,
	0x1c, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x57, 0x69, 0x74, 0x68, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22, 0x39, 0x0a, 0x1d, 0x53, 0x69, 0x67,
	0x6e, 0x75, 0x70, 0x57, 0x69, 0x74, 0x68, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x46, 0x0a, 0x18, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x4b, 0x0a, 0x19,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x1b, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x22, 0x51, 0x0a, 0x1c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x4d, 0x0a, 0x1f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x22, 0x68, 0x0a, 0x20, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2c, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x68, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x32, 0xbe, 0x04, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x76, 0x0a, 0x15, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x57, 0x69,
	0x74, 0x68, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2d, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x57, 0x69, 0x74, 0x68, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x75, 0x70, 0x57, 0x69, 0x74, 0x68, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x11,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x29, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x73, 0x0a, 0x14, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x57, 0x69, 0x74, 0x68, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x2c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7f, 0x0a,
	0x18, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x30, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x22, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xd3, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x41,
	0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x75, 0x72, 0x61, 0x62, 0x68, 0x2d, 0x73,
	0x64, 0x65, 0x2f, 0x6f, 0x74, 0x70, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x0e, 0x41, 0x75, 0x74,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x41, 0x75,
	0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a, 0x41,
	0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x41, 0x75, 0x74, 0x68,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_auth_service_v1_auth_proto_rawDescOnce sync.Once
	file_auth_service_v1_auth_proto_rawDescData = file_auth_service_v1_auth_proto_rawDesc
)

func file_auth_service_v1_auth_proto_rawDescGZIP() []byte {
	file_auth_service_v1_auth_proto_rawDescOnce.Do(func() {
		file_auth_service_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_service_v1_auth_proto_rawDescData)
	})
	return file_auth_service_v1_auth_proto_rawDescData
}

var file_auth_service_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_auth_service_v1_auth_proto_goTypes = []interface{}{
	(*SignupWithPhoneNumberRequest)(nil),     // 0: auth_service.v1.SignupWithPhoneNumberRequest
	(*SignupWithPhoneNumberResponse)(nil),    // 1: auth_service.v1.SignupWithPhoneNumberResponse
	(*VerifyPhoneNumberRequest)(nil),         // 2: auth_service.v1.VerifyPhoneNumberRequest
	(*VerifyPhoneNumberResponse)(nil),        // 3: auth_service.v1.VerifyPhoneNumberResponse
	(*LoginWithPhoneNumberRequest)(nil),      // 4: auth_service.v1.LoginWithPhoneNumberRequest
	(*LoginWithPhoneNumberResponse)(nil),     // 5: auth_service.v1.LoginWithPhoneNumberResponse
	(*ValidatePhoneNumberLoginRequest)(nil),  // 6: auth_service.v1.ValidatePhoneNumberLoginRequest
	(*ValidatePhoneNumberLoginResponse)(nil), // 7: auth_service.v1.ValidatePhoneNumberLoginResponse
	(*GetProfileRequest)(nil),                // 8: auth_service.v1.GetProfileRequest
	(*GetProfileResponse)(nil),               // 9: auth_service.v1.GetProfileResponse
}
var file_auth_service_v1_auth_proto_depIdxs = []int32{
	0, // 0: auth_service.v1.AuthService.SignupWithPhoneNumber:input_type -> auth_service.v1.SignupWithPhoneNumberRequest
	2, // 1: auth_service.v1.AuthService.VerifyPhoneNumber:input_type -> auth_service.v1.VerifyPhoneNumberRequest
	4, // 2: auth_service.v1.AuthService.LoginWithPhoneNumber:input_type -> auth_service.v1.LoginWithPhoneNumberRequest
	6, // 3: auth_service.v1.AuthService.ValidatePhoneNumberLogin:input_type -> auth_service.v1.ValidatePhoneNumberLoginRequest
	8, // 4: auth_service.v1.AuthService.GetProfile:input_type -> auth_service.v1.GetProfileRequest
	1, // 5: auth_service.v1.AuthService.SignupWithPhoneNumber:output_type -> auth_service.v1.SignupWithPhoneNumberResponse
	3, // 6: auth_service.v1.AuthService.VerifyPhoneNumber:output_type -> auth_service.v1.VerifyPhoneNumberResponse
	5, // 7: auth_service.v1.AuthService.LoginWithPhoneNumber:output_type -> auth_service.v1.LoginWithPhoneNumberResponse
	7, // 8: auth_service.v1.AuthService.ValidatePhoneNumberLogin:output_type -> auth_service.v1.ValidatePhoneNumberLoginResponse
	9, // 9: auth_service.v1.AuthService.GetProfile:output_type -> auth_service.v1.GetProfileResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_service_v1_auth_proto_init() }
func file_auth_service_v1_auth_proto_init() {
	if File_auth_service_v1_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_service_v1_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupWithPhoneNumberRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_service_v1_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupWithPhoneNumberResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_service_v1_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyPhoneNumberRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_service_v1_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyPhoneNumberResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_service_v1_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginWithPhoneNumberRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_service_v1_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginWithPhoneNumberResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_service_v1_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatePhoneNumberLoginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_service_v1_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatePhoneNumberLoginResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_service_v1_auth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_service_v1_auth_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_service_v1_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_service_v1_auth_proto_goTypes,
		DependencyIndexes: file_auth_service_v1_auth_proto_depIdxs,
		MessageInfos:      file_auth_service_v1_auth_proto_msgTypes,
	}.Build()
	File_auth_service_v1_auth_proto = out.File
	file_auth_service_v1_auth_proto_rawDesc = nil
	file_auth_service_v1_auth_proto_goTypes = nil
	file_auth_service_v1_auth_proto_depIdxs = nil
}
