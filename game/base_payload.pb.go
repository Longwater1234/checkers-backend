// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.1
// source: base_payload.proto

package game

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// GUI display text from server to client
	Notice string `protobuf:"bytes,2,opt,name=notice,proto3" json:"notice,omitempty"`
	// Types that are valid to be assigned to Inner:
	//
	//	*BasePayload_Welcome
	//	*BasePayload_MovePayload
	//	*BasePayload_Start
	//	*BasePayload_ExitPayload
	//	*BasePayload_CapturePayload
	//	*BasePayload_WinlosePayload
	Inner         isBasePayload_Inner `protobuf_oneof:"inner"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BasePayload) Reset() {
	*x = BasePayload{}
	mi := &file_base_payload_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BasePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasePayload) ProtoMessage() {}

func (x *BasePayload) ProtoReflect() protoreflect.Message {
	mi := &file_base_payload_proto_msgTypes[0]
	if x != nil {
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

func (x *BasePayload) GetInner() isBasePayload_Inner {
	if x != nil {
		return x.Inner
	}
	return nil
}

func (x *BasePayload) GetWelcome() *WelcomePayload {
	if x != nil {
		if x, ok := x.Inner.(*BasePayload_Welcome); ok {
			return x.Welcome
		}
	}
	return nil
}

func (x *BasePayload) GetMovePayload() *MovePayload {
	if x != nil {
		if x, ok := x.Inner.(*BasePayload_MovePayload); ok {
			return x.MovePayload
		}
	}
	return nil
}

func (x *BasePayload) GetStart() *StartPayload {
	if x != nil {
		if x, ok := x.Inner.(*BasePayload_Start); ok {
			return x.Start
		}
	}
	return nil
}

func (x *BasePayload) GetExitPayload() *ExitPayload {
	if x != nil {
		if x, ok := x.Inner.(*BasePayload_ExitPayload); ok {
			return x.ExitPayload
		}
	}
	return nil
}

func (x *BasePayload) GetCapturePayload() *CapturePayload {
	if x != nil {
		if x, ok := x.Inner.(*BasePayload_CapturePayload); ok {
			return x.CapturePayload
		}
	}
	return nil
}

func (x *BasePayload) GetWinlosePayload() *WinLosePayload {
	if x != nil {
		if x, ok := x.Inner.(*BasePayload_WinlosePayload); ok {
			return x.WinlosePayload
		}
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

type BasePayload_WinlosePayload struct {
	WinlosePayload *WinLosePayload `protobuf:"bytes,8,opt,name=winlose_payload,json=winlosePayload,proto3,oneof"`
}

func (*BasePayload_Welcome) isBasePayload_Inner() {}

func (*BasePayload_MovePayload) isBasePayload_Inner() {}

func (*BasePayload_Start) isBasePayload_Inner() {}

func (*BasePayload_ExitPayload) isBasePayload_Inner() {}

func (*BasePayload_CapturePayload) isBasePayload_Inner() {}

func (*BasePayload_WinlosePayload) isBasePayload_Inner() {}

type WelcomePayload struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Color of this player pieces
	MyTeam TeamColor `protobuf:"varint,1,opt,name=my_team,json=myTeam,proto3,enum=chk.payload.TeamColor" json:"my_team,omitempty"`
	// in the form of X.Y.Z
	ServerVersion string `protobuf:"bytes,2,opt,name=server_version,json=serverVersion,proto3" json:"server_version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WelcomePayload) Reset() {
	*x = WelcomePayload{}
	mi := &file_base_payload_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WelcomePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WelcomePayload) ProtoMessage() {}

