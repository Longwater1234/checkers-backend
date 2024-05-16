package player

import (
	"log"

	"golang.org/x/net/websocket"
)

type PlayerType uint16

const (
	RED PlayerType = iota + 48834
	BLACK
)

type Player struct {
	Conn   *websocket.Conn // client connection
	Name   string          // Name can only be X or O
	Pieces []int16         // cell indexes used by this player
	Dead   chan bool       // whether player has disconnected
}

func (t PlayerType) SimpleName() string {
	switch t {
	case RED:
		return "RED"
	case BLACK:
		return "BLACK"
	}
	return "unknown"
}

// SendMessage as JSON to this player
func (p *Player) SendMessage(payload any) {
	err := websocket.JSON.Send(p.Conn, payload)
	if err != nil {
		log.Printf("Failed to sendMessage to %s. Cause %+v", p.Name, err)
		p.Dead <- true
	}
}
