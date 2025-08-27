package main

import (
	"checkers-backend/game"
	"checkers-backend/player"
	"checkers-backend/room"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"

	_ "net/http/pprof"

	"golang.org/x/net/websocket"
)

const maxRequestSize int = 1 << 10 // 1KB

var numPlayers atomic.Uint32             // total number of LIVE players
var lobby = make(chan *player.Player, 1) // waiting room for players

func main() {
	portNum, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		portNum = 9876
	}
	port := strconv.Itoa(portNum)

	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(writer, `<p>This is a websocket server. Dial ws://{requestURI}/game </p>`)
	})

	http.Handle("/game", websocket.Handler(wsHandler))

	go room.ListenForJoins(lobby)
	log.Println("Server listening at http://127.0.0.1:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// wsHandler handles every new WS connection and redirects Player to Lobby
func wsHandler(ws *websocket.Conn) {
	ws.MaxPayloadBytes = maxRequestSize
	defer ws.Close()

	var clientIp = ws.Request().RemoteAddr
	deadChan := make(chan bool, 1)
	p := &player.Player{
		Conn:   ws,
		Pieces: make([]int32, 12),
		Dead:   deadChan,
	}

	// for each pair joining, the First will always be RED
	if numPlayers.Load()%2 == 0 {
		p.Name = game.TeamColor_TEAM_RED.String()
	} else {
		p.Name = game.TeamColor_TEAM_BLACK.String()
	}
	numPlayers.Add(1)
	lobby <- p

	log.Println("Someone connected", clientIp, "Total players:", numPlayers.Load())
	<-deadChan                 // block until player exits
	numPlayers.Add(^uint32(0)) // if player exits, minus 1
	log.Println(p.Name, "just left the game. Total players:", numPlayers.Load())
}
