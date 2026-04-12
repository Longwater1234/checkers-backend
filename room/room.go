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

	//make random pieceId's for both players
	if errx := generatePieces(p1, p2); errx != nil {
		gameOver <- true
		log.Panic("cannot generate pieces", errx)
	}

	notifyMatchStart(p1, p2)

	var isPlayerRedTurn = true            // Whose turn is it now? RED always starts.
	var gameMap = generateGameMap(p1, p2) // map of cell index --> pieces.

	// cleanup after match ends
	defer func() {
		close(gameOver)
		clear(gameMap)
		p1.Pieces = nil
		p2.Pieces = nil
	}()

	// MAIN GAME LOOP (each player has MAX 30 sec to respond)
	for {
		current, opponent := p1, p2
		if !isPlayerRedTurn {
			current, opponent = p2, p1
		}
		// which team triggered the exit
		fromTeam := game.TeamColor_TEAM_RED
		if current.Name == game.TeamColor_TEAM_BLACK.String() {
			fromTeam = game.TeamColor_TEAM_BLACK
		}
		// read from connection
		var rawBytes []byte
		current.Conn.SetReadDeadline(time.Now().Add(30 * time.Second))
		if err := websocket.Message.Receive(current.Conn, &rawBytes); err != nil {
			log.Println(current.Name, "disconnected. Cause:", err)
			opponent.SendMessage(&game.BasePayload{
				Notice: "Opponent has left the game!",
				Inner: &game.BasePayload_ExitPayload{
					ExitPayload: &game.ExitPayload{
						FromTeam: fromTeam,
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

		if payload.GetMovePayload() != nil {
			// ============== MESSAGE_TYPE :: "move" ==================== //
			if valid := processMovePiece(&payload, gameMap, current, opponent); !valid {
				gameOver <- true
				return
			}
			isPlayerRedTurn = !isPlayerRedTurn
		} else if payload.GetCapturePayload() != nil {
			// ============== MESSAGE_TYPE :: "capture" ==================== //
			capture := payload.GetCapturePayload()
			isKingBefore := getKingStatusBefore(capture, gameMap)
			if valid := processCapturePiece(&payload, gameMap, current, opponent); !valid {
				gameOver <- true
				return
			}
			if game.HasWinner(current, opponent) {
				time.Sleep(3 * time.Second)
				gameOver <- true
				return
			}
			isKingNow := getKingStatusAfter(capture, gameMap)
			currentCell := capture.GetDestination().GetCellIndex()
			var needCheck bool = isKingBefore == isKingNow
			if needCheck && hasExtraTargets(current, currentCell, gameMap) {
				log.Println(current.Name, "has extra targets!")
				continue
			}
			isPlayerRedTurn = !isPlayerRedTurn
			// ... RETURN TO TOP^
		}
	}

}

// notifyMatchStart to both players, and distribute their pieces
func notifyMatchStart(p1 *player.Player, p2 *player.Player) {
	startPayload := &game.StartPayload{
		PiecesRed:   p1.Pieces,
		PiecesBlack: p2.Pieces,
	}

	p1.SendMessage(&game.BasePayload{
		Notice: "Opponent joined. Make your first move!",
		Inner:  &game.BasePayload_Start{Start: startPayload},
	})

	p2.SendMessage(&game.BasePayload{
		Notice: "Match has begun. Waiting for RED to move!",
		Inner:  &game.BasePayload_Start{Start: startPayload},
	})
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
