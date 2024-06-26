// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: storage/veto.proto

package storage

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Veto_ActionType int32

const (
	Veto_VETO_ACTION_INVALID Veto_ActionType = 0
	Veto_VETO                Veto_ActionType = 1
	Veto_PICK                Veto_ActionType = 2
	Veto_SIDE                Veto_ActionType = 3 // Custom?
	Veto_SKIP                Veto_ActionType = 4
)

// Enum value maps for Veto_ActionType.
var (
	Veto_ActionType_name = map[int32]string{
		0: "VETO_ACTION_INVALID",
		1: "VETO",
		2: "PICK",
		3: "SIDE",
		4: "SKIP",
	}
	Veto_ActionType_value = map[string]int32{
		"VETO_ACTION_INVALID": 0,
		"VETO":                1,
		"PICK":                2,
		"SIDE":                3,
		"SKIP":                4,
	}
)

func (x Veto_ActionType) Enum() *Veto_ActionType {
	p := new(Veto_ActionType)
	*p = x
	return p
}

func (x Veto_ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Veto_ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_veto_proto_enumTypes[0].Descriptor()
}

func (Veto_ActionType) Type() protoreflect.EnumType {
	return &file_storage_veto_proto_enumTypes[0]
}

func (x Veto_ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Veto_ActionType.Descriptor instead.
func (Veto_ActionType) EnumDescriptor() ([]byte, []int) {
	return file_storage_veto_proto_rawDescGZIP(), []int{0, 0}
}

type Veto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Settings *Veto_Settings      `protobuf:"bytes,1,opt,name=settings,proto3" json:"settings,omitempty"`
	States   []*Veto_ActionState `protobuf:"bytes,2,rep,name=states,proto3" json:"states,omitempty"`
}

func (x *Veto) Reset() {
	*x = Veto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_veto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Veto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Veto) ProtoMessage() {}

func (x *Veto) ProtoReflect() protoreflect.Message {
	mi := &file_storage_veto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Veto.ProtoReflect.Descriptor instead.
func (*Veto) Descriptor() ([]byte, []int) {
	return file_storage_veto_proto_rawDescGZIP(), []int{0}
}

func (x *Veto) GetSettings() *Veto_Settings {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *Veto) GetStates() []*Veto_ActionState {
	if x != nil {
		return x.States
	}
	return nil
}

type Veto_Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Team   string          `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	Action Veto_ActionType `protobuf:"varint,2,opt,name=action,proto3,enum=storage.Veto_ActionType" json:"action,omitempty"`
}

func (x *Veto_Action) Reset() {
	*x = Veto_Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_veto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Veto_Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Veto_Action) ProtoMessage() {}

func (x *Veto_Action) ProtoReflect() protoreflect.Message {
	mi := &file_storage_veto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Veto_Action.ProtoReflect.Descriptor instead.
func (*Veto_Action) Descriptor() ([]byte, []int) {
	return file_storage_veto_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Veto_Action) GetTeam() string {
	if x != nil {
		return x.Team
	}
	return ""
}

func (x *Veto_Action) GetAction() Veto_ActionType {
	if x != nil {
		return x.Action
	}
	return Veto_VETO_ACTION_INVALID
}

type Veto_ActionState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time      *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Selection string                 `protobuf:"bytes,2,opt,name=selection,proto3" json:"selection,omitempty"`
}

func (x *Veto_ActionState) Reset() {
	*x = Veto_ActionState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_veto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Veto_ActionState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Veto_ActionState) ProtoMessage() {}

func (x *Veto_ActionState) ProtoReflect() protoreflect.Message {
	mi := &file_storage_veto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Veto_ActionState.ProtoReflect.Descriptor instead.
func (*Veto_ActionState) Descriptor() ([]byte, []int) {
	return file_storage_veto_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Veto_ActionState) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Veto_ActionState) GetSelection() string {
	if x != nil {
		return x.Selection
	}
	return ""
}

type Veto_Settings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Maps    []string       `protobuf:"bytes,1,rep,name=maps,proto3" json:"maps,omitempty"`
	Actions []*Veto_Action `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *Veto_Settings) Reset() {
	*x = Veto_Settings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_veto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Veto_Settings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Veto_Settings) ProtoMessage() {}

