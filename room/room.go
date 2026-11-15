package room

import (
	"checkers-backend/game"
	"checkers-backend/player"
	"log"
	"time"

	"golang.org/x/net/websocket"
	"google.golang.org/protobuf/proto"
)

// StartMatch between the two players `p1` and `p2`. Closes `gameOver` when the match ends.
func StartMatch(p1 *player.Player, p2 *player.Player, gameOver chan<- struct{}) {
	log.Println("🟢 Match has begun!")

	var gameMap map[int32]*game.Piece
	defer func() {
		clear(gameMap)
		p1.Pieces = nil
		p2.Pieces = nil
		close(gameOver)
	}()

	// make random pieceId's for both players
	if errx := generatePieceIds(p1, p2); errx != nil {
		log.Panic("[ERROR] Severe problem: cannot generate random numbers ", errx)
	}

	notifyMatchStart(p1, p2)

	var isPlayerRedTurn = true        // Whose turn is it now? RED always starts.
	gameMap = generateGameMap(p1, p2) // map of cell index --> pieces.

	// MAIN GAME LOOP (each player has MAX 30 sec to respond)
	for {
		hunter, opponent := p1, p2
		if !isPlayerRedTurn {
			hunter, opponent = p2, p1
		}

		// which team triggered the exit
		fromTeam := game.TeamColor_TEAM_RED
		if hunter.Name == game.TeamColor_TEAM_BLACK.String() {
			fromTeam = game.TeamColor_TEAM_BLACK
		}

		// read from connection
		var rawBytes []byte
		hunter.Conn.SetReadDeadline(time.Now().Add(30 * time.Second))
		if err := websocket.Message.Receive(hunter.Conn, &rawBytes); err != nil {
			log.Println(hunter.Name, "disconnected. Cause:", err)
			opponent.SendMessage(&game.BasePayload{
				Notice: "Opponent has left the game!",
				Inner: &game.BasePayload_ExitPayload{
					ExitPayload: &game.ExitPayload{
						FromTeam: fromTeam,
					},
				},
			})
			return
		}

		var request game.BasePayload
		if err := proto.Unmarshal(rawBytes, &request); err != nil {
			log.Println("[ERROR] failed to parse protobuf", err)
			return
		}

		if request.GetMovePayload() != nil {
			// ============== MESSAGE_TYPE :: "move" ==================== //
			if valid := processMovePiece(&request, gameMap, hunter, opponent); !valid {
				return
			}
			if game.HasWinner(hunter, opponent, gameMap) {
				time.Sleep(3 * time.Second)
				return
			}
			isPlayerRedTurn = !isPlayerRedTurn
		} else if request.GetCapturePayload() != nil {
			// ============== MESSAGE_TYPE :: "capture" ==================== //
			capture := request.GetCapturePayload()
			isKingBefore := getKingStatusBefore(capture, gameMap)
			if valid := processCapturePiece(&request, gameMap, hunter, opponent); !valid {
				return
			}
			if game.HasWinner(hunter, opponent, gameMap) {
				time.Sleep(3 * time.Second)
				return
			}
			isKingNow := getKingStatusAfter(capture, gameMap)
			currentCell := capture.GetDestination().GetCellIndex()
			var needCheck bool = isKingBefore == isKingNow
			if needCheck && hasExtraTargets(hunter, currentCell, gameMap) {
				log.Println(hunter.Name, "has extra targets!")
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