func (x *WelcomePayload) ProtoReflect() protoreflect.Message {
	mi := &file_base_payload_proto_msgTypes[1]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// Pieces for red player
	PiecesRed []int32 `protobuf:"varint,1,rep,packed,name=pieces_red,json=piecesRed,proto3" json:"pieces_red,omitempty"`
	// pieces for black player
	PiecesBlack   []int32 `protobuf:"varint,2,rep,packed,name=pieces_black,json=piecesBlack,proto3" json:"pieces_black,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartPayload) Reset() {
	*x = StartPayload{}
	mi := &file_base_payload_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartPayload) ProtoMessage() {}

func (x *StartPayload) ProtoReflect() protoreflect.Message {
	mi := &file_base_payload_proto_msgTypes[2]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// from which player
	FromTeam TeamColor `protobuf:"varint,1,opt,name=from_team,json=fromTeam,proto3,enum=chk.payload.TeamColor" json:"from_team,omitempty"`
	// moving piece id
	PieceId int32 `protobuf:"varint,2,opt,name=piece_id,json=pieceId,proto3" json:"piece_id,omitempty"`
	// older cell index for this pieceId
	SourceCell int32 `protobuf:"varint,3,opt,name=source_cell,json=sourceCell,proto3" json:"source_cell,omitempty"`
	// destination cell
	Destination   *MovePayload_Detination `protobuf:"bytes,4,opt,name=destination,proto3" json:"destination,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MovePayload) Reset() {
	*x = MovePayload{}
	mi := &file_base_payload_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MovePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovePayload) ProtoMessage() {}

