// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.6.1
// source: player.proto

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

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email              string      `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Deck               *DeckObject `protobuf:"bytes,2,opt,name=deck,proto3" json:"deck,omitempty"`
	Graveyard          *Zone       `protobuf:"bytes,3,opt,name=graveyard,proto3" json:"graveyard,omitempty"`
	Hand               *Zone       `protobuf:"bytes,4,opt,name=hand,proto3" json:"hand,omitempty"`
	LifeTotal          int64       `protobuf:"varint,5,opt,name=lifeTotal,proto3" json:"lifeTotal,omitempty"`
	CommanderDamage    int64       `protobuf:"varint,6,opt,name=commanderDamage,proto3" json:"commanderDamage,omitempty"`
	PoisonCounters     int64       `protobuf:"varint,7,opt,name=poisonCounters,proto3" json:"poisonCounters,omitempty"`
	EnergyCounters     int64       `protobuf:"varint,8,opt,name=energyCounters,proto3" json:"energyCounters,omitempty"`
	ExperienceCounters int64       `protobuf:"varint,9,opt,name=experienceCounters,proto3" json:"experienceCounters,omitempty"`
	IsMonarch          bool        `protobuf:"varint,10,opt,name=isMonarch,proto3" json:"isMonarch,omitempty"`
	IsGameOwner        bool        `protobuf:"varint,11,opt,name=isGameOwner,proto3" json:"isGameOwner,omitempty"`
	CurrentPhase       string      `protobuf:"bytes,12,opt,name=currentPhase,proto3" json:"currentPhase,omitempty"`
	CurrentStep        string      `protobuf:"bytes,13,opt,name=currentStep,proto3" json:"currentStep,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	mi := &file_player_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{0}
}

func (x *Player) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Player) GetDeck() *DeckObject {
	if x != nil {
		return x.Deck
	}
	return nil
}

func (x *Player) GetGraveyard() *Zone {
	if x != nil {
		return x.Graveyard
	}
	return nil
}

func (x *Player) GetHand() *Zone {
	if x != nil {
		return x.Hand
	}
	return nil
}

func (x *Player) GetLifeTotal() int64 {
	if x != nil {
		return x.LifeTotal
	}
	return 0
}

func (x *Player) GetCommanderDamage() int64 {
	if x != nil {
		return x.CommanderDamage
	}
	return 0
}

func (x *Player) GetPoisonCounters() int64 {
	if x != nil {
		return x.PoisonCounters
	}
	return 0
}

func (x *Player) GetEnergyCounters() int64 {
	if x != nil {
		return x.EnergyCounters
	}
	return 0
}

func (x *Player) GetExperienceCounters() int64 {
	if x != nil {
		return x.ExperienceCounters
	}
	return 0
}

func (x *Player) GetIsMonarch() bool {
	if x != nil {
		return x.IsMonarch
	}
	return false
}

func (x *Player) GetIsGameOwner() bool {
	if x != nil {
		return x.IsGameOwner
	}
	return false
}

func (x *Player) GetCurrentPhase() string {
	if x != nil {
		return x.CurrentPhase
	}
	return ""
}

func (x *Player) GetCurrentStep() string {
	if x != nil {
		return x.CurrentStep
	}
	return ""
}

var File_player_proto protoreflect.FileDescriptor

var file_player_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x67, 0x61, 0x6d, 0x65, 0x1a, 0x0a, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0a, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x03, 0x0a,
	0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x24, 0x0a,
	0x04, 0x64, 0x65, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x6b, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x04, 0x64,
	0x65, 0x63, 0x6b, 0x12, 0x28, 0x0a, 0x09, 0x67, 0x72, 0x61, 0x76, 0x65, 0x79, 0x61, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x5a, 0x6f,
	0x6e, 0x65, 0x52, 0x09, 0x67, 0x72, 0x61, 0x76, 0x65, 0x79, 0x61, 0x72, 0x64, 0x12, 0x1e, 0x0a,
	0x04, 0x68, 0x61, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x04, 0x68, 0x61, 0x6e, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x6c, 0x69, 0x66, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x44,
	0x61, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x6f, 0x69, 0x73, 0x6f, 0x6e, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x70,
	0x6f, 0x69, 0x73, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x12, 0x26, 0x0a,
	0x0e, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x12, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x4d, 0x6f, 0x6e, 0x61, 0x72,
	0x63, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x4d, 0x6f, 0x6e, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x47, 0x61, 0x6d, 0x65, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x47, 0x61, 0x6d, 0x65,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x50, 0x68, 0x61, 0x73, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x50, 0x68, 0x61, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x65, 0x70, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x65, 0x70, 0x42, 0x2a, 0x5a, 0x28, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x65, 0x76, 0x65, 0x7a,
	0x61, 0x6c, 0x75, 0x6b, 0x2f, 0x61, 0x72, 0x63, 0x61, 0x6e, 0x65, 0x2d, 0x67, 0x61, 0x6d, 0x65,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_player_proto_rawDescOnce sync.Once
	file_player_proto_rawDescData = file_player_proto_rawDesc
)

func file_player_proto_rawDescGZIP() []byte {
	file_player_proto_rawDescOnce.Do(func() {
		file_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_player_proto_rawDescData)
	})
	return file_player_proto_rawDescData
}

var file_player_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_player_proto_goTypes = []any{
	(*Player)(nil),     // 0: game.Player
	(*DeckObject)(nil), // 1: game.DeckObject
	(*Zone)(nil),       // 2: game.Zone
}
var file_player_proto_depIdxs = []int32{
	1, // 0: game.Player.deck:type_name -> game.DeckObject
	2, // 1: game.Player.graveyard:type_name -> game.Zone
	2, // 2: game.Player.hand:type_name -> game.Zone
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_player_proto_init() }
func file_player_proto_init() {
	if File_player_proto != nil {
		return
	}
	file_zone_proto_init()
	file_deck_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_player_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_player_proto_goTypes,
		DependencyIndexes: file_player_proto_depIdxs,
		MessageInfos:      file_player_proto_msgTypes,
	}.Build()
	File_player_proto = out.File
	file_player_proto_rawDesc = nil
	file_player_proto_goTypes = nil
	file_player_proto_depIdxs = nil
}
