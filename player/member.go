package player

import (
	"log"
	"slices"

	"golang.org/x/net/websocket"
	"google.golang.org/protobuf/proto"
)

type Player struct {
	Conn   *websocket.Conn // client's WS connection
	Name   string          // Name can only be RED or BLACK
	Pieces []int32         // pieces IDs owned by this player. Max count 12
	Dead   chan bool       // to signal player has disconnected
}

// SendMessage as PROTOBUF to this player
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

// LosePiece removes captured `targetPieceId` from player's basket
func (p *Player) LosePiece(targetPieceId int32) {
	for i := 0; i < len(p.Pieces); i++ {
		if p.Pieces[i] == targetPieceId {
			p.Pieces = slices.Delete(p.Pieces, i, i+1)
			return
		}
	}
}
