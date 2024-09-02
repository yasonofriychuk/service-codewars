// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: service/service.proto

package codewars

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

type GetUserInfoIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserInfoIn) Reset() {
	*x = GetUserInfoIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoIn) ProtoMessage() {}

func (x *GetUserInfoIn) ProtoReflect() protoreflect.Message {
	mi := &file_service_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoIn.ProtoReflect.Descriptor instead.
func (*GetUserInfoIn) Descriptor() ([]byte, []int) {
	return file_service_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserInfoIn) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type GetUserInfoOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username            string                         `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Name                string                         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Honor               uint32                         `protobuf:"varint,3,opt,name=honor,proto3" json:"honor,omitempty"`
	Clan                string                         `protobuf:"bytes,4,opt,name=clan,proto3" json:"clan,omitempty"`
	LeaderboardPosition int64                          `protobuf:"varint,5,opt,name=leaderboardPosition,proto3" json:"leaderboardPosition,omitempty"`
	Skills              []string                       `protobuf:"bytes,6,rep,name=skills,proto3" json:"skills,omitempty"`
	Ranks               *GetUserInfoOut_Ranks          `protobuf:"bytes,7,opt,name=ranks,proto3" json:"ranks,omitempty"`
	CodeChallenges      *GetUserInfoOut_CodeChallenges `protobuf:"bytes,8,opt,name=codeChallenges,proto3" json:"codeChallenges,omitempty"`
}

func (x *GetUserInfoOut) Reset() {
	*x = GetUserInfoOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoOut) ProtoMessage() {}

func (x *GetUserInfoOut) ProtoReflect() protoreflect.Message {
	mi := &file_service_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoOut.ProtoReflect.Descriptor instead.
func (*GetUserInfoOut) Descriptor() ([]byte, []int) {
	return file_service_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserInfoOut) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetUserInfoOut) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetUserInfoOut) GetHonor() uint32 {
	if x != nil {
		return x.Honor
	}
	return 0
}

func (x *GetUserInfoOut) GetClan() string {
	if x != nil {
		return x.Clan
	}
	return ""
}

func (x *GetUserInfoOut) GetLeaderboardPosition() int64 {
	if x != nil {
		return x.LeaderboardPosition
	}
	return 0
}

func (x *GetUserInfoOut) GetSkills() []string {
	if x != nil {
		return x.Skills
	}
	return nil
}

func (x *GetUserInfoOut) GetRanks() *GetUserInfoOut_Ranks {
	if x != nil {
		return x.Ranks
	}
	return nil
}

func (x *GetUserInfoOut) GetCodeChallenges() *GetUserInfoOut_CodeChallenges {
	if x != nil {
		return x.CodeChallenges
	}
	return nil
}

type GetUserInfoOut_Ranks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Overall   *GetUserInfoOut_Rank            `protobuf:"bytes,1,opt,name=overall,proto3" json:"overall,omitempty"`
	Languages map[string]*GetUserInfoOut_Rank `protobuf:"bytes,2,rep,name=languages,proto3" json:"languages,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetUserInfoOut_Ranks) Reset() {
	*x = GetUserInfoOut_Ranks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoOut_Ranks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoOut_Ranks) ProtoMessage() {}

