package main

import (
	"fmt"
	"sync/atomic"

	"golang.org/x/net/websocket"
)

var numPlayers atomic.Uint32 // total number of LIVE players

// wsHandler assigns name to Player and redirects to Lobby
func wsHandler(ws *websocket.Conn) {
	ws.MaxPayloadBytes = 1024
	defer ws.Close()
}
func main() {
	fmt.Println("hello world!")
}
