// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: user/user.proto

package userv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // Идентификатор пользователя
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserData *User `protobuf:"bytes,1,opt,name=user_data,json=userData,proto3" json:"user_data,omitempty"` // Информация о пользователе
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserResponse) GetUserData() *User {
	if x != nil {
		return x.UserData
	}
	return nil
}

type PartialUpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`            // Идентификатор пользователя
	UserData   *User                  `protobuf:"bytes,2,opt,name=user_data,json=userData,proto3" json:"user_data,omitempty"`       // Данные пользователя для обновления
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"` // Указание полей для обновления
}

func (x *PartialUpdateUserRequest) Reset() {
	*x = PartialUpdateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartialUpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartialUpdateUserRequest) ProtoMessage() {}

func (x *PartialUpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartialUpdateUserRequest.ProtoReflect.Descriptor instead.
func (*PartialUpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *PartialUpdateUserRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PartialUpdateUserRequest) GetUserData() *User {
	if x != nil {
		return x.UserData
	}
	return nil
}

func (x *PartialUpdateUserRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type PartialUpdateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdatedUser *User `protobuf:"bytes,1,opt,name=updated_user,json=updatedUser,proto3" json:"updated_user,omitempty"` // Обновленные данные пользователя
}

func (x *PartialUpdateUserResponse) Reset() {
	*x = PartialUpdateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartialUpdateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartialUpdateUserResponse) ProtoMessage() {}

func (x *PartialUpdateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartialUpdateUserResponse.ProtoReflect.Descriptor instead.
func (*PartialUpdateUserResponse) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *PartialUpdateUserResponse) GetUpdatedUser() *User {
	if x != nil {
		return x.UpdatedUser
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TelegramId        int64  `protobuf:"varint,2,opt,name=telegram_id,json=telegramId,proto3" json:"telegram_id,omitempty"`
	TelegramUsername  string `protobuf:"bytes,3,opt,name=telegram_username,json=telegramUsername,proto3" json:"telegram_username,omitempty"`
	TelegramFirstName string `protobuf:"bytes,4,opt,name=telegram_first_name,json=telegramFirstName,proto3" json:"telegram_first_name,omitempty"`
	TelegramLastName  string `protobuf:"bytes,5,opt,name=telegram_last_name,json=telegramLastName,proto3" json:"telegram_last_name,omitempty"`
	TelegramAvatarUrl string `protobuf:"bytes,6,opt,name=telegram_avatar_url,json=telegramAvatarUrl,proto3" json:"telegram_avatar_url,omitempty"`
	LanguageCode      string `protobuf:"bytes,7,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	IsPremium         bool   `protobuf:"varint,8,opt,name=is_premium,json=isPremium,proto3" json:"is_premium,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetTelegramId() int64 {
	if x != nil {
		return x.TelegramId
	}
	return 0
}

func (x *User) GetTelegramUsername() string {
	if x != nil {
		return x.TelegramUsername
	}
	return ""
}

func (x *User) GetTelegramFirstName() string {
	if x != nil {
		return x.TelegramFirstName
	}
	return ""
}

func (x *User) GetTelegramLastName() string {
	if x != nil {
		return x.TelegramLastName
	}
	return ""
}

func (x *User) GetTelegramAvatarUrl() string {
	if x != nil {
		return x.TelegramAvatarUrl
	}
	return ""
}

func (x *User) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

func (x *User) GetIsPremium() bool {
	if x != nil {
		return x.IsPremium
	}
	return false
}

var File_user_user_proto protoreflect.FileDescriptor

var file_user_user_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0b, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x37, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x27, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x22, 0x99, 0x01, 0x0a, 0x18, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x27, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x4a, 0x0a, 0x19, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61,
	0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x22, 0xb6, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11,
	0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61,
	0x6d, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x65, 0x6c,
	0x65, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d,
	0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x65, 0x6c,
	0x65, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x4c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x65, 0x6c, 0x65, 0x67,
	0x72, 0x61, 0x6d, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x32, 0x95, 0x01, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a,
	0x11, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61,
	0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61,
	0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x62, 0x6f, 0x62, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_user_user_proto_rawDescOnce sync.Once
	file_user_user_proto_rawDescData = file_user_user_proto_rawDesc
)

func file_user_user_proto_rawDescGZIP() []byte {
	file_user_user_proto_rawDescOnce.Do(func() {
		file_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_user_proto_rawDescData)
	})
	return file_user_user_proto_rawDescData
}

var file_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_user_user_proto_goTypes = []any{
	(*UserRequest)(nil),               // 0: user.UserRequest
	(*UserResponse)(nil),              // 1: user.UserResponse
	(*PartialUpdateUserRequest)(nil),  // 2: user.PartialUpdateUserRequest
	(*PartialUpdateUserResponse)(nil), // 3: user.PartialUpdateUserResponse
	(*User)(nil),                      // 4: user.User
	(*fieldmaskpb.FieldMask)(nil),     // 5: google.protobuf.FieldMask
}
var file_user_user_proto_depIdxs = []int32{
	4, // 0: user.UserResponse.user_data:type_name -> user.User
	4, // 1: user.PartialUpdateUserRequest.user_data:type_name -> user.User
	5, // 2: user.PartialUpdateUserRequest.update_mask:type_name -> google.protobuf.FieldMask
	4, // 3: user.PartialUpdateUserResponse.updated_user:type_name -> user.User
	0, // 4: user.UserService.GetUser:input_type -> user.UserRequest
	2, // 5: user.UserService.PartialUpdateUser:input_type -> user.PartialUpdateUserRequest
	1, // 6: user.UserService.GetUser:output_type -> user.UserResponse
	3, // 7: user.UserService.PartialUpdateUser:output_type -> user.PartialUpdateUserResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_user_user_proto_init() }
func file_user_user_proto_init() {
	if File_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_user_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UserRequest); i {
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
		file_user_user_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UserResponse); i {
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
		file_user_user_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PartialUpdateUserRequest); i {
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
		file_user_user_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PartialUpdateUserResponse); i {
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
		file_user_user_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*User); i {
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
			RawDescriptor: file_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_user_proto_goTypes,
		DependencyIndexes: file_user_user_proto_depIdxs,
		MessageInfos:      file_user_user_proto_msgTypes,
	}.Build()
	File_user_user_proto = out.File
	file_user_user_proto_rawDesc = nil
	file_user_user_proto_goTypes = nil
	file_user_user_proto_depIdxs = nil
}
