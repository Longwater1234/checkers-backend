package main

import (
	"checkers-backend/game"
	"checkers-backend/player"
	"checkers-backend/room"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	_ "net/http/pprof"

	"golang.org/x/net/websocket"
)

const serverVersion = "1.0.12"
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

	go listenForJoins()
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

	//for each pair joining, the First will always be RED
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

// Keep Listening for new players joining lobby. Then forward a pair to new match room
func listenForJoins() {
	for {
		//welcome 1st player
		p1 := <-lobby
		var msgOne = &game.BasePayload{
			Notice: "Connected. You are Team RED. Waiting for opponent...",
			Inner: &game.BasePayload_Welcome{
				Welcome: &game.WelcomePayload{
					MyTeam:        game.TeamColor_TEAM_RED,
					ServerVersion: serverVersion,
				},
			},
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		go p1.StartHeartbeat(ctx)
		p1.SendMessage(msgOne)

		//waiting for 2nd player to join (TIMEOUT at 30 seconds)
		select {
		case p2 := <-lobby:
			cancel()
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

			// start the match in new goroutine
			go func(p1, p2 *player.Player) {
				//Sleep REQUIRED for [p2] client to process prev message
				time.Sleep(200 * time.Millisecond)
				gameOver := make(chan bool, 1)
				room.StartMatch(p1, p2, gameOver)
				<-gameOver // block until match ends
				log.Println("🔴 GAME OVER!")
				p1.Dead <- true
				p2.Dead <- true
			}(p1, p2)

		case <-ctx.Done():
			// timeout reached. No other player joined! Goodbye p1!
			p1.SendMessage(&game.BasePayload{
				Notice: "No other players at this moment. Try again later!",
				Inner: &game.BasePayload_ExitPayload{
					ExitPayload: &game.ExitPayload{
						FromTeam: game.TeamColor_TEAM_UNSPECIFIED,
					},
				},
			})
			p1.Dead <- true

		case <-p1.Quit:
			//[p1] has quit before match started
			cancel()
			p1.Dead <- true
		}
		// goto TOP... wait for another pair to join.
	}

}
