package main

import (
	"checkers-backend/player"
	"crypto/rand"
	"fmt"
	"math/big"
	"sync/atomic"

	"golang.org/x/net/websocket"
)

const SERVER_VERSION = "1.0.0"
const upperLimit int16 = 0x7FFF

var numPlayers atomic.Uint32             // total number of LIVE players
var lobby = make(chan *player.Player, 2) // waiting room for players

// wsHandler assigns name to Player and redirects to Lobby
func wsHandler(ws *websocket.Conn) {
	ws.MaxPayloadBytes = 1024
	defer ws.Close()
}

func main() {
	fmt.Println("hello world!")
	ans, err := rand.Int(rand.Reader, big.NewInt(int64(upperLimit)))
	if err != nil {
		panic(err)
	}
	fmt.Println(ans)
}
