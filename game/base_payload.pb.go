// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.1
// source: base_payload.proto

package game

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

type TeamColor int32

const (
	TeamColor_TEAM_UNSPECIFIED TeamColor = 0
	TeamColor_TEAM_RED         TeamColor = 1
	TeamColor_TEAM_BLACK       TeamColor = 2
)

// Enum value maps for TeamColor.
var (
	TeamColor_name = map[int32]string{
		0: "TEAM_UNSPECIFIED",
		1: "TEAM_RED",
		2: "TEAM_BLACK",
	}
	TeamColor_value = map[string]int32{
		"TEAM_UNSPECIFIED": 0,
		"TEAM_RED":         1,
		"TEAM_BLACK":       2,
	}
)

func (x TeamColor) Enum() *TeamColor {
	p := new(TeamColor)
	*p = x
	return p
}

func (x TeamColor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TeamColor) Descriptor() protoreflect.EnumDescriptor {
	return file_base_payload_proto_enumTypes[0].Descriptor()
}

func (TeamColor) Type() protoreflect.EnumType {
	return &file_base_payload_proto_enumTypes[0]
}

func (x TeamColor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TeamColor.Descriptor instead.
func (TeamColor) EnumDescriptor() ([]byte, []int) {
	return file_base_payload_proto_rawDescGZIP(), []int{0}
}

type BasePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// GUI display text from server to client
	Notice string `protobuf:"bytes,2,opt,name=notice,proto3" json:"notice,omitempty"`
	// Types that are assignable to Inner:
	//
	//	*BasePayload_Welcome
	//	*BasePayload_MovePayload
	//	*BasePayload_Start
	//	*BasePayload_ExitPayload
	//	*BasePayload_CapturePayload
	Inner isBasePayload_Inner `protobuf_oneof:"inner"`
}

func (x *BasePayload) Reset() {
	*x = BasePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasePayload) ProtoMessage() {}

func (x *BasePayload) ProtoReflect() protoreflect.Message {
	mi := &file_base_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasePayload.ProtoReflect.Descriptor instead.
func (*BasePayload) Descriptor() ([]byte, []int) {
	return file_base_payload_proto_rawDescGZIP(), []int{0}
}

func (x *BasePayload) GetNotice() string {
	if x != nil {
		return x.Notice
	}
	return ""
}

func (m *BasePayload) GetInner() isBasePayload_Inner {
	if m != nil {
		return m.Inner
	}
	return nil
}

func (x *BasePayload) GetWelcome() *WelcomePayload {
	if x, ok := x.GetInner().(*BasePayload_Welcome); ok {
		return x.Welcome
	}
	return nil
}

func (x *BasePayload) GetMovePayload() *MovePayload {
	if x, ok := x.GetInner().(*BasePayload_MovePayload); ok {
		return x.MovePayload
	}
	return nil
}

func (x *BasePayload) GetStart() *StartPayload {
	if x, ok := x.GetInner().(*BasePayload_Start); ok {
		return x.Start
	}
	return nil
}

func (x *BasePayload) GetExitPayload() *ExitPayload {
	if x, ok := x.GetInner().(*BasePayload_ExitPayload); ok {
		return x.ExitPayload
	}
	return nil
}

func (x *BasePayload) GetCapturePayload() *CapturePayload {
	if x, ok := x.GetInner().(*BasePayload_CapturePayload); ok {
		return x.CapturePayload
	}
	return nil
}

type isBasePayload_Inner interface {
	isBasePayload_Inner()
}

type BasePayload_Welcome struct {
	Welcome *WelcomePayload `protobuf:"bytes,3,opt,name=welcome,proto3,oneof"`
}

type BasePayload_MovePayload struct {
	MovePayload *MovePayload `protobuf:"bytes,4,opt,name=move_payload,json=movePayload,proto3,oneof"`
}

type BasePayload_Start struct {
	Start *StartPayload `protobuf:"bytes,5,opt,name=start,proto3,oneof"`
}

type BasePayload_ExitPayload struct {
	ExitPayload *ExitPayload `protobuf:"bytes,6,opt,name=exit_payload,json=exitPayload,proto3,oneof"`
}

type BasePayload_CapturePayload struct {
	CapturePayload *CapturePayload `protobuf:"bytes,7,opt,name=capture_payload,json=capturePayload,proto3,oneof"`
}

func (*BasePayload_Welcome) isBasePayload_Inner() {}

func (*BasePayload_MovePayload) isBasePayload_Inner() {}

func (*BasePayload_Start) isBasePayload_Inner() {}

func (*BasePayload_ExitPayload) isBasePayload_Inner() {}

func (*BasePayload_CapturePayload) isBasePayload_Inner() {}

type WelcomePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Color of this player pieces
	MyTeam TeamColor `protobuf:"varint,1,opt,name=my_team,json=myTeam,proto3,enum=chk.payload.TeamColor" json:"my_team,omitempty"`
	// in the form of X.Y.Z
	ServerVersion string `protobuf:"bytes,2,opt,name=server_version,json=serverVersion,proto3" json:"server_version,omitempty"`
}

