package room

import (
	"checkers-backend/game"
	"checkers-backend/player"
)

// handleCapturePiece made by Player `p` against `opponent`. Returns TRUE if all is OK. Else returns FALSE.
func handleCapturePiece(basePayload *game.BasePayload, gameMap map[int32]*game.Piece, p, opponent *player.Player) bool {
	success := validateCapture(basePayload.GetCapturePayload(), gameMap, opponent)
	if !success {
		p.SendMessage(&game.BasePayload{
			Notice: "Illegal move!",
			Inner: &game.BasePayload_ExitPayload{
				ExitPayload: &game.ExitPayload{
					FromTeam: game.TeamColor_TEAM_UNSPECIFIED,
				},
			},
		})
		opponent.SendMessage(&game.BasePayload{
			Notice: "Your opponent got kicked out!",
			Inner: &game.BasePayload_ExitPayload{
				ExitPayload: &game.ExitPayload{
					FromTeam: game.TeamColor_TEAM_UNSPECIFIED,
				},
			},
		})
		return false
	}
	// Else, forward the "CAPTURE" payload to opponent
	opponent.SendMessage(basePayload)
	return true
}

// validateCapture and updates map if VALID when opponent's piece is attacked by current player
func validateCapture(captureReq *game.CapturePayload, gameMap map[int32]*game.Piece, opponent *player.Player) bool {
	if captureReq.GetDetails() == nil || captureReq.GetHunterDestCell() == nil {
		return false
	}
	hunterPieceId := captureReq.GetHunterPieceId()
	hunterSrc := captureReq.GetDetails().GetHunterSrcCell()

	//check hunter params
	hunterPiecePtr, exists := gameMap[hunterSrc]
	if !exists || hunterPieceId != hunterPiecePtr.Id {
		return false
	}

	preyPieceId := captureReq.GetDetails().PreyPieceId
	preyCell := captureReq.GetDetails().PreyCellIdx

	//check Prey params
	preyPiecePtr, exists := gameMap[preyCell]
	if !exists || preyPieceId != preyPiecePtr.Id {
		return false
	}

	//check if destCell already has a Piece or not
	destCell := captureReq.GetHunterDestCell()
	_, hasValue := gameMap[destCell.GetCellIndex()]
	if hasValue {
		return false
	}

	success := hunterPiecePtr.MoveCapture(&game.Vec2{
		X: destCell.GetX(),
		Y: destCell.GetY(),
	})

	if !success {
		return false
	}

	delete(gameMap, hunterSrc)                        // set hunter's old location empty!
	delete(gameMap, preyCell)                         // set Prey's old location empty!
	gameMap[destCell.GetCellIndex()] = hunterPiecePtr // move hunter to new location
	opponent.LosePiece(preyPieceId)                   // the opponent loses 1 piece
	return true
}
