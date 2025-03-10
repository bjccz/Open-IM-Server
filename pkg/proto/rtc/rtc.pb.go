// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.5
// source: rtc.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CommonResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode int32  `protobuf:"varint,1,opt,name=errCode,proto3" json:"errCode,omitempty"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
}

func (x *CommonResp) Reset() {
	*x = CommonResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResp) ProtoMessage() {}

func (x *CommonResp) ProtoReflect() protoreflect.Message {
	mi := &file_rtc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResp.ProtoReflect.Descriptor instead.
func (*CommonResp) Descriptor() ([]byte, []int) {
	return file_rtc_proto_rawDescGZIP(), []int{0}
}

func (x *CommonResp) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *CommonResp) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

type GroupInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupID       string `protobuf:"bytes,1,opt,name=groupID,proto3" json:"groupID,omitempty"`
	GroupName     string `protobuf:"bytes,2,opt,name=groupName,proto3" json:"groupName,omitempty"`
	Notification  string `protobuf:"bytes,3,opt,name=notification,proto3" json:"notification,omitempty"`
	Introduction  string `protobuf:"bytes,4,opt,name=introduction,proto3" json:"introduction,omitempty"`
	FaceURL       string `protobuf:"bytes,5,opt,name=faceURL,proto3" json:"faceURL,omitempty"`
	OwnerUserID   string `protobuf:"bytes,6,opt,name=ownerUserID,proto3" json:"ownerUserID,omitempty"`
	CreateTime    uint32 `protobuf:"varint,7,opt,name=createTime,proto3" json:"createTime,omitempty"`
	MemberCount   uint32 `protobuf:"varint,8,opt,name=memberCount,proto3" json:"memberCount,omitempty"`
	Ex            string `protobuf:"bytes,9,opt,name=ex,proto3" json:"ex,omitempty"`
	Status        int32  `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"`
	CreatorUserID string `protobuf:"bytes,11,opt,name=creatorUserID,proto3" json:"creatorUserID,omitempty"`
	GroupType     int32  `protobuf:"varint,12,opt,name=groupType,proto3" json:"groupType,omitempty"`
}

func (x *GroupInfo) Reset() {
	*x = GroupInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupInfo) ProtoMessage() {}

func (x *GroupInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rtc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupInfo.ProtoReflect.Descriptor instead.
func (*GroupInfo) Descriptor() ([]byte, []int) {
	return file_rtc_proto_rawDescGZIP(), []int{1}
}

func (x *GroupInfo) GetGroupID() string {
	if x != nil {
		return x.GroupID
	}
	return ""
}

func (x *GroupInfo) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *GroupInfo) GetNotification() string {
	if x != nil {
		return x.Notification
	}
	return ""
}

func (x *GroupInfo) GetIntroduction() string {
	if x != nil {
		return x.Introduction
	}
	return ""
}

func (x *GroupInfo) GetFaceURL() string {
	if x != nil {
		return x.FaceURL
	}
	return ""
}

func (x *GroupInfo) GetOwnerUserID() string {
	if x != nil {
		return x.OwnerUserID
	}
	return ""
}

func (x *GroupInfo) GetCreateTime() uint32 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *GroupInfo) GetMemberCount() uint32 {
	if x != nil {
		return x.MemberCount
	}
	return 0
}

func (x *GroupInfo) GetEx() string {
	if x != nil {
		return x.Ex
	}
	return ""
}

func (x *GroupInfo) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GroupInfo) GetCreatorUserID() string {
	if x != nil {
		return x.CreatorUserID
	}
	return ""
}

func (x *GroupInfo) GetGroupType() int32 {
	if x != nil {
		return x.GroupType
	}
	return 0
}

type GroupMemberFullInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupID        string `protobuf:"bytes,1,opt,name=groupID,proto3" json:"groupID,omitempty"`
	UserID         string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	RoleLevel      int32  `protobuf:"varint,3,opt,name=roleLevel,proto3" json:"roleLevel,omitempty"`
	JoinTime       int32  `protobuf:"varint,4,opt,name=joinTime,proto3" json:"joinTime,omitempty"`
	Nickname       string `protobuf:"bytes,5,opt,name=nickname,proto3" json:"nickname,omitempty"`
	FaceURL        string `protobuf:"bytes,6,opt,name=faceURL,proto3" json:"faceURL,omitempty"`
	AppMangerLevel int32  `protobuf:"varint,7,opt,name=appMangerLevel,proto3" json:"appMangerLevel,omitempty"` //if >0
	JoinSource     int32  `protobuf:"varint,8,opt,name=joinSource,proto3" json:"joinSource,omitempty"`
	OperatorUserID string `protobuf:"bytes,9,opt,name=operatorUserID,proto3" json:"operatorUserID,omitempty"`
	Ex             string `protobuf:"bytes,10,opt,name=ex,proto3" json:"ex,omitempty"`
}