func (x *WelcomePayload) Reset() {
	*x = WelcomePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_payload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WelcomePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WelcomePayload) ProtoMessage() {}

func (x *WelcomePayload) ProtoReflect() protoreflect.Message {
	mi := &file_base_payload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WelcomePayload.ProtoReflect.Descriptor instead.
func (*WelcomePayload) Descriptor() ([]byte, []int) {
	return file_base_payload_proto_rawDescGZIP(), []int{1}
}

func (x *WelcomePayload) GetMyTeam() TeamColor {
	if x != nil {
		return x.MyTeam
	}
	return TeamColor_TEAM_UNSPECIFIED
}

func (x *WelcomePayload) GetServerVersion() string {
	if x != nil {
		return x.ServerVersion
	}
	return ""
}

type StartPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Pieces for red player
	PiecesRed []int32 `protobuf:"varint,1,rep,packed,name=pieces_red,json=piecesRed,proto3" json:"pieces_red,omitempty"`
	// pieces for black player
	PiecesBlack []int32 `protobuf:"varint,2,rep,packed,name=pieces_black,json=piecesBlack,proto3" json:"pieces_black,omitempty"`
}

func (x *StartPayload) Reset() {
	*x = StartPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_payload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartPayload) ProtoMessage() {}

func (x *StartPayload) ProtoReflect() protoreflect.Message {
	mi := &file_base_payload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartPayload.ProtoReflect.Descriptor instead.
func (*StartPayload) Descriptor() ([]byte, []int) {
	return file_base_payload_proto_rawDescGZIP(), []int{2}
}

func (x *StartPayload) GetPiecesRed() []int32 {
	if x != nil {
		return x.PiecesRed
	}
	return nil
}

func (x *StartPayload) GetPiecesBlack() []int32 {
	if x != nil {
		return x.PiecesBlack
	}
	return nil
}

type MovePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// from which player
	FromTeam TeamColor `protobuf:"varint,1,opt,name=from_team,json=fromTeam,proto3,enum=chk.payload.TeamColor" json:"from_team,omitempty"`
	// moving piece id
	PieceId int32 `protobuf:"varint,2,opt,name=piece_id,json=pieceId,proto3" json:"piece_id,omitempty"`
	// older cell index for this pieceId
	SourceCell int32 `protobuf:"varint,3,opt,name=source_cell,json=sourceCell,proto3" json:"source_cell,omitempty"`
	// destination cell
	DestCell *MovePayload_DestCell `protobuf:"bytes,4,opt,name=dest_cell,json=destCell,proto3" json:"dest_cell,omitempty"`
}

func (x *MovePayload) Reset() {
	*x = MovePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_payload_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovePayload) ProtoMessage() {}

func (x *MovePayload) ProtoReflect() protoreflect.Message {
	mi := &file_base_payload_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovePayload.ProtoReflect.Descriptor instead.
func (*MovePayload) Descriptor() ([]byte, []int) {
	return file_base_payload_proto_rawDescGZIP(), []int{3}
}

func (x *MovePayload) GetFromTeam() TeamColor {
	if x != nil {
		return x.FromTeam
	}
	return TeamColor_TEAM_UNSPECIFIED
}

func (x *MovePayload) GetPieceId() int32 {
	if x != nil {
		return x.PieceId
	}
	return 0
}

func (x *MovePayload) GetSourceCell() int32 {
	if x != nil {
		return x.SourceCell
	}
	return 0
}

func (x *MovePayload) GetDestCell() *MovePayload_DestCell {
	if x != nil {
		return x.DestCell
	}
	return nil
}

// whenever any of the players exits, or server terminates Match
type ExitPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// which player initiated exit (RED, BLACK, or unspecified)
	FromTeam TeamColor `protobuf:"varint,1,opt,name=from_team,json=fromTeam,proto3,enum=chk.payload.TeamColor" json:"from_team,omitempty"`
}

