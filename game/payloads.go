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

type BasePayload struct {
	MessageType MessageType `json:"messageType"` // type of message

}
type WelcomePayload struct {
	BasePayload
	Notice string            `json:"notice"` //text for UI
	MyTeam player.PlayerType `json:"myTeam"` // equal to color of Player's pieces
}

type StartPayload struct {
	BasePayload
	Notice      string  `json:"notice"`      // text for the UI
	PiecesRed   []int16 `json:"piecesRed"`   // red's pieces
	PiecesBlack []int16 `json:"piecesBlack"` // black's pieces
}

type Pos struct {
	X float32 `json:"x"` // cell x position
	Y float32 `json:"y"` // cell y position
}

type MovePayload struct {
	BasePayload
	FromTeam player.PlayerType `json:"fromTeam"` // which Player made the move
	PieceId  int16             `json:"PieceId"`  // the moving pieceId
	DestCell Pos               `json:"destCell"` // target cell position
	SrcCell  int16             `json:"srcCell"`  // souce cell index
}
