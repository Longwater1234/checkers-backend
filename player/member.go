package player

import (
	"context"
	"log"
	"slices"
	"sync"
	"time"

	"golang.org/x/net/websocket"
	"google.golang.org/protobuf/proto"
)

type Player struct {
	Conn   *websocket.Conn // client's WS connection
	Name   string          // Name can only be RED or BLACK
	Pieces []int32         // pieces IDs owned by this player. Max size 12
	Quit   <-chan struct{} // to RECEIVE signal this player has quit BEFORE match starts

	dead     chan struct{} // closed when this player is gone after the match ends.
	deadOnce sync.Once
}

// pingCodec is used to PING the client
var pingCodec = websocket.Codec{Marshal: func(v any) (data []byte, payloadType byte, err error) {
	return nil, websocket.PingFrame, nil
}}

// New creates a Player for the given connection.
func New(conn *websocket.Conn) *Player {
	return &Player{
		Conn:   conn,
		Pieces: make([]int32, 12),
		dead:   make(chan struct{}),
	}
}

// Done is closed once this player is gone.
func (p *Player) Done() <-chan struct{} {
	return p.dead
}

// SendMessage as protobuf to this player, and if it fails, eject this player.
func (p *Player) SendMessage(payload proto.Message) {
	bb, err := proto.Marshal(payload)
	if err != nil {
		p.Kill()
		return
	}
	p.Conn.SetWriteDeadline(time.Now().Add(5 * time.Second)) // don't let a sleeping client block us
	defer p.Conn.SetWriteDeadline(time.Time{})               // resets the timer

	if err := websocket.Message.Send(p.Conn, bb); err != nil {
		log.Println("Failed to sendMessage to", p.Name, ". Reason: ", err)
		p.Kill()
	}
}

// LosePiece removes captured `targetPieceId` from player's ownership
func (p *Player) LosePiece(targetPieceId int32) {
	for i := range p.Pieces {
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
// It sets `p.Quit` BEFORE returning (so callers never read a nil channel), then pings in a goroutine.
func (p *Player) StartHeartbeat(ctx context.Context) {
	qq := make(chan struct{}, 1)
	p.Quit = qq
	go func() {
		tt := time.NewTicker(time.Second)
		defer tt.Stop()
		for {
			select {
			case <-tt.C:
				if err := pingCodec.Send(p.Conn, nil); err != nil {
					// This player has quit early
					qq <- struct{}{}
					close(qq)
					return
				}
			case <-ctx.Done():
				// timeout in lobby (OR p2 just joined)
				return
			}
		}
	}()
}

// Kill marks this player as gone. It is safe to call multiple times.
func (p *Player) Kill() {
	p.deadOnce.Do(func() { close(p.dead) })
}
