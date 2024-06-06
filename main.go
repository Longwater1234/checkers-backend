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

	_ "github.com/goccy/go-json"

	"golang.org/x/net/websocket"
)

var SERVER_VERSION = "1.0.0"

const customMaxPayload int = 2 << 10 //2KB

var numPlayers atomic.Uint32             // total number of LIVE players
var lobby = make(chan *player.Player, 1) // waiting room for players

// wsHandler assigns name to Player and redirects to Lobby
func wsHandler(ws *websocket.Conn) {
	ws.MaxPayloadBytes = customMaxPayload
	defer ws.Close()

	var clientIp = ws.Request().RemoteAddr

	p := &player.Player{
		Conn:   ws,
		Pieces: make([]int16, 12),
		Dead:   make(chan bool, 1),
	}
	defer close(p.Dead)

	//for each pair joining, the 1st will be player 1 (RED)
	if numPlayers.Load()%2 == 0 {
		p.Name = player.RED.SimpleName()
	} else {
		p.Name = player.BLACK.SimpleName()
	}
	numPlayers.Add(1)
	lobby <- p //send player to lobby

	log.Println("Someone connected", clientIp, "Total players:", numPlayers.Load())
	<-p.Dead                   //block until player exits
	numPlayers.Add(^uint32(0)) // if player exits, minus 1
	log.Println(clientIp, p.Name, "just left the game. Total players:", numPlayers.Load())
}

func main() {
	portNum, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		portNum = 9876
	}
	port := strconv.Itoa(portNum)

	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, `<p>This is a socket game server. Dial ws://%s:%s/game </p>`, r.URL.Host, port)
	})
	http.Handle("/game", websocket.Handler(wsHandler))

	go listenForJoins()
	log.Println("Server listening at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// Keep Listening for new players joining lobby
func listenForJoins() {
	for {
		log.Println("LOBBY:", "cap", cap(lobby), "len", len(lobby))
		p1 := <-lobby

		p1.SendMessage(&game.WelcomePayload{
			BasePayload: game.BasePayload{
				MessageType: game.WELCOME,
			},
			MyTeam: player.RED,
			Notice: "Connected. You are Team RED. Waiting for opponent...",
		})

		p2 := <-lobby //waiting for 2nd player to join
		p2.SendMessage(&game.WelcomePayload{
			BasePayload: game.BasePayload{
				MessageType: game.WELCOME,
			},
			Notice: "Connected. You are Team BLACK. Match is starting!",
			MyTeam: player.BLACK,
		})

		//start the match in new goroutine
		go func(p1 *player.Player, p2 *player.Player) {
			gameOver := make(chan bool, 1)
			room.StartMatch(p1, p2, gameOver)
			//block until match ends
			<-gameOver
			log.Println("🔴 GAME OVER!")
			p1.Dead <- true
			p2.Dead <- true
		}(p1, p2)
	}

}
