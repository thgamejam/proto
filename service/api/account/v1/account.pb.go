// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: api/account/v1/account.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// CreateEMailAccountReq 创建邮箱账号请求
type CreateEMailAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 密文密码 len:64-256
	Ciphertext string `protobuf:"bytes,1,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	// 公钥摘要 len:16
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// 电子邮箱地址
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CreateEMailAccountReq) Reset() {
	*x = CreateEMailAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_v1_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEMailAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEMailAccountReq) ProtoMessage() {}

func (x *CreateEMailAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_v1_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEMailAccountReq.ProtoReflect.Descriptor instead.
func (*CreateEMailAccountReq) Descriptor() ([]byte, []int) {
	return file_api_account_v1_account_proto_rawDescGZIP(), []int{0}
}

func (x *CreateEMailAccountReq) GetCiphertext() string {
	if x != nil {
		return x.Ciphertext
	}
	return ""
}

func (x *CreateEMailAccountReq) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *CreateEMailAccountReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// CreateAccountReply 创建账号回复
type CreateEMailAccountReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateEMailAccountReply) Reset() {
	*x = CreateEMailAccountReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_v1_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEMailAccountReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEMailAccountReply) ProtoMessage() {}

func (x *CreateEMailAccountReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_v1_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEMailAccountReply.ProtoReflect.Descriptor instead.
func (*CreateEMailAccountReply) Descriptor() ([]byte, []int) {
	return file_api_account_v1_account_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEMailAccountReply) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// GetAccountReq 使用ID获取账号请求
type GetAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAccountReq) Reset() {
	*x = GetAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_v1_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountReq) ProtoMessage() {}

func (x *GetAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_v1_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountReq.ProtoReflect.Descriptor instead.
func (*GetAccountReq) Descriptor() ([]byte, []int) {
	return file_api_account_v1_account_proto_rawDescGZIP(), []int{2}
}

func (x *GetAccountReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// GetAccountReply 使用ID获取账号回复
type GetAccountReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid    string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`        // uuid
	Email   string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`      // 电子邮箱地址
	TelCode uint32 `protobuf:"varint,3,opt,name=telCode,proto3" json:"telCode,omitempty"` // 国际区号
	Phone   string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`      // 电话号
	Status  uint32 `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`   // 账号状态
	UserID  uint32 `protobuf:"varint,6,opt,name=userID,proto3" json:"userID,omitempty"`   // 用户表ID
}

func (x *GetAccountReply) Reset() {
	*x = GetAccountReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_v1_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountReply) ProtoMessage() {}

func (x *GetAccountReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_v1_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountReply.ProtoReflect.Descriptor instead.
func (*GetAccountReply) Descriptor() ([]byte, []int) {
	return file_api_account_v1_account_proto_rawDescGZIP(), []int{3}
}

func (x *GetAccountReply) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GetAccountReply) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetAccountReply) GetTelCode() uint32 {
	if x != nil {
		return x.TelCode
	}
	return 0
}

func (x *GetAccountReply) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *GetAccountReply) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetAccountReply) GetUserID() uint32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

// ExistAccountEMailReq 检查邮箱是否存在请求
type ExistAccountEMailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"` // 电子邮箱地址
}

func (x *ExistAccountEMailReq) Reset() {
	*x = ExistAccountEMailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_v1_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistAccountEMailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistAccountEMailReq) ProtoMessage() {}

