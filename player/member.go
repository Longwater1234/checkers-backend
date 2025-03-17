package player

import (
	"context"
	"log"
	"slices"
	"time"

	"golang.org/x/net/websocket"
	"google.golang.org/protobuf/proto"
)

type Player struct {
	Conn   *websocket.Conn // client's WS connection
	Name   string          // Name can only be RED or BLACK
	Pieces []int32         // pieces IDs owned by this player. Max size 12
	Dead   chan<- bool     // to SEND signal this player left AFTER match starts
	Quit   <-chan bool     // to RECEIVE signal this player has quit BEFORE match starts
}

// pingCodec is used to PING the client
var pingCodec = websocket.Codec{Marshal: func(v any) (data []byte, payloadType byte, err error) {
	return nil, websocket.PingFrame, nil
}}

// SendMessage as PROTOBUF to this player
func (p *Player) SendMessage(payload proto.Message) {
	bb, err := proto.Marshal(payload)
	if err != nil {
		log.Println("Failed to encode message", err)
		p.Dead <- true
	}
	if err := websocket.Message.Send(p.Conn, bb); err != nil {
		log.Println("Failed to sendMessage to", p.Name, ".Reason: ", err)
		p.Dead <- true
	}
}

// LosePiece removes captured `targetPieceId` from player's ownership
func (p *Player) LosePiece(targetPieceId int32) {
	for i := 0; i < len(p.Pieces); i++ {
		if p.Pieces[i] == targetPieceId {
			p.Pieces = slices.Delete(p.Pieces, i, i+1)
			return
		}
	}
}

// HasThisPiece returns TRUE if this player owns the given `pieceId`
func (p *Player) HasThisPiece(pieceId int32) bool {
	return slices.Contains(p.Pieces, pieceId)
}

// StartHeartbeat for checking (every second) if this Player is still connected (when waiting for opponent)
func (p *Player) StartHeartbeat(ctx context.Context) {
	tt := time.NewTicker(time.Second)
	defer tt.Stop()
	qq := make(chan bool, 1)
	p.Quit = qq
	for {
		select {
		case <-tt.C:
			if err := pingCodec.Send(p.Conn, nil); err != nil {
				// This player has quit early
				qq <- true
				close(qq)
				return
			}
		case <-ctx.Done():
			// timeout waiting for [p2] expired
			tt.Stop()
			return
		}

	}
}