func (x *ExitPayload) Reset() {
	*x = ExitPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_payload_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExitPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExitPayload) ProtoMessage() {}

func (x *ExitPayload) ProtoReflect() protoreflect.Message {
	mi := &file_base_payload_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExitPayload.ProtoReflect.Descriptor instead.
func (*ExitPayload) Descriptor() ([]byte, []int) {
	return file_base_payload_proto_rawDescGZIP(), []int{4}
}

func (x *ExitPayload) GetFromTeam() TeamColor {
	if x != nil {
		return x.FromTeam
	}
	return TeamColor_TEAM_UNSPECIFIED
}

// when hunter player is capturing opponent's piece
type CapturePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// from which player
	FromTeam TeamColor `protobuf:"varint,1,opt,name=from_team,json=fromTeam,proto3,enum=chk.payload.TeamColor" json:"from_team,omitempty"`
	// the attacking player's piece
	HunterPieceId int32 `protobuf:"varint,2,opt,name=hunter_piece_id,json=hunterPieceId,proto3" json:"hunter_piece_id,omitempty"`
	// source cell of hunter
	HunterCell int32                         `protobuf:"varint,3,opt,name=hunter_cell,json=hunterCell,proto3" json:"hunter_cell,omitempty"`
	Details    *CapturePayload_TargetDetails `protobuf:"bytes,4,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *CapturePayload) Reset() {
	*x = CapturePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_payload_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapturePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapturePayload) ProtoMessage() {}

func (x *CapturePayload) ProtoReflect() protoreflect.Message {
	mi := &file_base_payload_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapturePayload.ProtoReflect.Descriptor instead.
func (*CapturePayload) Descriptor() ([]byte, []int) {
	return file_base_payload_proto_rawDescGZIP(), []int{5}
}

func (x *CapturePayload) GetFromTeam() TeamColor {
	if x != nil {
		return x.FromTeam
	}
	return TeamColor_TEAM_UNSPECIFIED
}

func (x *CapturePayload) GetHunterPieceId() int32 {
	if x != nil {
		return x.HunterPieceId
	}
	return 0
}

func (x *CapturePayload) GetHunterCell() int32 {
	if x != nil {
		return x.HunterCell
	}
	return 0
}

func (x *CapturePayload) GetDetails() *CapturePayload_TargetDetails {
	if x != nil {
		return x.Details
	}
	return nil
}

// Represents destination cell
type MovePayload_DestCell struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CellIndex int32   `protobuf:"varint,1,opt,name=cell_index,json=cellIndex,proto3" json:"cell_index,omitempty"`
	X         float32 `protobuf:"fixed32,2,opt,name=x,proto3" json:"x,omitempty"`
	Y         float32 `protobuf:"fixed32,3,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *MovePayload_DestCell) Reset() {
	*x = MovePayload_DestCell{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_payload_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovePayload_DestCell) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovePayload_DestCell) ProtoMessage() {}

func (x *MovePayload_DestCell) ProtoReflect() protoreflect.Message {
	mi := &file_base_payload_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovePayload_DestCell.ProtoReflect.Descriptor instead.
func (*MovePayload_DestCell) Descriptor() ([]byte, []int) {
	return file_base_payload_proto_rawDescGZIP(), []int{3, 0}
}

func (x *MovePayload_DestCell) GetCellIndex() int32 {
	if x != nil {
		return x.CellIndex
	}
	return 0
}

func (x *MovePayload_DestCell) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *MovePayload_DestCell) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

// extra details about this capture
type CapturePayload_TargetDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// * ID of the piece that is to be captured
	PreyPieceId int32 `protobuf:"varint,1,opt,name=prey_piece_id,json=preyPieceId,proto3" json:"prey_piece_id,omitempty"`
	// the cell Index hosting this target piece
	PreyCellIdx int32 `protobuf:"varint,2,opt,name=prey_cell_idx,json=preyCellIdx,proto3" json:"prey_cell_idx,omitempty"`
	// Destination of hunterPiece after capturing prey
	HunterNextCell int32 `protobuf:"varint,3,opt,name=hunter_next_cell,json=hunterNextCell,proto3" json:"hunter_next_cell,omitempty"`
}

