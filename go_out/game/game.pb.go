// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: game/game.proto

package game

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetRemoteLitePlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRemoteLitePlayerRequest) Reset() {
	*x = GetRemoteLitePlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRemoteLitePlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRemoteLitePlayerRequest) ProtoMessage() {}

func (x *GetRemoteLitePlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRemoteLitePlayerRequest.ProtoReflect.Descriptor instead.
func (*GetRemoteLitePlayerRequest) Descriptor() ([]byte, []int) {
	return file_game_game_proto_rawDescGZIP(), []int{0}
}

func (x *GetRemoteLitePlayerRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetRemoteLitePlayerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *LitePlayer `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *GetRemoteLitePlayerReply) Reset() {
	*x = GetRemoteLitePlayerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRemoteLitePlayerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRemoteLitePlayerReply) ProtoMessage() {}

func (x *GetRemoteLitePlayerReply) ProtoReflect() protoreflect.Message {
	mi := &file_game_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRemoteLitePlayerReply.ProtoReflect.Descriptor instead.
func (*GetRemoteLitePlayerReply) Descriptor() ([]byte, []int) {
	return file_game_game_proto_rawDescGZIP(), []int{1}
}

func (x *GetRemoteLitePlayerReply) GetInfo() *LitePlayer {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdatePlayerExpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdatePlayerExpRequest) Reset() {
	*x = UpdatePlayerExpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePlayerExpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePlayerExpRequest) ProtoMessage() {}

func (x *UpdatePlayerExpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePlayerExpRequest.ProtoReflect.Descriptor instead.
func (*UpdatePlayerExpRequest) Descriptor() ([]byte, []int) {
	return file_game_game_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePlayerExpRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdatePlayerExpReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Exp int64 `protobuf:"varint,2,opt,name=exp,proto3" json:"exp,omitempty"`
}

func (x *UpdatePlayerExpReply) Reset() {
	*x = UpdatePlayerExpReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePlayerExpReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePlayerExpReply) ProtoMessage() {}

func (x *UpdatePlayerExpReply) ProtoReflect() protoreflect.Message {
	mi := &file_game_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePlayerExpReply.ProtoReflect.Descriptor instead.
func (*UpdatePlayerExpReply) Descriptor() ([]byte, []int) {
	return file_game_game_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePlayerExpReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdatePlayerExpReply) GetExp() int64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

var File_game_game_proto protoreflect.FileDescriptor

var file_game_game_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x1a, 0x11, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x1a, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4c, 0x69, 0x74, 0x65, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x28, 0x0a, 0x16, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x78, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x45, 0x78, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x65, 0x78, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x78, 0x70, 0x32, 0xb7,
	0x01, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x74, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x78, 0x70, 0x12, 0x1c, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x45, 0x78, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x78,
	0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x67, 0x61, 0x6d, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_game_proto_rawDescOnce sync.Once
	file_game_game_proto_rawDescData = file_game_game_proto_rawDesc
)

func file_game_game_proto_rawDescGZIP() []byte {
	file_game_game_proto_rawDescOnce.Do(func() {
		file_game_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_game_proto_rawDescData)
	})
	return file_game_game_proto_rawDescData
}

var file_game_game_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_game_game_proto_goTypes = []interface{}{
	(*GetRemoteLitePlayerRequest)(nil), // 0: game.GetRemoteLitePlayerRequest
	(*GetRemoteLitePlayerReply)(nil),   // 1: game.GetRemoteLitePlayerReply
	(*UpdatePlayerExpRequest)(nil),     // 2: game.UpdatePlayerExpRequest
	(*UpdatePlayerExpReply)(nil),       // 3: game.UpdatePlayerExpReply
	(*LitePlayer)(nil),                 // 4: game.LitePlayer
}
var file_game_game_proto_depIdxs = []int32{
	4, // 0: game.GetRemoteLitePlayerReply.info:type_name -> game.LitePlayer
	0, // 1: game.GameService.GetRemoteLitePlayer:input_type -> game.GetRemoteLitePlayerRequest
	2, // 2: game.GameService.UpdatePlayerExp:input_type -> game.UpdatePlayerExpRequest
	1, // 3: game.GameService.GetRemoteLitePlayer:output_type -> game.GetRemoteLitePlayerReply
	3, // 4: game.GameService.UpdatePlayerExp:output_type -> game.UpdatePlayerExpReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_game_game_proto_init() }
func file_game_game_proto_init() {
	if File_game_game_proto != nil {
		return
	}
	file_game_player_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_game_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRemoteLitePlayerRequest); i {
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
		file_game_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRemoteLitePlayerReply); i {
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
		file_game_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePlayerExpRequest); i {
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
		file_game_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePlayerExpReply); i {
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
			RawDescriptor: file_game_game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_game_game_proto_goTypes,
		DependencyIndexes: file_game_game_proto_depIdxs,
		MessageInfos:      file_game_game_proto_msgTypes,
	}.Build()
	File_game_game_proto = out.File
	file_game_game_proto_rawDesc = nil
	file_game_game_proto_goTypes = nil
	file_game_game_proto_depIdxs = nil
}