func (x *MovePayload) ProtoReflect() protoreflect.Message {
	mi := &file_base_payload_proto_msgTypes[3]
	if x != nil {
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

func (x *MovePayload) GetDestination() *MovePayload_Detination {
	if x != nil {
		return x.Destination
	}
	return nil
}

// whenever any of the players exits, or server terminates Match
type ExitPayload struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// which player initiated exit (RED, BLACK, or unspecified)
	FromTeam      TeamColor `protobuf:"varint,1,opt,name=from_team,json=fromTeam,proto3,enum=chk.payload.TeamColor" json:"from_team,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExitPayload) Reset() {
	*x = ExitPayload{}
	mi := &file_base_payload_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExitPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExitPayload) ProtoMessage() {}

func (x *ExitPayload) ProtoReflect() protoreflect.Message {
	mi := &file_base_payload_proto_msgTypes[4]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// from which player
	FromTeam TeamColor `protobuf:"varint,1,opt,name=from_team,json=fromTeam,proto3,enum=chk.payload.TeamColor" json:"from_team,omitempty"`
	// the attacking player's piece
	HunterPieceId int32                             `protobuf:"varint,2,opt,name=hunter_piece_id,json=hunterPieceId,proto3" json:"hunter_piece_id,omitempty"`
	Details       *CapturePayload_TargetDetails     `protobuf:"bytes,4,opt,name=details,proto3" json:"details,omitempty"`
	Destination   *CapturePayload_HunterDestination `protobuf:"bytes,5,opt,name=destination,proto3" json:"destination,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CapturePayload) Reset() {
	*x = CapturePayload{}
	mi := &file_base_payload_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CapturePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapturePayload) ProtoMessage() {}

func (x *CapturePayload) ProtoReflect() protoreflect.Message {
	mi := &file_base_payload_proto_msgTypes[5]
	if x != nil {
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

func (x *CapturePayload) GetDetails() *CapturePayload_TargetDetails {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *CapturePayload) GetDestination() *CapturePayload_HunterDestination {
	if x != nil {
		return x.Destination
	}
	return nil
}

// When one of player wins. Also, this marks the end of the match
type WinLosePayload struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// which player won
	Winner        TeamColor `protobuf:"varint,1,opt,name=winner,proto3,enum=chk.payload.TeamColor" json:"winner,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WinLosePayload) Reset() {
	*x = WinLosePayload{}
	mi := &file_base_payload_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WinLosePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinLosePayload) ProtoMessage() {}

func (x *WinLosePayload) ProtoReflect() protoreflect.Message {
	mi := &file_base_payload_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinLosePayload.ProtoReflect.Descriptor instead.
func (*WinLosePayload) Descriptor() ([]byte, []int) {
	return file_base_payload_proto_rawDescGZIP(), []int{6}
}

func (x *WinLosePayload) GetWinner() TeamColor {
	if x != nil {
		return x.Winner
	}
	return TeamColor_TEAM_UNSPECIFIED
}

// Where will this player land on
type MovePayload_Detination struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CellIndex     int32                  `protobuf:"varint,1,opt,name=cell_index,json=cellIndex,proto3" json:"cell_index,omitempty"`
	X             float32                `protobuf:"fixed32,2,opt,name=x,proto3" json:"x,omitempty"`
	Y             float32                `protobuf:"fixed32,3,opt,name=y,proto3" json:"y,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MovePayload_Detination) Reset() {
	*x = MovePayload_Detination{}
	mi := &file_base_payload_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MovePayload_Detination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovePayload_Detination) ProtoMessage() {}

func (x *MovePayload_Detination) ProtoReflect() protoreflect.Message {
	mi := &file_base_payload_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovePayload_Detination.ProtoReflect.Descriptor instead.
func (*MovePayload_Detination) Descriptor() ([]byte, []int) {
	return file_base_payload_proto_rawDescGZIP(), []int{3, 0}
}

func (x *MovePayload_Detination) GetCellIndex() int32 {
	if x != nil {
		return x.CellIndex
	}
	return 0
}

func (x *MovePayload_Detination) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *MovePayload_Detination) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

// extra details about this capture
type CapturePayload_TargetDetails struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the piece that is to be captured
	PreyPieceId int32 `protobuf:"varint,1,opt,name=prey_piece_id,json=preyPieceId,proto3" json:"prey_piece_id,omitempty"`
	// the cell Index hosting this target piece
	PreyCellIdx int32 `protobuf:"varint,2,opt,name=prey_cell_idx,json=preyCellIdx,proto3" json:"prey_cell_idx,omitempty"`
	// source cell of hunter
	HunterSrcCell int32 `protobuf:"varint,3,opt,name=hunter_src_cell,json=hunterSrcCell,proto3" json:"hunter_src_cell,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CapturePayload_TargetDetails) Reset() {
	*x = CapturePayload_TargetDetails{}
	mi := &file_base_payload_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CapturePayload_TargetDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapturePayload_TargetDetails) ProtoMessage() {}

func (x *CapturePayload_TargetDetails) ProtoReflect() protoreflect.Message {
	mi := &file_base_payload_proto_msgTypes[8]
	if x != nil {
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

func (x *CapturePayload_TargetDetails) GetHunterSrcCell() int32 {
	if x != nil {
		return x.HunterSrcCell
	}
	return 0
}

// Destination of hunterPiece after capturing prey
type CapturePayload_HunterDestination struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CellIndex     int32                  `protobuf:"varint,1,opt,name=cell_index,json=cellIndex,proto3" json:"cell_index,omitempty"`
	X             float32                `protobuf:"fixed32,2,opt,name=x,proto3" json:"x,omitempty"`
	Y             float32                `protobuf:"fixed32,3,opt,name=y,proto3" json:"y,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CapturePayload_HunterDestination) Reset() {
	*x = CapturePayload_HunterDestination{}
	mi := &file_base_payload_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CapturePayload_HunterDestination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapturePayload_HunterDestination) ProtoMessage() {}

func (x *CapturePayload_HunterDestination) ProtoReflect() protoreflect.Message {
	mi := &file_base_payload_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapturePayload_HunterDestination.ProtoReflect.Descriptor instead.
func (*CapturePayload_HunterDestination) Descriptor() ([]byte, []int) {
	return file_base_payload_proto_rawDescGZIP(), []int{5, 1}
}

func (x *CapturePayload_HunterDestination) GetCellIndex() int32 {
	if x != nil {
		return x.CellIndex
	}
	return 0
}

func (x *CapturePayload_HunterDestination) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *CapturePayload_HunterDestination) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

var File_base_payload_proto protoreflect.FileDescriptor

var file_base_payload_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x68, 0x6b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0xa8, 0x03, 0x0a, 0x0b, 0x42, 0x61, 0x73, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
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
	0x74, 0x75, 0x72, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x46, 0x0a, 0x0f, 0x77,
	0x69, 0x6e, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x68, 0x6b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x2e, 0x57, 0x69, 0x6e, 0x4c, 0x6f, 0x73, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x48, 0x00, 0x52, 0x0e, 0x77, 0x69, 0x6e, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x68, 0x0a, 0x0e,
	0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2f,
	0x0a, 0x07, 0x6d, 0x79, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x63, 0x68, 0x6b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x54, 0x65,
	0x61, 0x6d, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x06, 0x6d, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x12,
	0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x50, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x69, 0x65, 0x63, 0x65, 0x73,
	0x5f, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x70, 0x69, 0x65, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x69, 0x65, 0x63, 0x65, 0x73, 0x5f,
	0x62, 0x6c, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x69, 0x65,
	0x63, 0x65, 0x73, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x22, 0x8e, 0x02, 0x0a, 0x0b, 0x4d, 0x6f, 0x76,
	0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x33, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d,
	0x5f, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x68,
	0x6b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x6f,
	0x6c, 0x6f, 0x72, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x19, 0x0a,
	0x08, 0x70, 0x69, 0x65, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x70, 0x69, 0x65, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x63, 0x65, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x65, 0x6c, 0x6c, 0x12, 0x45, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x63, 0x68, 0x6b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x4d, 0x6f, 0x76,
	0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x44, 0x65, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x47, 0x0a, 0x0a, 0x44, 0x65, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x63, 0x65, 0x6c, 0x6c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x0c, 0x0a,
	0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x22, 0x42, 0x0a, 0x0b, 0x45, 0x78, 0x69,
	0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x33, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d,
	0x5f, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x68,
	0x6b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x6f,
	0x6c, 0x6f, 0x72, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x61, 0x6d, 0x22, 0xd4, 0x03,
	0x0a, 0x0e, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x33, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x68, 0x6b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x08, 0x66, 0x72, 0x6f,
	0x6d, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x26, 0x0a, 0x0f, 0x68, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x5f,
	0x70, 0x69, 0x65, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d,
	0x68, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x69, 0x65, 0x63, 0x65, 0x49, 0x64, 0x12, 0x43, 0x0a,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x63, 0x68, 0x6b, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x43, 0x61, 0x70,
	0x74, 0x75, 0x72, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x4f, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x68, 0x6b, 0x2e, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x48, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x7f, 0x0a, 0x0d, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x79, 0x5f, 0x70, 0x69, 0x65,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x72, 0x65,
	0x79, 0x50, 0x69, 0x65, 0x63, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x79,
	0x5f, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x70, 0x72, 0x65, 0x79, 0x43, 0x65, 0x6c, 0x6c, 0x49, 0x64, 0x78, 0x12, 0x26, 0x0a, 0x0f,
	0x68, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x72, 0x63, 0x5f, 0x63, 0x65, 0x6c, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x68, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x72, 0x63,
	0x43, 0x65, 0x6c, 0x6c, 0x1a, 0x4e, 0x0a, 0x11, 0x48, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x65, 0x6c,
	0x6c, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63,
	0x65, 0x6c, 0x6c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x01, 0x79, 0x22, 0x40, 0x0a, 0x0e, 0x57, 0x69, 0x6e, 0x4c, 0x6f, 0x73, 0x65, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x68, 0x6b, 0x2e, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x06,
	0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2a, 0x3f, 0x0a, 0x09, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x6f,
	0x6c, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x45, 0x41,
	0x4d, 0x5f, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x45, 0x41, 0x4d, 0x5f,
	0x42, 0x4c, 0x41, 0x43, 0x4b, 0x10, 0x02, 0x42, 0x17, 0x5a, 0x15, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x65, 0x72, 0x73, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x61, 0x6d, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_base_payload_proto_rawDescOnce sync.Once
	file_base_payload_proto_rawDescData []byte
)

