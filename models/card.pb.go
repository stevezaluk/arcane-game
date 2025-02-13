// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.6.1
// source: card.proto

package models

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

type CardObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description       string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Type              string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	SubTypes          []string `protobuf:"bytes,4,rep,name=subTypes,proto3" json:"subTypes,omitempty"`
	ColorIdentity     []string `protobuf:"bytes,5,rep,name=colorIdentity,proto3" json:"colorIdentity,omitempty"`
	ConvertedManaCost int64    `protobuf:"varint,6,opt,name=convertedManaCost,proto3" json:"convertedManaCost,omitempty"`
	Toughness         string   `protobuf:"bytes,7,opt,name=toughness,proto3" json:"toughness,omitempty"`
	Power             string   `protobuf:"bytes,8,opt,name=power,proto3" json:"power,omitempty"`
	Image             string   `protobuf:"bytes,9,opt,name=image,proto3" json:"image,omitempty"`
	IsTapped          bool     `protobuf:"varint,10,opt,name=isTapped,proto3" json:"isTapped,omitempty"`
	IsFaceDown        bool     `protobuf:"varint,11,opt,name=isFaceDown,proto3" json:"isFaceDown,omitempty"`
	WasPlayedThisTurn bool     `protobuf:"varint,12,opt,name=wasPlayedThisTurn,proto3" json:"wasPlayedThisTurn,omitempty"`
	Owner             string   `protobuf:"bytes,13,opt,name=owner,proto3" json:"owner,omitempty"`
	Controller        string   `protobuf:"bytes,14,opt,name=controller,proto3" json:"controller,omitempty"`
}

func (x *CardObject) Reset() {
	*x = CardObject{}
	mi := &file_card_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CardObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardObject) ProtoMessage() {}

func (x *CardObject) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardObject.ProtoReflect.Descriptor instead.
func (*CardObject) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{0}
}

func (x *CardObject) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CardObject) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CardObject) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CardObject) GetSubTypes() []string {
	if x != nil {
		return x.SubTypes
	}
	return nil
}

func (x *CardObject) GetColorIdentity() []string {
	if x != nil {
		return x.ColorIdentity
	}
	return nil
}

func (x *CardObject) GetConvertedManaCost() int64 {
	if x != nil {
		return x.ConvertedManaCost
	}
	return 0
}

func (x *CardObject) GetToughness() string {
	if x != nil {
		return x.Toughness
	}
	return ""
}

func (x *CardObject) GetPower() string {
	if x != nil {
		return x.Power
	}
	return ""
}

func (x *CardObject) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CardObject) GetIsTapped() bool {
	if x != nil {
		return x.IsTapped
	}
	return false
}

func (x *CardObject) GetIsFaceDown() bool {
	if x != nil {
		return x.IsFaceDown
	}
	return false
}

func (x *CardObject) GetWasPlayedThisTurn() bool {
	if x != nil {
		return x.WasPlayedThisTurn
	}
	return false
}

func (x *CardObject) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *CardObject) GetController() string {
	if x != nil {
		return x.Controller
	}
	return ""
}

var File_card_proto protoreflect.FileDescriptor

var file_card_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x61,
	0x6d, 0x65, 0x22, 0xb0, 0x03, 0x0a, 0x0a, 0x43, 0x61, 0x72, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a,
	0x11, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x4d, 0x61, 0x6e, 0x61, 0x43, 0x6f,
	0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x65, 0x64, 0x4d, 0x61, 0x6e, 0x61, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x6f, 0x75, 0x67, 0x68, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x6f, 0x75, 0x67, 0x68, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x77,
	0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x54, 0x61, 0x70, 0x70, 0x65,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x54, 0x61, 0x70, 0x70, 0x65,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x46, 0x61, 0x63, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x61, 0x63, 0x65, 0x44, 0x6f, 0x77,
	0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x77, 0x61, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x54, 0x68,
	0x69, 0x73, 0x54, 0x75, 0x72, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x77, 0x61,
	0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x54, 0x68, 0x69, 0x73, 0x54, 0x75, 0x72, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x65, 0x76, 0x65, 0x7a, 0x61, 0x6c, 0x75, 0x6b, 0x2f, 0x61,
	0x72, 0x63, 0x61, 0x6e, 0x65, 0x2d, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_card_proto_rawDescOnce sync.Once
	file_card_proto_rawDescData = file_card_proto_rawDesc
)

func file_card_proto_rawDescGZIP() []byte {
	file_card_proto_rawDescOnce.Do(func() {
		file_card_proto_rawDescData = protoimpl.X.CompressGZIP(file_card_proto_rawDescData)
	})
	return file_card_proto_rawDescData
}

var file_card_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_card_proto_goTypes = []any{
	(*CardObject)(nil), // 0: game.CardObject
}
var file_card_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_card_proto_init() }
func file_card_proto_init() {
	if File_card_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_card_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_card_proto_goTypes,
		DependencyIndexes: file_card_proto_depIdxs,
		MessageInfos:      file_card_proto_msgTypes,
	}.Build()
	File_card_proto = out.File
	file_card_proto_rawDesc = nil
	file_card_proto_goTypes = nil
	file_card_proto_depIdxs = nil
}
