package game

import "checkers-backend/player"

type MessageType uint16

const (
	WELCOME MessageType = iota + 49
	START
	EXIT
	MOVE
	CAPTURE
	WIN
	LOSE
)

type WelcomePayload struct {
	MessageType MessageType       `json:"messageType"` // type of message
	MyTeam      player.PlayerType `json:"myTeam"`      // equal to color of Player's pieces
	PiecesRed   []int16           `json:"piecesRed"`   // red's pieces
	PiecesBlack []int16           `json:"piecesBlack"` // black's pieces
}

type Pos struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type MovePayload struct {
	FromTeam       player.PlayerType `json:"fromTeam"`
	CurrentPieceId int16             `json:"currentPieceId"`
	DestPos        Pos               `json:"destPos"`
	SrcCell        int16             `json:"srcCell"`
}