func (x *CapturePayload_TargetDetails) Reset() {
	*x = CapturePayload_TargetDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_payload_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapturePayload_TargetDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapturePayload_TargetDetails) ProtoMessage() {}

func (x *CapturePayload_TargetDetails) ProtoReflect() protoreflect.Message {
	mi := &file_base_payload_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapturePayload_TargetDetails.ProtoReflect.Descriptor instead.
func (*CapturePayload_TargetDetails) Descriptor() ([]byte, []int) {
	return file_base_payload_proto_rawDescGZIP(), []int{5, 0}
}

func (x *CapturePayload_TargetDetails) GetPreyPieceId() int32 {
	if x != nil {
		return x.PreyPieceId
	}
	return 0
}

func (x *CapturePayload_TargetDetails) GetPreyCellIdx() int32 {
	if x != nil {
		return x.PreyCellIdx
	}
	return 0
}

func (x *CapturePayload_TargetDetails) GetHunterNextCell() int32 {
	if x != nil {
		return x.HunterNextCell
	}
	return 0
}

var File_base_payload_proto protoreflect.FileDescriptor

var file_base_payload_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x68, 0x6b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0xe0, 0x02, 0x0a, 0x0b, 0x42, 0x61, 0x73, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x77, 0x65, 0x6c,
	0x63, 0x6f, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x68, 0x6b,
	0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x00, 0x52, 0x07, 0x77, 0x65, 0x6c, 0x63, 0x6f,
	0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x68, 0x6b, 0x2e, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x48, 0x00, 0x52, 0x0b, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x31, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x63, 0x68, 0x6b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x00, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x68, 0x6b,
	0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x45, 0x78, 0x69, 0x74, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x48, 0x00, 0x52, 0x0b, 0x65, 0x78, 0x69, 0x74, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x46, 0x0a, 0x0f, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63,
	0x68, 0x6b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x75,
	0x72, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x61, 0x70,
	0x74, 0x75, 0x72, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x69,
	0x6e, 0x6e, 0x65, 0x72, 0x22, 0x68, 0x0a, 0x0e, 0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2f, 0x0a, 0x07, 0x6d, 0x79, 0x5f, 0x74, 0x65, 0x61,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x68, 0x6b, 0x2e, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52,
	0x06, 0x6d, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x50,
	0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x69, 0x65, 0x63, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x09, 0x70, 0x69, 0x65, 0x63, 0x65, 0x73, 0x52, 0x65, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x69, 0x65, 0x63, 0x65, 0x73, 0x5f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x69, 0x65, 0x63, 0x65, 0x73, 0x42, 0x6c, 0x61, 0x63, 0x6b,
	0x22, 0x85, 0x02, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x33, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x68, 0x6b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x08, 0x66, 0x72, 0x6f,
	0x6d, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x69, 0x65, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x69, 0x65, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x65, 0x6c, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x65, 0x6c,
	0x6c, 0x12, 0x3e, 0x0a, 0x09, 0x64, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x65, 0x6c, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x68, 0x6b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x44,
	0x65, 0x73, 0x74, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x08, 0x64, 0x65, 0x73, 0x74, 0x43, 0x65, 0x6c,
	0x6c, 0x1a, 0x45, 0x0a, 0x08, 0x44, 0x65, 0x73, 0x74, 0x43, 0x65, 0x6c, 0x6c, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x63, 0x65, 0x6c, 0x6c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x0c, 0x0a, 0x01,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x22, 0x42, 0x0a, 0x0b, 0x45, 0x78, 0x69, 0x74,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x33, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f,
	0x74, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x68, 0x6b,
	0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x61, 0x6d, 0x22, 0xd7, 0x02, 0x0a,
	0x0e, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x33, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x68, 0x6b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x2e, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d,
	0x54, 0x65, 0x61, 0x6d, 0x12, 0x26, 0x0a, 0x0f, 0x68, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x70,
	0x69, 0x65, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x68,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x69, 0x65, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x68, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x68, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x65, 0x6c, 0x6c, 0x12, 0x43, 0x0a,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x63, 0x68, 0x6b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x43, 0x61, 0x70,
	0x74, 0x75, 0x72, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x1a, 0x81, 0x01, 0x0a, 0x0d, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x79, 0x5f, 0x70, 0x69, 0x65,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x72, 0x65,
	0x79, 0x50, 0x69, 0x65, 0x63, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x79,
	0x5f, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x70, 0x72, 0x65, 0x79, 0x43, 0x65, 0x6c, 0x6c, 0x49, 0x64, 0x78, 0x12, 0x28, 0x0a, 0x10,
	0x68, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x65, 0x6c, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x68, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x4e, 0x65,
	0x78, 0x74, 0x43, 0x65, 0x6c, 0x6c, 0x2a, 0x3f, 0x0a, 0x09, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x6f,
	0x6c, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x45, 0x41,
	0x4d, 0x5f, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x45, 0x41, 0x4d, 0x5f,
	0x42, 0x4c, 0x41, 0x43, 0x4b, 0x10, 0x02, 0x42, 0x17, 0x5a, 0x15, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x65, 0x72, 0x73, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x61, 0x6d, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_base_payload_proto_rawDescOnce sync.Once
	file_base_payload_proto_rawDescData = file_base_payload_proto_rawDesc
)

