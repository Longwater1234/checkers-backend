package game

type MessageType uint16

const (
	WELCOME MessageType = iota + 9978
	START
	EXIT
	MOVE
	CAPTURE
	WIN
	LOSE
	DRAW
)

type WelcomePayload struct {
	Pieces []int16 `json:"pieces"`
	Name   string  `json:"name"`
}

type JumpPayload struct {
	FromUser       string `json:"fromUser"`
	CurrentPieceId int16  `json:"currentPieceId"`
	DestCell       uint8  `json:"destCell"`
	SrcCell        uint8  `json:"srcCell"`
}