func (x *GroupMemberFullInfo) Reset() {
	*x = GroupMemberFullInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupMemberFullInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMemberFullInfo) ProtoMessage() {}

func (x *GroupMemberFullInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rtc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupMemberFullInfo.ProtoReflect.Descriptor instead.
func (*GroupMemberFullInfo) Descriptor() ([]byte, []int) {
	return file_rtc_proto_rawDescGZIP(), []int{2}
}

func (x *GroupMemberFullInfo) GetGroupID() string {
	if x != nil {
		return x.GroupID
	}
	return ""
}

func (x *GroupMemberFullInfo) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GroupMemberFullInfo) GetRoleLevel() int32 {
	if x != nil {
		return x.RoleLevel
	}
	return 0
}

func (x *GroupMemberFullInfo) GetJoinTime() int32 {
	if x != nil {
		return x.JoinTime
	}
	return 0
}

func (x *GroupMemberFullInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *GroupMemberFullInfo) GetFaceURL() string {
	if x != nil {
		return x.FaceURL
	}
	return ""
}

func (x *GroupMemberFullInfo) GetAppMangerLevel() int32 {
	if x != nil {
		return x.AppMangerLevel
	}
	return 0
}

func (x *GroupMemberFullInfo) GetJoinSource() int32 {
	if x != nil {
		return x.JoinSource
	}
	return 0
}

func (x *GroupMemberFullInfo) GetOperatorUserID() string {
	if x != nil {
		return x.OperatorUserID
	}
	return ""
}

func (x *GroupMemberFullInfo) GetEx() string {
	if x != nil {
		return x.Ex
	}
	return ""
}

type ParticipantMetaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupInfo       *GroupInfo           `protobuf:"bytes,1,opt,name=groupInfo,proto3" json:"groupInfo,omitempty"`
	GroupMemberInfo *GroupMemberFullInfo `protobuf:"bytes,2,opt,name=groupMemberInfo,proto3" json:"groupMemberInfo,omitempty"`
	UserInfo        *PublicUserInfo      `protobuf:"bytes,3,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
}

func (x *ParticipantMetaData) Reset() {
	*x = ParticipantMetaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParticipantMetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipantMetaData) ProtoMessage() {}

func (x *ParticipantMetaData) ProtoReflect() protoreflect.Message {
	mi := &file_rtc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipantMetaData.ProtoReflect.Descriptor instead.
func (*ParticipantMetaData) Descriptor() ([]byte, []int) {
	return file_rtc_proto_rawDescGZIP(), []int{3}
}

func (x *ParticipantMetaData) GetGroupInfo() *GroupInfo {
	if x != nil {
		return x.GroupInfo
	}
	return nil
}

func (x *ParticipantMetaData) GetGroupMemberInfo() *GroupMemberFullInfo {
	if x != nil {
		return x.GroupMemberInfo
	}
	return nil
}

func (x *ParticipantMetaData) GetUserInfo() *PublicUserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type PublicUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	FaceURL  string `protobuf:"bytes,3,opt,name=faceURL,proto3" json:"faceURL,omitempty"`
	Gender   int32  `protobuf:"varint,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Ex       string `protobuf:"bytes,5,opt,name=ex,proto3" json:"ex,omitempty"`
}

func (x *PublicUserInfo) Reset() {
	*x = PublicUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicUserInfo) ProtoMessage() {}

func (x *PublicUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rtc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicUserInfo.ProtoReflect.Descriptor instead.
func (*PublicUserInfo) Descriptor() ([]byte, []int) {
	return file_rtc_proto_rawDescGZIP(), []int{4}
}

func (x *PublicUserInfo) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *PublicUserInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *PublicUserInfo) GetFaceURL() string {
	if x != nil {
		return x.FaceURL
	}
	return ""
}

