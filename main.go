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

	_ "github.com/goccy/go-json"

	"golang.org/x/net/websocket"
)

var SERVER_VERSION = "1.0.0"

const upperLimit int16 = 0x7FFF      //random ID max value (short_max)
const customMaxPayload int = 2 << 10 //2KB

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
		p.Name = player.RED.SimpleName()
	} else {
		p.Name = player.BLACK.SimpleName()
	}
	numPlayers.Add(1)
	lobby <- p

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

// Listen for new players joining lobby
func listenForJoins() {
	for {
		log.Println("LOBBY:", "cap", cap(lobby), "len", len(lobby))
		p1 := <-lobby
		//generate 16 random numbers (upto SHORT_MAX) as pieceId for this player
		for i := 0; i < len(p1.Pieces); i++ {
			val, err := rand.Int(rand.Reader, big.NewInt(int64(upperLimit)))
			if err != nil {
				p1.Dead <- true
				log.Panic("could not generate random number", err)
				break
			}
			p1.Pieces[i] = int16(val.Int64())
		}

		playerTwoPieces := make([]int16, 12)
		for i := 0; i < len(playerTwoPieces); i++ {
			val, err := rand.Int(rand.Reader, big.NewInt(int64(upperLimit)))
			if err != nil {
				log.Panic("could not generate random number", err)
			}
			playerTwoPieces[i] = int16(val.Int64())
		}

		p1.SendMessage(&game.WelcomePayload{
			MessageType: game.WELCOME,
			MyTeam:      player.RED,
			Notice:      "Connected. Waiting for opponent...",
			PiecesRed:   p1.Pieces,
			PiecesBlack: playerTwoPieces,
		})

		p2 := <-lobby                    //waiting for 2nd player to join
		copy(p2.Pieces, playerTwoPieces) // copy pieces to p2
		p2.SendMessage(&game.WelcomePayload{
			MessageType: game.WELCOME,
			Notice:      "Connected. Game is starting!",
			MyTeam:      player.BLACK,
			PiecesRed:   p1.Pieces,
			PiecesBlack: p2.Pieces,
		})

		playerTwoPieces = nil // we dont need this anymore

		//TODO start a timer of 30 seconds to wait for 2nd player. if timeout, close p1 connection
		var payload game.WelcomePayload
		if err := websocket.JSON.Receive(p1.Conn, &payload); err != nil {
			log.Println(p1.Name, "disconnected. Cause:", err.Error())
			p1.Dead <- true
			//return
		}

		//var payload game.WelcomePayload
		if err := websocket.JSON.Receive(p2.Conn, &payload); err != nil {
			log.Println(p2.Name, "disconnected. Cause:", err.Error())
			p2.Dead <- true
			//return
		}
		//start the match in new goroutine
		// go func(p1 *player.Player, p2 *player.Player) {
		// 	gameOver := make(chan bool, 1)
		// 	room.StartMatch(p1, p2, gameOver)
		// 	//block until match ends
		// 	<-gameOver
		// 	log.Println("🔴 GAME OVER!")
		// 	p1.Dead <- true
		// 	p2.Dead <- true
		// }(p1, p2)
	}

}
