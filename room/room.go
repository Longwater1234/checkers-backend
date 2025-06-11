package room

import (
	"checkers-backend/game"
	"checkers-backend/player"
	"log"
	"time"

	"golang.org/x/net/websocket"
	"google.golang.org/protobuf/proto"
)

// StartMatch between the two players `p1` and `p2`. If match ends, send signal through `gameOver` channel
func StartMatch(p1 *player.Player, p2 *player.Player, gameOver chan<- bool) {
	log.Println("🟢 Match has begun!")

	//make random pieceId's for both
	generatePieces(p1, p2, gameOver)

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
	var gameMap = generateGameMap(p1, p2) // map of cell index --> pieces.

	// free up memory after match ends
	defer func() {
		gameOver <- true
		clear(gameMap)
		p1.Pieces = nil
		p2.Pieces = nil
	}()

	//START GAME MAIN LOOP
	for {
		if isPlayerRedTurn {
			// ============= IT'S PLAYER 1 (RED's) TURN =============//
			var rawBytes []byte
			p1.Conn.SetReadDeadline(time.Now().Add(time.Second * 40))
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

			//if MESSAGE TYPE == "move"
			if payload.GetMovePayload() != nil {
				//log.Println("move", payload.GetMovePayload().String())
				if valid := processMovePiece(&payload, gameMap, p1, p2); !valid {
					gameOver <- true
					return
				}
				isPlayerRedTurn = false
			} else if payload.GetCapturePayload() != nil {
				//if MESSAGE TYPE == "capture"
				isKingBefore := getKingStatusBefore(payload.GetCapturePayload(), gameMap)
				if valid := processCapturePiece(&payload, gameMap, p1, p2); !valid {
					gameOver <- true
					return
				}
				if game.HasWinner(p1, p2) {
					time.Sleep(3 * time.Second)
					gameOver <- true
					return
				}
				isKingNow := getKingStatusAfter(payload.GetCapturePayload(), gameMap)
				currentCell := payload.GetCapturePayload().Destination.CellIndex
				var needCheck bool = isKingBefore == isKingNow
				// CHECK for extra opportunities for P1. if NONE, toggle turns
				if needCheck && hasExtraTargets(p1, currentCell, gameMap) {
					log.Println(p1.Name, " has extra targets!")
					continue
				}
				isPlayerRedTurn = false
			}
		} else if !isPlayerRedTurn {
			// ============= IT'S PLAYER 2 (BLACK's) TURN =============//
			var rawBytes []byte
			p2.Conn.SetReadDeadline(time.Now().Add(time.Second * 40))
			if err := websocket.Message.Receive(p2.Conn, &rawBytes); err != nil {
				log.Println(p2.Name, "disconnected. Cause:", err.Error())
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

			//if MESSAGE TYPE == "move"
			if payload.GetMovePayload() != nil {
				//log.Println("move", payload.GetMovePayload().String())
				valid := processMovePiece(&payload, gameMap, p2, p1)
				if !valid {
					gameOver <- true
					return
				}
				isPlayerRedTurn = true
			} else if payload.GetCapturePayload() != nil {
				//if MESSAGE TYPE == "capture"
				//log.Println("capture", payload.GetCapturePayload().String())
				isKingBefore := getKingStatusBefore(payload.GetCapturePayload(), gameMap)
				valid := processCapturePiece(&payload, gameMap, p2, p1)
				if !valid {
					gameOver <- true
					return
				}
				if game.HasWinner(p2, p1) {
					time.Sleep(3 * time.Second)
					gameOver <- true
					return
				}
				//check for extra opportunities for P2. if NONE, toggle turns
				isKingNow := getKingStatusAfter(payload.GetCapturePayload(), gameMap)
				currentCell := payload.GetCapturePayload().Destination.CellIndex
				var needCheck bool = isKingBefore == isKingNow
				if needCheck && hasExtraTargets(p2, currentCell, gameMap) {
					log.Println(p2.Name, " has extra targets!")
					continue
				}
				isPlayerRedTurn = true
			}
			// ... return to top
		}
	}
}

// getKingStatusBefore capturing the opponent. Returns TRUE if piece is King at its original Cell
func getKingStatusBefore(capturePayload *game.CapturePayload, gameMap map[int32]*game.Piece) bool {
	if capturePayload == nil || capturePayload.GetDetails() == nil {
		return false
	}
	srcCell := capturePayload.GetDetails().GetHunterSrcCell()
	if piecePtr, exists := gameMap[srcCell]; !exists {
		return false
	} else {
		return piecePtr.IsKing
	}
}

// getKingStatusAfter capturing opponent.  Returns TRUE if piece is King at destination Cell
func getKingStatusAfter(capturePayload *game.CapturePayload, gameMap map[int32]*game.Piece) bool {
	if capturePayload == nil || capturePayload.GetDestination() == nil {
		return false
	}
	destCell := capturePayload.GetDestination().GetCellIndex()
	if piecePtr, exists := gameMap[destCell]; !exists {
		return false
	} else {
		return piecePtr.IsKing
	}
}