func (x *PublicUserInfo) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *PublicUserInfo) GetEx() string {
	if x != nil {
		return x.Ex
	}
	return ""
}

type GetJoinTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room        string               `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	Identity    string               `protobuf:"bytes,2,opt,name=identity,proto3" json:"identity,omitempty"`
	MetaData    *ParticipantMetaData `protobuf:"bytes,3,opt,name=metaData,proto3" json:"metaData,omitempty"`
	OperationID string               `protobuf:"bytes,4,opt,name=operationID,proto3" json:"operationID,omitempty"`
}

func (x *GetJoinTokenReq) Reset() {
	*x = GetJoinTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJoinTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJoinTokenReq) ProtoMessage() {}

func (x *GetJoinTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_rtc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJoinTokenReq.ProtoReflect.Descriptor instead.
func (*GetJoinTokenReq) Descriptor() ([]byte, []int) {
	return file_rtc_proto_rawDescGZIP(), []int{5}
}

func (x *GetJoinTokenReq) GetRoom() string {
	if x != nil {
		return x.Room
	}
	return ""
}

func (x *GetJoinTokenReq) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *GetJoinTokenReq) GetMetaData() *ParticipantMetaData {
	if x != nil {
		return x.MetaData
	}
	return nil
}

func (x *GetJoinTokenReq) GetOperationID() string {
	if x != nil {
		return x.OperationID
	}
	return ""
}

type GetJoinTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonResp *CommonResp `protobuf:"bytes,1,opt,name=CommonResp,proto3" json:"CommonResp,omitempty"`
	Jwt        string      `protobuf:"bytes,2,opt,name=jwt,proto3" json:"jwt,omitempty"`
	LiveURL    string      `protobuf:"bytes,3,opt,name=liveURL,proto3" json:"liveURL,omitempty"`
}

func (x *GetJoinTokenResp) Reset() {
	*x = GetJoinTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJoinTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJoinTokenResp) ProtoMessage() {}

func (x *GetJoinTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_rtc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJoinTokenResp.ProtoReflect.Descriptor instead.
func (*GetJoinTokenResp) Descriptor() ([]byte, []int) {
	return file_rtc_proto_rawDescGZIP(), []int{6}
}

func (x *GetJoinTokenResp) GetCommonResp() *CommonResp {
	if x != nil {
		return x.CommonResp
	}
	return nil
}

func (x *GetJoinTokenResp) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

func (x *GetJoinTokenResp) GetLiveURL() string {
	if x != nil {
		return x.LiveURL
	}
	return ""
}

var File_rtc_proto protoreflect.FileDescriptor

var file_rtc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x74, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x63,
	0x22, 0x3e, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x4d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67,
	0x22, 0xf5, 0x02, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18,
	0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e,
	0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x66, 0x61, 0x63, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x66, 0x61, 0x63, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x65, 0x78, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x22, 0xb7, 0x02, 0x0a, 0x13, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x63, 0x65,
	0x55, 0x52, 0x4c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x63, 0x65, 0x55,
	0x52, 0x4c, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x4d, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x70, 0x70, 0x4d,
	0x61, 0x6e, 0x67, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x6a, 0x6f,
	0x69, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x6a, 0x6f, 0x69, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x65, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x65, 0x78, 0x22, 0xb8, 0x01, 0x0a, 0x13, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x09, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x72, 0x74, 0x63, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x42, 0x0a, 0x0f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x72, 0x74, 0x63, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2f, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x72, 0x74, 0x63, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x86, 0x01,
	0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x63, 0x65, 0x55, 0x52, 0x4c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x63, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x16,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x65, 0x78, 0x22, 0x99, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4a, 0x6f,
	0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f,
	0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72,
	0x74, 0x63, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x22, 0x6f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x74, 0x63,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x0a, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x76,
	0x65, 0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x69, 0x76, 0x65,
	0x55, 0x52, 0x4c, 0x32, 0x49, 0x0a, 0x0a, 0x52, 0x74, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x14, 0x2e, 0x72, 0x74, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x72, 0x74, 0x63, 0x2e, 0x47, 0x65,
	0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0b,
	0x5a, 0x09, 0x2e, 0x2f, 0x72, 0x74, 0x63, 0x3b, 0x72, 0x74, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_rtc_proto_rawDescOnce sync.Once
	file_rtc_proto_rawDescData = file_rtc_proto_rawDesc
)

func file_rtc_proto_rawDescGZIP() []byte {
	file_rtc_proto_rawDescOnce.Do(func() {
		file_rtc_proto_rawDescData = protoimpl.X.CompressGZIP(file_rtc_proto_rawDescData)
	})
	return file_rtc_proto_rawDescData
}

var file_rtc_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_rtc_proto_goTypes = []interface{}{
	(*CommonResp)(nil),          // 0: rtc.CommonResp
	(*GroupInfo)(nil),           // 1: rtc.GroupInfo
	(*GroupMemberFullInfo)(nil), // 2: rtc.GroupMemberFullInfo
	(*ParticipantMetaData)(nil), // 3: rtc.ParticipantMetaData
	(*PublicUserInfo)(nil),      // 4: rtc.PublicUserInfo
	(*GetJoinTokenReq)(nil),     // 5: rtc.GetJoinTokenReq
	(*GetJoinTokenResp)(nil),    // 6: rtc.GetJoinTokenResp
}
var file_rtc_proto_depIdxs = []int32{
	1, // 0: rtc.ParticipantMetaData.groupInfo:type_name -> rtc.GroupInfo
	2, // 1: rtc.ParticipantMetaData.groupMemberInfo:type_name -> rtc.GroupMemberFullInfo
	4, // 2: rtc.ParticipantMetaData.userInfo:type_name -> rtc.PublicUserInfo
	3, // 3: rtc.GetJoinTokenReq.metaData:type_name -> rtc.ParticipantMetaData
	0, // 4: rtc.GetJoinTokenResp.CommonResp:type_name -> rtc.CommonResp
	5, // 5: rtc.RtcService.GetJoinToken:input_type -> rtc.GetJoinTokenReq
	6, // 6: rtc.RtcService.GetJoinToken:output_type -> rtc.GetJoinTokenResp
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_rtc_proto_init() }
func file_rtc_proto_init() {
	if File_rtc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rtc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResp); i {
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
		file_rtc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupInfo); i {
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
		file_rtc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupMemberFullInfo); i {
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
		file_rtc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParticipantMetaData); i {
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
		file_rtc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicUserInfo); i {
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
		file_rtc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJoinTokenReq); i {
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
		file_rtc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJoinTokenResp); i {
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
			RawDescriptor: file_rtc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rtc_proto_goTypes,
		DependencyIndexes: file_rtc_proto_depIdxs,
		MessageInfos:      file_rtc_proto_msgTypes,
	}.Build()
	File_rtc_proto = out.File
	file_rtc_proto_rawDesc = nil
	file_rtc_proto_goTypes = nil
	file_rtc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RtcServiceClient is the client API for RtcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RtcServiceClient interface {
	GetJoinToken(ctx context.Context, in *GetJoinTokenReq, opts ...grpc.CallOption) (*GetJoinTokenResp, error)
}

type rtcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRtcServiceClient(cc grpc.ClientConnInterface) RtcServiceClient {
	return &rtcServiceClient{cc}
}

func (c *rtcServiceClient) GetJoinToken(ctx context.Context, in *GetJoinTokenReq, opts ...grpc.CallOption) (*GetJoinTokenResp, error) {
	out := new(GetJoinTokenResp)
	err := c.cc.Invoke(ctx, "/rtc.RtcService/GetJoinToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RtcServiceServer is the server API for RtcService service.
type RtcServiceServer interface {
	GetJoinToken(context.Context, *GetJoinTokenReq) (*GetJoinTokenResp, error)
}

// UnimplementedRtcServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRtcServiceServer struct {
}

func (*UnimplementedRtcServiceServer) GetJoinToken(context.Context, *GetJoinTokenReq) (*GetJoinTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJoinToken not implemented")
}

func RegisterRtcServiceServer(s *grpc.Server, srv RtcServiceServer) {
	s.RegisterService(&_RtcService_serviceDesc, srv)
}

func _RtcService_GetJoinToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJoinTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RtcServiceServer).GetJoinToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rtc.RtcService/GetJoinToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RtcServiceServer).GetJoinToken(ctx, req.(*GetJoinTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _RtcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rtc.RtcService",
	HandlerType: (*RtcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJoinToken",
			Handler:    _RtcService_GetJoinToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rtc.proto",
}
