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
	"time"

	"golang.org/x/net/websocket"
)

const SERVER_VERSION = "2024.7.0"

const maxRequestSize int = 2 << 10 //2KB

var numPlayers atomic.Uint32             // total number of LIVE players
var lobby = make(chan *player.Player, 2) // waiting room for players

// wsHandler handles every new WS connection and redirects Player to Lobby
func wsHandler(ws *websocket.Conn) {
	ws.MaxPayloadBytes = maxRequestSize
	defer ws.Close()

	var clientIp = ws.Request().RemoteAddr

	p := &player.Player{
		Conn:   ws,
		Pieces: make([]int32, 12),
		Dead:   make(chan bool, 1),
	}
	defer close(p.Dead)

	//for each pair joining, the First will always be RED
	if numPlayers.Load()%2 == 0 {
		p.Name = game.TeamColor_TEAM_RED.String()
	} else {
		p.Name = game.TeamColor_TEAM_BLACK.String()
	}
	numPlayers.Add(1)
	lobby <- p

	log.Println("Someone connected", clientIp, "Total players:", numPlayers.Load())
	<-p.Dead                   // block until player exits
	numPlayers.Add(^uint32(0)) // if player exits, minus 1
	log.Println(p.Name, "just left the game. Total players:", numPlayers.Load())
}

func main() {
	portNum, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		portNum = 9876
	}
	port := strconv.Itoa(portNum)

	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, `<p>This is a socket server. Dial ws://%s:%s/game </p>`, r.URL.Host, port)
	})
	http.Handle("/game", websocket.Handler(wsHandler))

	go listenForJoins()
	log.Println("Server listening at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// Keep Listening for new players joining lobby
func listenForJoins() {
	for {
		//welcome 1st player
		p1 := <-lobby
		log.Println("LOBBY:", "cap", cap(lobby), "len", len(lobby))
		var msgOne = &game.BasePayload{
			Notice: "Connected. You are Team RED. Waiting for opponent...",
			Inner: &game.BasePayload_Welcome{
				Welcome: &game.WelcomePayload{
					MyTeam:        game.TeamColor_TEAM_RED,
					ServerVersion: SERVER_VERSION,
				},
			},
		}
		p1.SendMessage(msgOne)

		//waiting for 2nd player to join (TIMEOUT at 30 seconds)
		t := time.NewTimer(30 * time.Second)
		select {
		case p2 := <-lobby:
			t.Stop()
			// welcome 2nd player
			var msgTwo = &game.BasePayload{
				Notice: "Connected. You are Team BLACK. Match is starting!",
				Inner: &game.BasePayload_Welcome{
					Welcome: &game.WelcomePayload{
						MyTeam: game.TeamColor_TEAM_BLACK,
					},
				},
			}
			p2.SendMessage(msgTwo)

			//start the match in new goroutine
			go func(p1, p2 *player.Player) {
				//Sleep necessary for [p2] Client to process prev message
				time.Sleep(100 * time.Millisecond)
				gameOver := make(chan bool, 1)
				room.RunMatch(p1, p2, gameOver)
				<-gameOver //block until match ends
				log.Println("🔴 GAME OVER!")
				close(gameOver)
				p1.Dead <- true
				p2.Dead <- true
			}(p1, p2)

		case <-t.C:
			// timeout reached. No other player joined! Goodbye p1!
			t.Stop()
			p1.SendMessage(&game.BasePayload{
				Notice: "No other players at this moment. Try again later!",
				Inner: &game.BasePayload_ExitPayload{
					ExitPayload: &game.ExitPayload{
						FromTeam: game.TeamColor_TEAM_UNSPECIFIED,
					},
				},
			})
			p1.Dead <- true
		}
		// goto TOP... wait for another pair to join.
	}

}
