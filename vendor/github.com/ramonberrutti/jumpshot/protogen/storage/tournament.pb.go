// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: storage/tournament.proto

package storage

import (
	v1 "github.com/ramonberrutti/jumpshot/protogen/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

type Tournament struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string                         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             string                         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description      string                         `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreateTime       *timestamppb.Timestamp         `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime       *timestamppb.Timestamp         `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	Version          uint64                         `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
	ParticipantsMode v1.Tournament_ParticipantsMode `protobuf:"varint,7,opt,name=participants_mode,json=participantsMode,proto3,enum=common.v1.Tournament_ParticipantsMode" json:"participants_mode,omitempty"`
	// Types that are assignable to Participants:
	//
	//	*Tournament_Teams
	//	*Tournament_Players
	Participants isTournament_Participants `protobuf_oneof:"participants"`
	// phases contains the phases of the tournament.
	Phases []*v1.Tournament_Phase `protobuf:"bytes,10,rep,name=phases,proto3" json:"phases,omitempty"`
	// additional_matches are created when are not related to a phase.
	AdditionalMatches []string                     `protobuf:"bytes,11,rep,name=additional_matches,json=additionalMatches,proto3" json:"additional_matches,omitempty"`
	Config            *v1.Tournament_Configuration `protobuf:"bytes,12,opt,name=config,proto3" json:"config,omitempty"`
	Admins            []*Tournament_Admin          `protobuf:"bytes,13,rep,name=admins,proto3" json:"admins,omitempty"`
	// created_by contain the id of the user that create the team.
	CreatedBy string              `protobuf:"bytes,14,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	State     v1.Tournament_State `protobuf:"varint,15,opt,name=state,proto3,enum=common.v1.Tournament_State" json:"state,omitempty"`
	// Type of game for this tournament.
	Game        string              `protobuf:"bytes,16,opt,name=game,proto3" json:"game,omitempty"`
	MatchConfig *MatchConfiguration `protobuf:"bytes,17,opt,name=match_config,json=matchConfig,proto3" json:"match_config,omitempty"`
}

func (x *Tournament) Reset() {
	*x = Tournament{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_tournament_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tournament) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tournament) ProtoMessage() {}

func (x *Tournament) ProtoReflect() protoreflect.Message {
	mi := &file_storage_tournament_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tournament.ProtoReflect.Descriptor instead.
func (*Tournament) Descriptor() ([]byte, []int) {
	return file_storage_tournament_proto_rawDescGZIP(), []int{0}
}

func (x *Tournament) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tournament) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tournament) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Tournament) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Tournament) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Tournament) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Tournament) GetParticipantsMode() v1.Tournament_ParticipantsMode {
	if x != nil {
		return x.ParticipantsMode
	}
	return v1.Tournament_ParticipantsMode(0)
}

func (m *Tournament) GetParticipants() isTournament_Participants {
	if m != nil {
		return m.Participants
	}
	return nil
}

func (x *Tournament) GetTeams() *v1.Tournament_TeamsParticipants {
	if x, ok := x.GetParticipants().(*Tournament_Teams); ok {
		return x.Teams
	}
	return nil
}

func (x *Tournament) GetPlayers() *v1.Tournament_PlayersParticipants {
	if x, ok := x.GetParticipants().(*Tournament_Players); ok {
		return x.Players
	}
	return nil
}

func (x *Tournament) GetPhases() []*v1.Tournament_Phase {
	if x != nil {
		return x.Phases
	}
	return nil
}

func (x *Tournament) GetAdditionalMatches() []string {
	if x != nil {
		return x.AdditionalMatches
	}
	return nil
}

func (x *Tournament) GetConfig() *v1.Tournament_Configuration {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Tournament) GetAdmins() []*Tournament_Admin {
	if x != nil {
		return x.Admins
	}
	return nil
}

func (x *Tournament) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Tournament) GetState() v1.Tournament_State {
	if x != nil {
		return x.State
	}
	return v1.Tournament_State(0)
}

func (x *Tournament) GetGame() string {
	if x != nil {
		return x.Game
	}
	return ""
}

func (x *Tournament) GetMatchConfig() *MatchConfiguration {
	if x != nil {
		return x.MatchConfig
	}
	return nil
}

type isTournament_Participants interface {
	isTournament_Participants()
}

type Tournament_Teams struct {
	Teams *v1.Tournament_TeamsParticipants `protobuf:"bytes,8,opt,name=teams,proto3,oneof"`
}

type Tournament_Players struct {
	Players *v1.Tournament_PlayersParticipants `protobuf:"bytes,9,opt,name=players,proto3,oneof"`
}

func (*Tournament_Teams) isTournament_Participants() {}

func (*Tournament_Players) isTournament_Participants() {}

// Match configuration.
type MatchConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// setup endpoint for the match.
	Setup *v1.Url `protobuf:"bytes,1,opt,name=setup,proto3" json:"setup,omitempty"`
	// cancel endpoint for the match.
	Cancel *v1.Url `protobuf:"bytes,2,opt,name=cancel,proto3" json:"cancel,omitempty"`
	// setup responses
	SetupResponses []*structpb.Struct `protobuf:"bytes,3,rep,name=setup_responses,json=setupResponses,proto3" json:"setup_responses,omitempty"`
}

func (x *MatchConfiguration) Reset() {
	*x = MatchConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_tournament_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchConfiguration) ProtoMessage() {}

func (x *MatchConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_storage_tournament_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchConfiguration.ProtoReflect.Descriptor instead.
func (*MatchConfiguration) Descriptor() ([]byte, []int) {
	return file_storage_tournament_proto_rawDescGZIP(), []int{1}
}

func (x *MatchConfiguration) GetSetup() *v1.Url {
	if x != nil {
		return x.Setup
	}
	return nil
}

func (x *MatchConfiguration) GetCancel() *v1.Url {
	if x != nil {
		return x.Cancel
	}
	return nil
}

func (x *MatchConfiguration) GetSetupResponses() []*structpb.Struct {
	if x != nil {
		return x.SetupResponses
	}
	return nil
}

// TournamentParticipantTeam index teams with tournaments
type TournamentParticipantTeam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tournaments []*TournamentParticipantTeam_Tournament `protobuf:"bytes,2,rep,name=tournaments,proto3" json:"tournaments,omitempty"`
	Version     uint64                                  `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	UpdateTime  *timestamppb.Timestamp                  `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *TournamentParticipantTeam) Reset() {
	*x = TournamentParticipantTeam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_tournament_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TournamentParticipantTeam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TournamentParticipantTeam) ProtoMessage() {}

func (x *TournamentParticipantTeam) ProtoReflect() protoreflect.Message {
	mi := &file_storage_tournament_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TournamentParticipantTeam.ProtoReflect.Descriptor instead.
func (*TournamentParticipantTeam) Descriptor() ([]byte, []int) {
	return file_storage_tournament_proto_rawDescGZIP(), []int{2}
}

func (x *TournamentParticipantTeam) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TournamentParticipantTeam) GetTournaments() []*TournamentParticipantTeam_Tournament {
	if x != nil {
		return x.Tournaments
	}
	return nil
}

func (x *TournamentParticipantTeam) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *TournamentParticipantTeam) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// TournamentParticipantPlayer index players with tournaments
type TournamentParticipantPlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tournaments []*TournamentParticipantPlayer_Tournament `protobuf:"bytes,2,rep,name=tournaments,proto3" json:"tournaments,omitempty"`
	Version     uint64                                    `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	UpdateTime  *timestamppb.Timestamp                    `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *TournamentParticipantPlayer) Reset() {
	*x = TournamentParticipantPlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_tournament_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TournamentParticipantPlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TournamentParticipantPlayer) ProtoMessage() {}

func (x *TournamentParticipantPlayer) ProtoReflect() protoreflect.Message {
	mi := &file_storage_tournament_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TournamentParticipantPlayer.ProtoReflect.Descriptor instead.
func (*TournamentParticipantPlayer) Descriptor() ([]byte, []int) {
	return file_storage_tournament_proto_rawDescGZIP(), []int{3}
}

func (x *TournamentParticipantPlayer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TournamentParticipantPlayer) GetTournaments() []*TournamentParticipantPlayer_Tournament {
	if x != nil {
		return x.Tournaments
	}
	return nil
}

func (x *TournamentParticipantPlayer) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *TournamentParticipantPlayer) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// admins contains the ids of the users that can manage the tournament.
type Tournament_Admin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // TODO: add role inside the tournament for admins.
}

func (x *Tournament_Admin) Reset() {
	*x = Tournament_Admin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_tournament_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tournament_Admin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tournament_Admin) ProtoMessage() {}

func (x *Tournament_Admin) ProtoReflect() protoreflect.Message {
	mi := &file_storage_tournament_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tournament_Admin.ProtoReflect.Descriptor instead.
func (*Tournament_Admin) Descriptor() ([]byte, []int) {
	return file_storage_tournament_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Tournament_Admin) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TournamentParticipantTeam_Tournament struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TournamentParticipantTeam_Tournament) Reset() {
	*x = TournamentParticipantTeam_Tournament{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_tournament_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TournamentParticipantTeam_Tournament) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TournamentParticipantTeam_Tournament) ProtoMessage() {}

func (x *TournamentParticipantTeam_Tournament) ProtoReflect() protoreflect.Message {
	mi := &file_storage_tournament_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TournamentParticipantTeam_Tournament.ProtoReflect.Descriptor instead.
func (*TournamentParticipantTeam_Tournament) Descriptor() ([]byte, []int) {
	return file_storage_tournament_proto_rawDescGZIP(), []int{2, 0}
}

func (x *TournamentParticipantTeam_Tournament) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TournamentParticipantPlayer_Tournament struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TournamentParticipantPlayer_Tournament) Reset() {
	*x = TournamentParticipantPlayer_Tournament{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_tournament_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TournamentParticipantPlayer_Tournament) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TournamentParticipantPlayer_Tournament) ProtoMessage() {}

func (x *TournamentParticipantPlayer_Tournament) ProtoReflect() protoreflect.Message {
	mi := &file_storage_tournament_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TournamentParticipantPlayer_Tournament.ProtoReflect.Descriptor instead.
func (*TournamentParticipantPlayer_Tournament) Descriptor() ([]byte, []int) {
	return file_storage_tournament_proto_rawDescGZIP(), []int{3, 0}
}

func (x *TournamentParticipantPlayer_Tournament) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_storage_tournament_proto protoreflect.FileDescriptor

var file_storage_tournament_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x72, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x06, 0x0a, 0x0a, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x11,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x52,
	0x10, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x3f, 0x0a, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x75,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x48, 0x00, 0x52, 0x05, 0x74, 0x65, 0x61,
	0x6d, 0x73, 0x12, 0x45, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x48, 0x00,
	0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x70, 0x68, 0x61,
	0x73, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x50, 0x68, 0x61, 0x73, 0x65, 0x52, 0x06, 0x70, 0x68, 0x61, 0x73, 0x65, 0x73, 0x12, 0x2d,
	0x0a, 0x12, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x61, 0x64, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x3b, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x31, 0x0a, 0x06, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x06, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x31, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67,
	0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x1a, 0x17, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x0e, 0x0a, 0x0c,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x22, 0xa4, 0x01, 0x0a,
	0x12, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x05, 0x73, 0x65, 0x74, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x72, 0x6c, 0x52, 0x05, 0x73, 0x65, 0x74, 0x75, 0x70, 0x12, 0x26, 0x0a, 0x06, 0x63, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x72, 0x6c, 0x52, 0x06, 0x63, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x12, 0x40, 0x0a, 0x0f, 0x73, 0x65, 0x74, 0x75, 0x70, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x0e, 0x73, 0x65, 0x74, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x73, 0x22, 0xf1, 0x01, 0x0a, 0x19, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x6e, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x54, 0x65, 0x61,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x4f, 0x0a, 0x0b, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x1c, 0x0a, 0x0a, 0x54, 0x6f, 0x75,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xf5, 0x01, 0x0a, 0x1b, 0x54, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x51, 0x0a, 0x0b, 0x74, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x74,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x1a, 0x1c, 0x0a, 0x0a, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42,
	0x8e, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42,
	0x0f, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x61, 0x6d, 0x6f, 0x6e, 0x62, 0x65, 0x72, 0x72, 0x75, 0x74, 0x74, 0x69, 0x2f, 0x6a, 0x75, 0x6d,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0xca, 0x02, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0xe2, 0x02, 0x13, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_tournament_proto_rawDescOnce sync.Once
	file_storage_tournament_proto_rawDescData = file_storage_tournament_proto_rawDesc
)

func file_storage_tournament_proto_rawDescGZIP() []byte {
	file_storage_tournament_proto_rawDescOnce.Do(func() {
		file_storage_tournament_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_tournament_proto_rawDescData)
	})
	return file_storage_tournament_proto_rawDescData
}

var file_storage_tournament_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_storage_tournament_proto_goTypes = []interface{}{
	(*Tournament)(nil),                             // 0: storage.Tournament
	(*MatchConfiguration)(nil),                     // 1: storage.MatchConfiguration
	(*TournamentParticipantTeam)(nil),              // 2: storage.TournamentParticipantTeam
	(*TournamentParticipantPlayer)(nil),            // 3: storage.TournamentParticipantPlayer
	(*Tournament_Admin)(nil),                       // 4: storage.Tournament.Admin
	(*TournamentParticipantTeam_Tournament)(nil),   // 5: storage.TournamentParticipantTeam.Tournament
	(*TournamentParticipantPlayer_Tournament)(nil), // 6: storage.TournamentParticipantPlayer.Tournament
	(*timestamppb.Timestamp)(nil),                  // 7: google.protobuf.Timestamp
	(v1.Tournament_ParticipantsMode)(0),            // 8: common.v1.Tournament.ParticipantsMode
	(*v1.Tournament_TeamsParticipants)(nil),        // 9: common.v1.Tournament.TeamsParticipants
	(*v1.Tournament_PlayersParticipants)(nil),      // 10: common.v1.Tournament.PlayersParticipants
	(*v1.Tournament_Phase)(nil),                    // 11: common.v1.Tournament.Phase
	(*v1.Tournament_Configuration)(nil),            // 12: common.v1.Tournament.Configuration
	(v1.Tournament_State)(0),                       // 13: common.v1.Tournament.State
	(*v1.Url)(nil),                                 // 14: common.v1.Url
	(*structpb.Struct)(nil),                        // 15: google.protobuf.Struct
}
var file_storage_tournament_proto_depIdxs = []int32{
	7,  // 0: storage.Tournament.create_time:type_name -> google.protobuf.Timestamp
	7,  // 1: storage.Tournament.update_time:type_name -> google.protobuf.Timestamp
	8,  // 2: storage.Tournament.participants_mode:type_name -> common.v1.Tournament.ParticipantsMode
	9,  // 3: storage.Tournament.teams:type_name -> common.v1.Tournament.TeamsParticipants
	10, // 4: storage.Tournament.players:type_name -> common.v1.Tournament.PlayersParticipants
	11, // 5: storage.Tournament.phases:type_name -> common.v1.Tournament.Phase
	12, // 6: storage.Tournament.config:type_name -> common.v1.Tournament.Configuration
	4,  // 7: storage.Tournament.admins:type_name -> storage.Tournament.Admin
	13, // 8: storage.Tournament.state:type_name -> common.v1.Tournament.State
	1,  // 9: storage.Tournament.match_config:type_name -> storage.MatchConfiguration
	14, // 10: storage.MatchConfiguration.setup:type_name -> common.v1.Url
	14, // 11: storage.MatchConfiguration.cancel:type_name -> common.v1.Url
	15, // 12: storage.MatchConfiguration.setup_responses:type_name -> google.protobuf.Struct
	5,  // 13: storage.TournamentParticipantTeam.tournaments:type_name -> storage.TournamentParticipantTeam.Tournament
	7,  // 14: storage.TournamentParticipantTeam.update_time:type_name -> google.protobuf.Timestamp
	6,  // 15: storage.TournamentParticipantPlayer.tournaments:type_name -> storage.TournamentParticipantPlayer.Tournament
	7,  // 16: storage.TournamentParticipantPlayer.update_time:type_name -> google.protobuf.Timestamp
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_storage_tournament_proto_init() }
func file_storage_tournament_proto_init() {
	if File_storage_tournament_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storage_tournament_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tournament); i {
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
		file_storage_tournament_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchConfiguration); i {
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
		file_storage_tournament_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TournamentParticipantTeam); i {
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
		file_storage_tournament_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TournamentParticipantPlayer); i {
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
		file_storage_tournament_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tournament_Admin); i {
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
		file_storage_tournament_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TournamentParticipantTeam_Tournament); i {
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
		file_storage_tournament_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TournamentParticipantPlayer_Tournament); i {
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
	file_storage_tournament_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Tournament_Teams)(nil),
		(*Tournament_Players)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_tournament_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_tournament_proto_goTypes,
		DependencyIndexes: file_storage_tournament_proto_depIdxs,
		MessageInfos:      file_storage_tournament_proto_msgTypes,
	}.Build()
	File_storage_tournament_proto = out.File
	file_storage_tournament_proto_rawDesc = nil
	file_storage_tournament_proto_goTypes = nil
	file_storage_tournament_proto_depIdxs = nil
}