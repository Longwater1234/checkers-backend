package game

type MessageType uint8

const (
	WELCOME MessageType = iota
	START
	EXIT
	MOVE
	CAPTURE
	WIN
	LOSE
)

type WelcomePayload struct {
	Pieces []int16 `json:"pieces"`
	Name   string  `json:"name"`
}

type Pos struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type MovePayload struct {
	FromUser       string `json:"fromUser"`
	CurrentPieceId int16  `json:"currentPieceId"`
	DestPos        Pos    `json:"destPos"`
	SrcCell        int16  `json:"srcCell"`
}
