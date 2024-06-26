package room

import (
	"checkers-backend/game"
	"checkers-backend/player"
	"log"

	"golang.org/x/net/websocket"
	"google.golang.org/protobuf/proto"
)

// RunMatch between the two players `p1` and `p2`. If match ends, send signal through `gameOver` channel
func RunMatch(p1 *player.Player, p2 *player.Player, gameOver chan bool) {
	log.Println("🟢 Match has begun!")

	//make random pieceId's for both
	generatePlayerPieces(p1, p2, gameOver)

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
				gameOver <- true
				return
			}

			var payload game.BasePayload
			if err := proto.Unmarshal(rawBytes, &payload); err != nil {
				log.Println("failed to parse protobuf", err)
				gameOver <- true
				return
			}

			log.Println(payload.String()) //FIXME delete me in production

			//if MESSAGE TYPE == "move"
			if payload.GetMovePayload() != nil {
				result := handleMovePiece(&payload, gameMap, p1, p2)
				if !result {
					gameOver <- true
					return
				}
			}
			//TODO ELSE IF MESSAGE_TYPE == "CAPTURE", VALIDATE CAPTURE HERE
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
				gameOver <- true
				return
			}

			var payload game.BasePayload
			if err := proto.Unmarshal(rawBytes, &payload); err != nil {
				log.Println("failed to parse protobuf", err)
				gameOver <- true
				return
			}

			log.Println(payload.String()) //FIXME delete me in production

			//CHECK MESSAGE TYPE EQUALs "move"
			if payload.GetMovePayload() != nil {
				result := handleMovePiece(&payload, gameMap, p2, p1)
				if !result {
					gameOver <- true
					return
				}
			}
			//TODO ELSE IF MESSAGE_TYPE == "CAPTURE", VALIDATE CAPTURE HERE
			isPlayerRedTurn = true
		}
	}
}