func (x *GetUserInfoOut_Ranks) ProtoReflect() protoreflect.Message {
	mi := &file_service_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoOut_Ranks.ProtoReflect.Descriptor instead.
func (*GetUserInfoOut_Ranks) Descriptor() ([]byte, []int) {
	return file_service_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetUserInfoOut_Ranks) GetOverall() *GetUserInfoOut_Rank {
	if x != nil {
		return x.Overall
	}
	return nil
}

func (x *GetUserInfoOut_Ranks) GetLanguages() map[string]*GetUserInfoOut_Rank {
	if x != nil {
		return x.Languages
	}
	return nil
}

type GetUserInfoOut_Rank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rank  int64  `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Color string `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	Score int64  `protobuf:"varint,4,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *GetUserInfoOut_Rank) Reset() {
	*x = GetUserInfoOut_Rank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoOut_Rank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoOut_Rank) ProtoMessage() {}

func (x *GetUserInfoOut_Rank) ProtoReflect() protoreflect.Message {
	mi := &file_service_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoOut_Rank.ProtoReflect.Descriptor instead.
func (*GetUserInfoOut_Rank) Descriptor() ([]byte, []int) {
	return file_service_service_proto_rawDescGZIP(), []int{1, 1}
}

func (x *GetUserInfoOut_Rank) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *GetUserInfoOut_Rank) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetUserInfoOut_Rank) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *GetUserInfoOut_Rank) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

type GetUserInfoOut_CodeChallenges struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalAuthored  uint32 `protobuf:"varint,1,opt,name=totalAuthored,proto3" json:"totalAuthored,omitempty"`
	TotalCompleted uint32 `protobuf:"varint,2,opt,name=totalCompleted,proto3" json:"totalCompleted,omitempty"`
}

func (x *GetUserInfoOut_CodeChallenges) Reset() {
	*x = GetUserInfoOut_CodeChallenges{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoOut_CodeChallenges) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoOut_CodeChallenges) ProtoMessage() {}

func (x *GetUserInfoOut_CodeChallenges) ProtoReflect() protoreflect.Message {
	mi := &file_service_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoOut_CodeChallenges.ProtoReflect.Descriptor instead.
func (*GetUserInfoOut_CodeChallenges) Descriptor() ([]byte, []int) {
	return file_service_service_proto_rawDescGZIP(), []int{1, 2}
}

func (x *GetUserInfoOut_CodeChallenges) GetTotalAuthored() uint32 {
	if x != nil {
		return x.TotalAuthored
	}
	return 0
}

func (x *GetUserInfoOut_CodeChallenges) GetTotalCompleted() uint32 {
	if x != nil {
		return x.TotalCompleted
	}
	return 0
}

var File_service_service_proto protoreflect.FileDescriptor

var file_service_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x77, 0x61, 0x72,
	0x73, 0x22, 0x23, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x49, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xe4, 0x05, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4f, 0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x6e,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x68, 0x6f, 0x6e, 0x6f, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6c, 0x61, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6c, 0x61, 0x6e, 0x12, 0x30, 0x0a, 0x13, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x13, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x34, 0x0a,
	0x05, 0x72, 0x61, 0x6e, 0x6b, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x4f, 0x75, 0x74, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x73, 0x52, 0x05, 0x72, 0x61,
	0x6e, 0x6b, 0x73, 0x12, 0x4f, 0x0a, 0x0e, 0x63, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6c, 0x6c,
	0x65, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x4f, 0x75, 0x74, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x73, 0x52, 0x0e, 0x63, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x73, 0x1a, 0xea, 0x01, 0x0a, 0x05, 0x52, 0x61, 0x6e, 0x6b, 0x73, 0x12, 0x37,
	0x0a, 0x07, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4f, 0x75, 0x74, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x07,
	0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x12, 0x4b, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x4f, 0x75, 0x74, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x73, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x73, 0x1a, 0x5b, 0x0a, 0x0e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x77, 0x61,
	0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4f, 0x75,
	0x74, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x5a, 0x0a, 0x04, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x5e, 0x0a,
	0x0e, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x73, 0x12,
	0x24, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x32, 0x4c, 0x0a,
	0x08, 0x43, 0x6f, 0x64, 0x65, 0x77, 0x61, 0x72, 0x73, 0x12, 0x40, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x77,
	0x61, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x49,
	0x6e, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4f, 0x75, 0x74, 0x42, 0x22, 0x5a, 0x20, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x63, 0x6f, 0x64, 0x65, 0x77, 0x61, 0x72, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_service_proto_rawDescOnce sync.Once
	file_service_service_proto_rawDescData = file_service_service_proto_rawDesc
)

func file_service_service_proto_rawDescGZIP() []byte {
	file_service_service_proto_rawDescOnce.Do(func() {
		file_service_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_service_proto_rawDescData)
	})
	return file_service_service_proto_rawDescData
}

var file_service_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_service_service_proto_goTypes = []any{
	(*GetUserInfoIn)(nil),                 // 0: codewars.GetUserInfoIn
	(*GetUserInfoOut)(nil),                // 1: codewars.GetUserInfoOut
	(*GetUserInfoOut_Ranks)(nil),          // 2: codewars.GetUserInfoOut.Ranks
	(*GetUserInfoOut_Rank)(nil),           // 3: codewars.GetUserInfoOut.Rank
	(*GetUserInfoOut_CodeChallenges)(nil), // 4: codewars.GetUserInfoOut.CodeChallenges
	nil,                                   // 5: codewars.GetUserInfoOut.Ranks.LanguagesEntry
}
var file_service_service_proto_depIdxs = []int32{
	2, // 0: codewars.GetUserInfoOut.ranks:type_name -> codewars.GetUserInfoOut.Ranks
	4, // 1: codewars.GetUserInfoOut.codeChallenges:type_name -> codewars.GetUserInfoOut.CodeChallenges
	3, // 2: codewars.GetUserInfoOut.Ranks.overall:type_name -> codewars.GetUserInfoOut.Rank
	5, // 3: codewars.GetUserInfoOut.Ranks.languages:type_name -> codewars.GetUserInfoOut.Ranks.LanguagesEntry
	3, // 4: codewars.GetUserInfoOut.Ranks.LanguagesEntry.value:type_name -> codewars.GetUserInfoOut.Rank
	0, // 5: codewars.Codewars.GetUserInfo:input_type -> codewars.GetUserInfoIn
	1, // 6: codewars.Codewars.GetUserInfo:output_type -> codewars.GetUserInfoOut
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_service_service_proto_init() }
func file_service_service_proto_init() {
	if File_service_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserInfoIn); i {
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
		file_service_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserInfoOut); i {
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
		file_service_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserInfoOut_Ranks); i {
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
		file_service_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserInfoOut_Rank); i {
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
		file_service_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserInfoOut_CodeChallenges); i {
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
			RawDescriptor: file_service_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_service_proto_goTypes,
		DependencyIndexes: file_service_service_proto_depIdxs,
		MessageInfos:      file_service_service_proto_msgTypes,
	}.Build()
	File_service_service_proto = out.File
	file_service_service_proto_rawDesc = nil
	file_service_service_proto_goTypes = nil
	file_service_service_proto_depIdxs = nil
}
