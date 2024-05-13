package main

import (
	"checkers-backend/game"
	"checkers-backend/player"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"

	"golang.org/x/net/websocket"
)

var SERVER_VERSION = "1.0.0"

const upperLimit int16 = 0x7FFF  //random gen max value
const customMaxPayload = 2 << 10 //2KB

var numPlayers atomic.Uint32             // total number of LIVE players
var lobby = make(chan *player.Player, 2) // waiting room for players

// wsHandler assigns name to Player and redirects to Lobby
func wsHandler(ws *websocket.Conn) {
	ws.MaxPayloadBytes = customMaxPayload
	defer ws.Close()

	var clientIp = ws.Request().RemoteAddr
	defer ws.Close()

	p := &player.Player{
		Conn:   ws,
		Pieces: make([]int16, 12),
		Dead:   make(chan bool, 1),
	}
	defer close(p.Dead)

	//for each pair joining, the 1st will be player 1
	if numPlayers.Load()%2 == 0 {
		p.Name = player.RED.String()
	} else {
		p.Name = player.BLACK.String()
	}

	log.Println("Someone connected", clientIp, "Total players:", numPlayers.Load())
	<-p.Dead
	numPlayers.Add(^uint32(0)) // if player exits, minus 1
	log.Println(clientIp, p.Name, "just left the game. Total players:", numPlayers.Load())

}

func main() {
	portNum, err := strconv.Atoi(os.Getenv("GAME_PORT"))
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

// listen for new players joining lobby
func listenForJoins() {
	//TODO
	for {
		log.Println("LOBBY:", "cap", cap(lobby), "len", len(lobby))
		p1 := <-lobby
		//generate 16 random numbers (upto SHORT_MAX) as pieceId for this player
		for i := 0; i < len(p1.Pieces); i++ {
			val, err := rand.Int(rand.Reader, big.NewInt(int64(upperLimit)))
			if err != nil {
				p1.Dead <- true
				log.Println("could not generate random number", err)
				break
			}
			p1.Pieces[i] = int16(val.Int64())
		}
		p1.SendMessage(&game.WelcomePayload{
			MessageType: game.WELCOME,
			MyTeam:      player.RED,
			PiecesRed:   p1.Pieces,
		})
	}
}
