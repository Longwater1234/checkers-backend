package player

import "golang.org/x/net/websocket"

type PlayerType uint16

const (
	PlayerOne PlayerType = iota + 48834
	PlayerTwo
)

type Player struct {
	Conn   *websocket.Conn // client connection
	Name   string          // Name can only be X or O
	Pieces []int           // cell indexes used by this player
	Dead   chan bool       // whether player has disconnected
}
