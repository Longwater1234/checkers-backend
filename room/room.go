package room

import (
	"checkers-backend/game"
	"checkers-backend/player"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"golang.org/x/net/websocket"
	"google.golang.org/protobuf/proto"
)

const upperLimit int16 = 0x7FFF //random ID max value (short_max)

// RunMatch between the two players. If game ends, send signal through `gameOver` channel
func RunMatch(p1 *player.Player, p2 *player.Player, gamOver chan bool) {
	log.Println("🟢 Match has begun!")

	//make random pieceId's for player 1
	for i := 0; i < len(p1.Pieces); i++ {
		val, err := rand.Int(rand.Reader, big.NewInt(int64(upperLimit)))
		if err != nil {
			gamOver <- true
			log.Panic("cannot generate random number", err)
		}
		p1.Pieces[i] = int32(val.Int64())
	}

	//make pieces for player 2
	for i := 0; i < len(p2.Pieces); i++ {
		val, err := rand.Int(rand.Reader, big.NewInt(int64(upperLimit)))
		if err != nil {
			gamOver <- true
			log.Panic("cannot generate random number", err)
		}
		p2.Pieces[i] = int32(val.Int64())
	}

	p1.SendMessage(&game.BasePayload{
		Notice: "Opponent joined. Make your first move!",
		Inner: &game.BasePayload_Start{
			Start: &game.StartPayload{
				PiecesRed:   p1.Pieces,
				PiecesBlack: p2.Pieces,
			},
		},
	})

	p2.SendMessage(&game.BasePayload{
		Notice: "Match has begun. Waiting for RED to move!",
		Inner: &game.BasePayload_Start{
			Start: &game.StartPayload{
				PiecesRed:   p1.Pieces,
				PiecesBlack: p2.Pieces,
			},
		},
	})

	var isPlayerRedTurn = true            // Who turn is it now? RED always starts.
	var gameMap = generateGameMap(p1, p2) // map of cell index --> pieces
	for i, v := range gameMap {
		fmt.Println(i, "->", v.Id)
	}

	//START GAME MAIN LOOP
	for {
		if isPlayerRedTurn {
			//IT'S PLAYER 1'S TURN
			var rawBytes []byte
			if err := websocket.Message.Receive(p1.Conn, &rawBytes); err != nil {
				log.Println(p1.Name, "disconnected. Cause:", err)
				p2.SendMessage(&game.BasePayload{
					Notice: "Opponent has left the game!",
					Inner: &game.BasePayload_ExitPayload{
						ExitPayload: &game.ExitPayload{
							FromTeam: game.TeamColor_TEAM_RED,
						},
					},
				})
				gamOver <- true
				return
			}

			var payload game.BasePayload
			if err := proto.Unmarshal(rawBytes, &payload); err != nil {
				log.Println("failed  to read protobuf", err)
				gamOver <- true
				return
			}

			if payload.GetMovePayload() != nil {
				log.Println("movePayload:", payload.GetMovePayload().String())
				//TODO validate move in separate function here
				// if result == false, kickout p1 and notify p2  (separate messages)
				// then end match at once
			}

			//FORWARD THE "MOVE" PAYLOAD TO PLAYER 2,
			if err := websocket.Message.Send(p2.Conn, rawBytes); err != nil {
				log.Println(p2.Name, "disconnected. Cause:", err.Error())
				p1.SendMessage(&game.BasePayload{
					Notice: "Opponent has left the game!",
					Inner: &game.BasePayload_ExitPayload{
						ExitPayload: &game.ExitPayload{
							FromTeam: game.TeamColor_TEAM_BLACK,
						},
					},
				})
				gamOver <- true
				return
			}

			isPlayerRedTurn = false
		} else {
			//IT'S PLAYER 2 (BLACK) TURN
			var rawBytes []byte
			if err := websocket.Message.Receive(p2.Conn, &rawBytes); err != nil {
				log.Println(p1.Name, "disconnected. Cause:", err.Error())
				p1.SendMessage(&game.BasePayload{
					Notice: "Opponent has left the game!",
					Inner: &game.BasePayload_ExitPayload{
						ExitPayload: &game.ExitPayload{
							FromTeam: game.TeamColor_TEAM_BLACK,
						},
					},
				})
				gamOver <- true
				return
			}

			var payload game.BasePayload
			if err := proto.Unmarshal(rawBytes, &payload); err != nil {
				log.Println("failed  to read protobuf", err)
				gamOver <- true
				return
			}

			log.Println(payload.String())

			if payload.GetMovePayload() != nil {
				log.Println("moveTo", payload.GetMovePayload().String())
				//TODO validate move in separate function here
				// if result == false, abort and notify p1 and p2  (separate messages)
				// then end match at once
			}

			//FORWARD THE "MOVE" PAYLOAD TO PLAYER 1, FOR UI UPDATE
			if err := websocket.Message.Send(p1.Conn, rawBytes); err != nil {
				log.Println(p2.Name, "disconnected. Cause:", err.Error())
				p2.SendMessage(&game.BasePayload{
					Notice: "Opponent has left the game!",
					Inner: &game.BasePayload_ExitPayload{
						ExitPayload: &game.ExitPayload{
							FromTeam: game.TeamColor_TEAM_RED,
						},
					},
				})
				gamOver <- true
				return
			}

			isPlayerRedTurn = true
		}
	}

}