func (x *ExistAccountEMailReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_v1_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistAccountEMailReq.ProtoReflect.Descriptor instead.
func (*ExistAccountEMailReq) Descriptor() ([]byte, []int) {
	return file_api_account_v1_account_proto_rawDescGZIP(), []int{4}
}

func (x *ExistAccountEMailReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// ExistAccountEMailReply 检查邮箱是否存在回复
type ExistAccountEMailReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"` // true 存在, false 不存在
}

func (x *ExistAccountEMailReply) Reset() {
	*x = ExistAccountEMailReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_v1_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistAccountEMailReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistAccountEMailReply) ProtoMessage() {}

func (x *ExistAccountEMailReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_v1_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistAccountEMailReply.ProtoReflect.Descriptor instead.
func (*ExistAccountEMailReply) Descriptor() ([]byte, []int) {
	return file_api_account_v1_account_proto_rawDescGZIP(), []int{5}
}

func (x *ExistAccountEMailReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

// VerifyPasswordReq 验证密码请求
type VerifyPasswordReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户登录标识, 未必是用户名 len:4-32
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// 密文密码 len:64-256
	Ciphertext string `protobuf:"bytes,2,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	// 公钥摘要 len:16
	Hash string `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *VerifyPasswordReq) Reset() {
	*x = VerifyPasswordReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_v1_account_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyPasswordReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyPasswordReq) ProtoMessage() {}

func (x *VerifyPasswordReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_v1_account_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyPasswordReq.ProtoReflect.Descriptor instead.
func (*VerifyPasswordReq) Descriptor() ([]byte, []int) {
	return file_api_account_v1_account_proto_rawDescGZIP(), []int{6}
}

func (x *VerifyPasswordReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *VerifyPasswordReq) GetCiphertext() string {
	if x != nil {
		return x.Ciphertext
	}
	return ""
}

func (x *VerifyPasswordReq) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

// VerifyPasswordReply 验证密码回复
type VerifyPasswordReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *VerifyPasswordReply) Reset() {
	*x = VerifyPasswordReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_v1_account_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyPasswordReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyPasswordReply) ProtoMessage() {}

func (x *VerifyPasswordReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_v1_account_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyPasswordReply.ProtoReflect.Descriptor instead.
func (*VerifyPasswordReply) Descriptor() ([]byte, []int) {
	return file_api_account_v1_account_proto_rawDescGZIP(), []int{7}
}

func (x *VerifyPasswordReply) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// SavePasswordReq 保存密码请求
type SavePasswordReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户id
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 密文密码 len:64-256
	Ciphertext string `protobuf:"bytes,2,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	// 公钥摘要 len:16
	Hash string `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *SavePasswordReq) Reset() {
	*x = SavePasswordReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_v1_account_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavePasswordReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavePasswordReq) ProtoMessage() {}

func (x *SavePasswordReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_v1_account_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavePasswordReq.ProtoReflect.Descriptor instead.
func (*SavePasswordReq) Descriptor() ([]byte, []int) {
	return file_api_account_v1_account_proto_rawDescGZIP(), []int{8}
}

func (x *SavePasswordReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SavePasswordReq) GetCiphertext() string {
	if x != nil {
		return x.Ciphertext
	}
	return ""
}

func (x *SavePasswordReq) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

// SavePasswordReply 保存密码回复
type SavePasswordReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SavePasswordReply) Reset() {
	*x = SavePasswordReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_v1_account_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavePasswordReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavePasswordReply) ProtoMessage() {}

func (x *SavePasswordReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_v1_account_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavePasswordReply.ProtoReflect.Descriptor instead.
func (*SavePasswordReply) Descriptor() ([]byte, []int) {
	return file_api_account_v1_account_proto_rawDescGZIP(), []int{9}
}

// BindUserReq 绑定用户请求
type BindUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserID uint32 `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *BindUserReq) Reset() {
	*x = BindUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_v1_account_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindUserReq) ProtoMessage() {}

func (x *BindUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_v1_account_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindUserReq.ProtoReflect.Descriptor instead.
func (*BindUserReq) Descriptor() ([]byte, []int) {
	return file_api_account_v1_account_proto_rawDescGZIP(), []int{10}
}

func (x *BindUserReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BindUserReq) GetUserID() uint32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

// BindUserReply 绑定用户回复
type BindUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BindUserReply) Reset() {
	*x = BindUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_v1_account_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindUserReply) ProtoMessage() {}

func (x *BindUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_v1_account_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindUserReply.ProtoReflect.Descriptor instead.
func (*BindUserReply) Descriptor() ([]byte, []int) {
	return file_api_account_v1_account_proto_rawDescGZIP(), []int{11}
}

// GetKeyReq 获取公钥请求
type GetKeyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetKeyReq) Reset() {
	*x = GetKeyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_v1_account_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeyReq) ProtoMessage() {}

func (x *GetKeyReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_v1_account_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeyReq.ProtoReflect.Descriptor instead.
func (*GetKeyReq) Descriptor() ([]byte, []int) {
	return file_api_account_v1_account_proto_rawDescGZIP(), []int{12}
}

// GetKeyReply 获取公钥回复
type GetKeyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"` // 密钥hash
	Key  string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`   // 公钥内容
}

func (x *GetKeyReply) Reset() {
	*x = GetKeyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_v1_account_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeyReply) ProtoMessage() {}

func (x *GetKeyReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_v1_account_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeyReply.ProtoReflect.Descriptor instead.
func (*GetKeyReply) Descriptor() ([]byte, []int) {
	return file_api_account_v1_account_proto_rawDescGZIP(), []int{13}
}

func (x *GetKeyReply) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *GetKeyReply) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_api_account_v1_account_proto protoreflect.FileDescriptor

var file_api_account_v1_account_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x4d,
	0x61, 0x69, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2a, 0x0a,
	0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x40, 0x18, 0x80, 0x02, 0x52, 0x0a, 0x63,
	0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x98, 0x01,
	0x10, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1f, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x18, 0x20, 0x60,
	0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x29, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x4d, 0x61, 0x69, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0x9b, 0x01,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x74,
	0x65, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x74, 0x65,
	0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x37, 0x0a, 0x14, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x4d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x18, 0x20, 0x60, 0x01, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x22, 0x28, 0x0a, 0x16, 0x45, 0x78, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x45, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x84,
	0x01, 0x0a, 0x11, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x04, 0x18,
	0x20, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x0a, 0x63,
	0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x40, 0x18, 0x80, 0x02, 0x52, 0x0a, 0x63, 0x69, 0x70,
	0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x98, 0x01, 0x10, 0x52,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x25, 0x0a, 0x13, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x74, 0x0a, 0x0f,
	0x53, 0x61, 0x76, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12,
	0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x2a, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x0a, 0x63, 0x69, 0x70, 0x68,
	0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42,
	0x07, 0x72, 0x05, 0x10, 0x40, 0x18, 0x80, 0x02, 0x52, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x98, 0x01, 0x10, 0x52, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x61, 0x76, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x47, 0x0a, 0x0b, 0x42, 0x69, 0x6e, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1f, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x22, 0x0f, 0x0a, 0x0d, 0x42, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x0b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x22, 0x33,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x32, 0xfe, 0x04, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x5e, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x4d, 0x61, 0x69, 0x6c, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x4d, 0x61, 0x69, 0x6c, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x4d, 0x61, 0x69,
	0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x46, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x19, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x11, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x22,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0c, 0x53, 0x61, 0x76, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x08, 0x42, 0x69, 0x6e, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x17, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4b,
	0x65, 0x79, 0x12, 0x15, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x42, 0x48, 0x0a, 0x19, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x42, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56,
	0x31, 0x50, 0x01, 0x5a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_account_v1_account_proto_rawDescOnce sync.Once
	file_api_account_v1_account_proto_rawDescData = file_api_account_v1_account_proto_rawDesc
)

func file_api_account_v1_account_proto_rawDescGZIP() []byte {
	file_api_account_v1_account_proto_rawDescOnce.Do(func() {
		file_api_account_v1_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_account_v1_account_proto_rawDescData)
	})
	return file_api_account_v1_account_proto_rawDescData
}

var file_api_account_v1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_api_account_v1_account_proto_goTypes = []interface{}{
	(*CreateEMailAccountReq)(nil),   // 0: account.v1.CreateEMailAccountReq
	(*CreateEMailAccountReply)(nil), // 1: account.v1.CreateEMailAccountReply
	(*GetAccountReq)(nil),           // 2: account.v1.GetAccountReq
	(*GetAccountReply)(nil),         // 3: account.v1.GetAccountReply
	(*ExistAccountEMailReq)(nil),    // 4: account.v1.ExistAccountEMailReq
	(*ExistAccountEMailReply)(nil),  // 5: account.v1.ExistAccountEMailReply
	(*VerifyPasswordReq)(nil),       // 6: account.v1.VerifyPasswordReq
	(*VerifyPasswordReply)(nil),     // 7: account.v1.VerifyPasswordReply
	(*SavePasswordReq)(nil),         // 8: account.v1.SavePasswordReq
	(*SavePasswordReply)(nil),       // 9: account.v1.SavePasswordReply
	(*BindUserReq)(nil),             // 10: account.v1.BindUserReq
	(*BindUserReply)(nil),           // 11: account.v1.BindUserReply
	(*GetKeyReq)(nil),               // 12: account.v1.GetKeyReq
	(*GetKeyReply)(nil),             // 13: account.v1.GetKeyReply
}
var file_api_account_v1_account_proto_depIdxs = []int32{
	0,  // 0: account.v1.Account.CreateEMailAccount:input_type -> account.v1.CreateEMailAccountReq
	2,  // 1: account.v1.Account.GetAccount:input_type -> account.v1.GetAccountReq
	2,  // 2: account.v1.Account.GetAccountByUserID:input_type -> account.v1.GetAccountReq
	4,  // 3: account.v1.Account.ExistAccountEMail:input_type -> account.v1.ExistAccountEMailReq
	6,  // 4: account.v1.Account.VerifyPassword:input_type -> account.v1.VerifyPasswordReq
	8,  // 5: account.v1.Account.SavePassword:input_type -> account.v1.SavePasswordReq
	10, // 6: account.v1.Account.BindUser:input_type -> account.v1.BindUserReq
	12, // 7: account.v1.Account.GetKey:input_type -> account.v1.GetKeyReq
	1,  // 8: account.v1.Account.CreateEMailAccount:output_type -> account.v1.CreateEMailAccountReply
	3,  // 9: account.v1.Account.GetAccount:output_type -> account.v1.GetAccountReply
	3,  // 10: account.v1.Account.GetAccountByUserID:output_type -> account.v1.GetAccountReply
	5,  // 11: account.v1.Account.ExistAccountEMail:output_type -> account.v1.ExistAccountEMailReply
	7,  // 12: account.v1.Account.VerifyPassword:output_type -> account.v1.VerifyPasswordReply
	9,  // 13: account.v1.Account.SavePassword:output_type -> account.v1.SavePasswordReply
	11, // 14: account.v1.Account.BindUser:output_type -> account.v1.BindUserReply
	13, // 15: account.v1.Account.GetKey:output_type -> account.v1.GetKeyReply
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_account_v1_account_proto_init() }
func file_api_account_v1_account_proto_init() {
	if File_api_account_v1_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_account_v1_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEMailAccountReq); i {
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
		file_api_account_v1_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEMailAccountReply); i {
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
		file_api_account_v1_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountReq); i {
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
		file_api_account_v1_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountReply); i {
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
		file_api_account_v1_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistAccountEMailReq); i {
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
		file_api_account_v1_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistAccountEMailReply); i {
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
		file_api_account_v1_account_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyPasswordReq); i {
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
		file_api_account_v1_account_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyPasswordReply); i {
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
		file_api_account_v1_account_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavePasswordReq); i {
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
		file_api_account_v1_account_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavePasswordReply); i {
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
		file_api_account_v1_account_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindUserReq); i {
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
		file_api_account_v1_account_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindUserReply); i {
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
		file_api_account_v1_account_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeyReq); i {
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
		file_api_account_v1_account_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeyReply); i {
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
			RawDescriptor: file_api_account_v1_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_account_v1_account_proto_goTypes,
		DependencyIndexes: file_api_account_v1_account_proto_depIdxs,
		MessageInfos:      file_api_account_v1_account_proto_msgTypes,
	}.Build()
	File_api_account_v1_account_proto = out.File
	file_api_account_v1_account_proto_rawDesc = nil
	file_api_account_v1_account_proto_goTypes = nil
	file_api_account_v1_account_proto_depIdxs = nil
}
