package player

import (
	"log"

	"golang.org/x/net/websocket"
	"google.golang.org/protobuf/proto"
)

type PlayerType uint16

const (
	RED PlayerType = iota + 48834
	BLACK
)

type Player struct {
	Conn   *websocket.Conn // client connection
	Name   string          // Name can only be RED or BLACK
	Pieces []int32         // cell indexes used by this player
	Dead   chan bool       // whether player has disconnected
}

// convert PlayerType to string
func (t PlayerType) SimpleName() string {
	switch t {
	case RED:
		return "RED"
	case BLACK:
		return "BLACK"
	}
	return "unknown"
}

// SendMessage - sends the payload as PROTOBUF to this player
func (p *Player) SendMessage(payload proto.Message) {
	bb, err := proto.Marshal(payload)
	if err != nil {
		log.Println("Failed to Marhal message", err)
		p.Dead <- true
	}
	if err := websocket.Message.Send(p.Conn, bb); err != nil {
		log.Println("Failed to sendMessage to", p.Name, ".Reason: ", err)
		p.Dead <- true
	}

}