func file_base_payload_proto_rawDescGZIP() []byte {
	file_base_payload_proto_rawDescOnce.Do(func() {
		file_base_payload_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_base_payload_proto_rawDesc), len(file_base_payload_proto_rawDesc)))
	})
	return file_base_payload_proto_rawDescData
}

var file_base_payload_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_base_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_base_payload_proto_goTypes = []any{
	(TeamColor)(0),                           // 0: chk.payload.TeamColor
	(*BasePayload)(nil),                      // 1: chk.payload.BasePayload
	(*WelcomePayload)(nil),                   // 2: chk.payload.WelcomePayload
	(*StartPayload)(nil),                     // 3: chk.payload.StartPayload
	(*MovePayload)(nil),                      // 4: chk.payload.MovePayload
	(*ExitPayload)(nil),                      // 5: chk.payload.ExitPayload
	(*CapturePayload)(nil),                   // 6: chk.payload.CapturePayload
	(*WinLosePayload)(nil),                   // 7: chk.payload.WinLosePayload
	(*MovePayload_Detination)(nil),           // 8: chk.payload.MovePayload.Detination
	(*CapturePayload_TargetDetails)(nil),     // 9: chk.payload.CapturePayload.TargetDetails
	(*CapturePayload_HunterDestination)(nil), // 10: chk.payload.CapturePayload.HunterDestination
}
var file_base_payload_proto_depIdxs = []int32{
	2,  // 0: chk.payload.BasePayload.welcome:type_name -> chk.payload.WelcomePayload
	4,  // 1: chk.payload.BasePayload.move_payload:type_name -> chk.payload.MovePayload
	3,  // 2: chk.payload.BasePayload.start:type_name -> chk.payload.StartPayload
	5,  // 3: chk.payload.BasePayload.exit_payload:type_name -> chk.payload.ExitPayload
	6,  // 4: chk.payload.BasePayload.capture_payload:type_name -> chk.payload.CapturePayload
	7,  // 5: chk.payload.BasePayload.winlose_payload:type_name -> chk.payload.WinLosePayload
	0,  // 6: chk.payload.WelcomePayload.my_team:type_name -> chk.payload.TeamColor
	0,  // 7: chk.payload.MovePayload.from_team:type_name -> chk.payload.TeamColor
	8,  // 8: chk.payload.MovePayload.destination:type_name -> chk.payload.MovePayload.Detination
	0,  // 9: chk.payload.ExitPayload.from_team:type_name -> chk.payload.TeamColor
	0,  // 10: chk.payload.CapturePayload.from_team:type_name -> chk.payload.TeamColor
	9,  // 11: chk.payload.CapturePayload.details:type_name -> chk.payload.CapturePayload.TargetDetails
	10, // 12: chk.payload.CapturePayload.destination:type_name -> chk.payload.CapturePayload.HunterDestination
	0,  // 13: chk.payload.WinLosePayload.winner:type_name -> chk.payload.TeamColor
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_base_payload_proto_init() }
func file_base_payload_proto_init() {
	if File_base_payload_proto != nil {
		return
	}
	file_base_payload_proto_msgTypes[0].OneofWrappers = []any{
		(*BasePayload_Welcome)(nil),
		(*BasePayload_MovePayload)(nil),
		(*BasePayload_Start)(nil),
		(*BasePayload_ExitPayload)(nil),
		(*BasePayload_CapturePayload)(nil),
		(*BasePayload_WinlosePayload)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_base_payload_proto_rawDesc), len(file_base_payload_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_payload_proto_goTypes,
		DependencyIndexes: file_base_payload_proto_depIdxs,
		EnumInfos:         file_base_payload_proto_enumTypes,
		MessageInfos:      file_base_payload_proto_msgTypes,
	}.Build()
	File_base_payload_proto = out.File
	file_base_payload_proto_goTypes = nil
	file_base_payload_proto_depIdxs = nil
}