func (x *Veto_Settings) ProtoReflect() protoreflect.Message {
	mi := &file_storage_veto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Veto_Settings.ProtoReflect.Descriptor instead.
func (*Veto_Settings) Descriptor() ([]byte, []int) {
	return file_storage_veto_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Veto_Settings) GetMaps() []string {
	if x != nil {
		return x.Maps
	}
	return nil
}

func (x *Veto_Settings) GetActions() []*Veto_Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

var File_storage_veto_proto protoreflect.FileDescriptor

var file_storage_veto_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x65, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9,
	0x03, 0x0a, 0x04, 0x56, 0x65, 0x74, 0x6f, 0x12, 0x32, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x56, 0x65, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x65, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x4e,
	0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x12, 0x30, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x65, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x5b,
	0x0a, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x4e, 0x0a, 0x08, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x70, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x70, 0x73, 0x12, 0x2e, 0x0a, 0x07, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x65, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x4d, 0x0a, 0x0a, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x56, 0x45, 0x54,
	0x4f, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x45, 0x54, 0x4f, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x50, 0x49, 0x43, 0x4b, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x49, 0x44, 0x45, 0x10, 0x03,
	0x12, 0x08, 0x0a, 0x04, 0x53, 0x4b, 0x49, 0x50, 0x10, 0x04, 0x42, 0x88, 0x01, 0x0a, 0x0b, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x09, 0x56, 0x65, 0x74, 0x6f,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74,
	0x69, 0x2f, 0x6a, 0x75, 0x6d, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x67, 0x65, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0xa2, 0x02, 0x03, 0x53, 0x58,
	0x58, 0xaa, 0x02, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0xca, 0x02, 0x07, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0xe2, 0x02, 0x13, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_veto_proto_rawDescOnce sync.Once
	file_storage_veto_proto_rawDescData = file_storage_veto_proto_rawDesc
)

func file_storage_veto_proto_rawDescGZIP() []byte {
	file_storage_veto_proto_rawDescOnce.Do(func() {
		file_storage_veto_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_veto_proto_rawDescData)
	})
	return file_storage_veto_proto_rawDescData
}

var file_storage_veto_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_storage_veto_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_storage_veto_proto_goTypes = []interface{}{
	(Veto_ActionType)(0),          // 0: storage.Veto.ActionType
	(*Veto)(nil),                  // 1: storage.Veto
	(*Veto_Action)(nil),           // 2: storage.Veto.Action
	(*Veto_ActionState)(nil),      // 3: storage.Veto.ActionState
	(*Veto_Settings)(nil),         // 4: storage.Veto.Settings
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_storage_veto_proto_depIdxs = []int32{
	4, // 0: storage.Veto.settings:type_name -> storage.Veto.Settings
	3, // 1: storage.Veto.states:type_name -> storage.Veto.ActionState
	0, // 2: storage.Veto.Action.action:type_name -> storage.Veto.ActionType
	5, // 3: storage.Veto.ActionState.time:type_name -> google.protobuf.Timestamp
	2, // 4: storage.Veto.Settings.actions:type_name -> storage.Veto.Action
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_storage_veto_proto_init() }
func file_storage_veto_proto_init() {
	if File_storage_veto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storage_veto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Veto); i {
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
		file_storage_veto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Veto_Action); i {
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
		file_storage_veto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Veto_ActionState); i {
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
		file_storage_veto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Veto_Settings); i {
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
			RawDescriptor: file_storage_veto_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_veto_proto_goTypes,
		DependencyIndexes: file_storage_veto_proto_depIdxs,
		EnumInfos:         file_storage_veto_proto_enumTypes,
		MessageInfos:      file_storage_veto_proto_msgTypes,
	}.Build()
	File_storage_veto_proto = out.File
	file_storage_veto_proto_rawDesc = nil
	file_storage_veto_proto_goTypes = nil
	file_storage_veto_proto_depIdxs = nil
}
