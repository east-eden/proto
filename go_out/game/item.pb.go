// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: game/item.proto

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

////////////////////////////////////////////////
// Item
type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	TypeId     int32 `protobuf:"varint,2,opt,name=TypeId,proto3" json:"TypeId,omitempty"`
	Num        int32 `protobuf:"varint,3,opt,name=Num,proto3" json:"Num,omitempty"`
	EquipObjId int64 `protobuf:"varint,4,opt,name=EquipObjId,proto3" json:"EquipObjId,omitempty"` // 装备者objid
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_game_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_game_item_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Item) GetTypeId() int32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *Item) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *Item) GetEquipObjId() int64 {
	if x != nil {
		return x.EquipObjId
	}
	return 0
}

type C2M_AddItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeId int32 `protobuf:"varint,1,opt,name=TypeId,proto3" json:"TypeId,omitempty"`
}

func (x *C2M_AddItem) Reset() {
	*x = C2M_AddItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2M_AddItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2M_AddItem) ProtoMessage() {}

func (x *C2M_AddItem) ProtoReflect() protoreflect.Message {
	mi := &file_game_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2M_AddItem.ProtoReflect.Descriptor instead.
func (*C2M_AddItem) Descriptor() ([]byte, []int) {
	return file_game_item_proto_rawDescGZIP(), []int{1}
}

func (x *C2M_AddItem) GetTypeId() int32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

type C2M_DelItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *C2M_DelItem) Reset() {
	*x = C2M_DelItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2M_DelItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2M_DelItem) ProtoMessage() {}

func (x *C2M_DelItem) ProtoReflect() protoreflect.Message {
	mi := &file_game_item_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2M_DelItem.ProtoReflect.Descriptor instead.
func (*C2M_DelItem) Descriptor() ([]byte, []int) {
	return file_game_item_proto_rawDescGZIP(), []int{2}
}

func (x *C2M_DelItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type C2M_UseItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId int64 `protobuf:"varint,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
}

func (x *C2M_UseItem) Reset() {
	*x = C2M_UseItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_item_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2M_UseItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2M_UseItem) ProtoMessage() {}

func (x *C2M_UseItem) ProtoReflect() protoreflect.Message {
	mi := &file_game_item_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2M_UseItem.ProtoReflect.Descriptor instead.
func (*C2M_UseItem) Descriptor() ([]byte, []int) {
	return file_game_item_proto_rawDescGZIP(), []int{3}
}

func (x *C2M_UseItem) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

type C2M_QueryItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *C2M_QueryItems) Reset() {
	*x = C2M_QueryItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_item_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2M_QueryItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2M_QueryItems) ProtoMessage() {}

func (x *C2M_QueryItems) ProtoReflect() protoreflect.Message {
	mi := &file_game_item_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2M_QueryItems.ProtoReflect.Descriptor instead.
func (*C2M_QueryItems) Descriptor() ([]byte, []int) {
	return file_game_item_proto_rawDescGZIP(), []int{4}
}

type M2C_ItemList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *M2C_ItemList) Reset() {
	*x = M2C_ItemList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_item_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M2C_ItemList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M2C_ItemList) ProtoMessage() {}

func (x *M2C_ItemList) ProtoReflect() protoreflect.Message {
	mi := &file_game_item_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M2C_ItemList.ProtoReflect.Descriptor instead.
func (*M2C_ItemList) Descriptor() ([]byte, []int) {
	return file_game_item_proto_rawDescGZIP(), []int{5}
}

func (x *M2C_ItemList) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type M2C_ItemAdd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *Item `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *M2C_ItemAdd) Reset() {
	*x = M2C_ItemAdd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_item_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M2C_ItemAdd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M2C_ItemAdd) ProtoMessage() {}

func (x *M2C_ItemAdd) ProtoReflect() protoreflect.Message {
	mi := &file_game_item_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M2C_ItemAdd.ProtoReflect.Descriptor instead.
func (*M2C_ItemAdd) Descriptor() ([]byte, []int) {
	return file_game_item_proto_rawDescGZIP(), []int{6}
}

func (x *M2C_ItemAdd) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

type M2C_ItemUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *Item `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *M2C_ItemUpdate) Reset() {
	*x = M2C_ItemUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_item_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M2C_ItemUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M2C_ItemUpdate) ProtoMessage() {}

func (x *M2C_ItemUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_game_item_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M2C_ItemUpdate.ProtoReflect.Descriptor instead.
func (*M2C_ItemUpdate) Descriptor() ([]byte, []int) {
	return file_game_item_proto_rawDescGZIP(), []int{7}
}

func (x *M2C_ItemUpdate) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

type M2C_DelItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId int64 `protobuf:"varint,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
}

func (x *M2C_DelItem) Reset() {
	*x = M2C_DelItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_item_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M2C_DelItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M2C_DelItem) ProtoMessage() {}

func (x *M2C_DelItem) ProtoReflect() protoreflect.Message {
	mi := &file_game_item_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M2C_DelItem.ProtoReflect.Descriptor instead.
func (*M2C_DelItem) Descriptor() ([]byte, []int) {
	return file_game_item_proto_rawDescGZIP(), []int{8}
}

func (x *M2C_DelItem) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

var File_game_item_proto protoreflect.FileDescriptor

var file_game_item_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x22, 0x60, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x75, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x4e, 0x75, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x71, 0x75,
	0x69, 0x70, 0x4f, 0x62, 0x6a, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x45,
	0x71, 0x75, 0x69, 0x70, 0x4f, 0x62, 0x6a, 0x49, 0x64, 0x22, 0x25, 0x0a, 0x0b, 0x43, 0x32, 0x4d,
	0x5f, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64,
	0x22, 0x1d, 0x0a, 0x0b, 0x43, 0x32, 0x4d, 0x5f, 0x44, 0x65, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22,
	0x25, 0x0a, 0x0b, 0x43, 0x32, 0x4d, 0x5f, 0x55, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16,
	0x0a, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x32, 0x4d, 0x5f, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x30, 0x0a, 0x0c, 0x4d, 0x32, 0x43, 0x5f,
	0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x2d, 0x0a, 0x0b, 0x4d, 0x32,
	0x43, 0x5f, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x64, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x69, 0x74, 0x65,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x30, 0x0a, 0x0e, 0x4d, 0x32, 0x43,
	0x5f, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x69,
	0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x25, 0x0a, 0x0b, 0x4d,
	0x32, 0x43, 0x5f, 0x44, 0x65, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x74,
	0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x49, 0x74, 0x65, 0x6d,
	0x49, 0x64, 0x42, 0x28, 0x5a, 0x26, 0x65, 0x2e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x6e,
	0x65, 0x74, 0x2f, 0x6d, 0x6d, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2f, 0x62, 0x6c, 0x61, 0x64,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_item_proto_rawDescOnce sync.Once
	file_game_item_proto_rawDescData = file_game_item_proto_rawDesc
)

func file_game_item_proto_rawDescGZIP() []byte {
	file_game_item_proto_rawDescOnce.Do(func() {
		file_game_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_item_proto_rawDescData)
	})
	return file_game_item_proto_rawDescData
}

