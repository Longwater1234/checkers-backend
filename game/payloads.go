package game

import "checkers-backend/player"

type MessageType uint16

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
	MessageType MessageType       `json:"messageType"`
	MyType      player.PlayerType `json:"myType"`
	PiecesRed   []int16           `json:"piecesRed"`
	PiecesBlack []int16           `json:"piecesBlack"`
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