func file_base_payload_proto_rawDescGZIP() []byte {
	file_base_payload_proto_rawDescOnce.Do(func() {
		file_base_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_payload_proto_rawDescData)
	})
	return file_base_payload_proto_rawDescData
}

var file_base_payload_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_base_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_base_payload_proto_goTypes = []interface{}{
	(TeamColor)(0),                       // 0: chk.payload.TeamColor
	(*BasePayload)(nil),                  // 1: chk.payload.BasePayload
	(*WelcomePayload)(nil),               // 2: chk.payload.WelcomePayload
	(*StartPayload)(nil),                 // 3: chk.payload.StartPayload
	(*MovePayload)(nil),                  // 4: chk.payload.MovePayload
	(*ExitPayload)(nil),                  // 5: chk.payload.ExitPayload
	(*CapturePayload)(nil),               // 6: chk.payload.CapturePayload
	(*MovePayload_DestCell)(nil),         // 7: chk.payload.MovePayload.DestCell
	(*CapturePayload_TargetDetails)(nil), // 8: chk.payload.CapturePayload.TargetDetails
}
var file_base_payload_proto_depIdxs = []int32{
	2,  // 0: chk.payload.BasePayload.welcome:type_name -> chk.payload.WelcomePayload
	4,  // 1: chk.payload.BasePayload.move_payload:type_name -> chk.payload.MovePayload
	3,  // 2: chk.payload.BasePayload.start:type_name -> chk.payload.StartPayload
	5,  // 3: chk.payload.BasePayload.exit_payload:type_name -> chk.payload.ExitPayload
	6,  // 4: chk.payload.BasePayload.capture_payload:type_name -> chk.payload.CapturePayload
	0,  // 5: chk.payload.WelcomePayload.my_team:type_name -> chk.payload.TeamColor
	0,  // 6: chk.payload.MovePayload.from_team:type_name -> chk.payload.TeamColor
	7,  // 7: chk.payload.MovePayload.dest_cell:type_name -> chk.payload.MovePayload.DestCell
	0,  // 8: chk.payload.ExitPayload.from_team:type_name -> chk.payload.TeamColor
	0,  // 9: chk.payload.CapturePayload.from_team:type_name -> chk.payload.TeamColor
	8,  // 10: chk.payload.CapturePayload.details:type_name -> chk.payload.CapturePayload.TargetDetails
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_base_payload_proto_init() }
func file_base_payload_proto_init() {
	if File_base_payload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_payload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasePayload); i {
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
		file_base_payload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WelcomePayload); i {
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
		file_base_payload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartPayload); i {
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
		file_base_payload_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovePayload); i {
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
		file_base_payload_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExitPayload); i {
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
		file_base_payload_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapturePayload); i {
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
		file_base_payload_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovePayload_DestCell); i {
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
		file_base_payload_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapturePayload_TargetDetails); i {
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
	file_base_payload_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*BasePayload_Welcome)(nil),
		(*BasePayload_MovePayload)(nil),
		(*BasePayload_Start)(nil),
		(*BasePayload_ExitPayload)(nil),
		(*BasePayload_CapturePayload)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_base_payload_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_payload_proto_goTypes,
		DependencyIndexes: file_base_payload_proto_depIdxs,
		EnumInfos:         file_base_payload_proto_enumTypes,
		MessageInfos:      file_base_payload_proto_msgTypes,
	}.Build()
	File_base_payload_proto = out.File
	file_base_payload_proto_rawDesc = nil
	file_base_payload_proto_goTypes = nil
	file_base_payload_proto_depIdxs = nil
}