var file_game_item_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_game_item_proto_goTypes = []interface{}{
	(*Item)(nil),           // 0: game.Item
	(*C2M_AddItem)(nil),    // 1: game.C2M_AddItem
	(*C2M_DelItem)(nil),    // 2: game.C2M_DelItem
	(*C2M_UseItem)(nil),    // 3: game.C2M_UseItem
	(*C2M_QueryItems)(nil), // 4: game.C2M_QueryItems
	(*M2C_ItemList)(nil),   // 5: game.M2C_ItemList
	(*M2C_ItemAdd)(nil),    // 6: game.M2C_ItemAdd
	(*M2C_ItemUpdate)(nil), // 7: game.M2C_ItemUpdate
	(*M2C_DelItem)(nil),    // 8: game.M2C_DelItem
}
var file_game_item_proto_depIdxs = []int32{
	0, // 0: game.M2C_ItemList.items:type_name -> game.Item
	0, // 1: game.M2C_ItemAdd.item:type_name -> game.Item
	0, // 2: game.M2C_ItemUpdate.item:type_name -> game.Item
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_game_item_proto_init() }
func file_game_item_proto_init() {
	if File_game_item_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_game_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2M_AddItem); i {
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
		file_game_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2M_DelItem); i {
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
		file_game_item_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2M_UseItem); i {
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
		file_game_item_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2M_QueryItems); i {
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
		file_game_item_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M2C_ItemList); i {
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
		file_game_item_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M2C_ItemAdd); i {
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
		file_game_item_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M2C_ItemUpdate); i {
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
		file_game_item_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M2C_DelItem); i {
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
			RawDescriptor: file_game_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_item_proto_goTypes,
		DependencyIndexes: file_game_item_proto_depIdxs,
		MessageInfos:      file_game_item_proto_msgTypes,
	}.Build()
	File_game_item_proto = out.File
	file_game_item_proto_rawDesc = nil
	file_game_item_proto_goTypes = nil
	file_game_item_proto_depIdxs = nil
}
