package room

import (
	"checkers-backend/game"
	"checkers-backend/player"
)

// processCapturePiece made by Player `p` against `opponent`. Returns TRUE if all is OK. Else returns FALSE.
func processCapturePiece(basePayload *game.BasePayload, gameMap map[int32]*game.Piece, p, opponent *player.Player) bool {
	capturePayload := basePayload.GetCapturePayload()
	success := validateCapture(capturePayload, gameMap)
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
	//all is OK, Opponent loses 1 piece
	preyPieceId := capturePayload.GetDetails().GetPreyPieceId()
	opponent.LosePiece(preyPieceId)
	opponent.SendMessage(basePayload)
	return true
}

// validateCapture when player `p` attacks by opponent's piece, AND then updates gameMap. returns TRUE if success
func validateCapture(captureReq *game.CapturePayload, gameMap map[int32]*game.Piece) bool {
	if captureReq.GetDetails() == nil || captureReq.GetDestination() == nil {
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
	preyCellIdx := captureReq.GetDetails().PreyCellIdx

	//check Prey params
	preyPiecePtr, exists := gameMap[preyCellIdx]
	if !exists || preyPieceId != preyPiecePtr.Id {
		return false
	}

	// check whether destCell already has a Piece
	destination := captureReq.GetDestination()
	_, hasValue := gameMap[destination.GetCellIndex()]
	if hasValue {
		return false
	}

	success := hunterPiecePtr.MoveCapture(&game.Vec2{
		X: destination.GetX(),
		Y: destination.GetY(),
	})

	if !success {
		return false
	}
	delete(gameMap, hunterSrc)                           // set hunter's old location empty!
	delete(gameMap, preyCellIdx)                         // set Prey's old location empty!
	gameMap[destination.GetCellIndex()] = hunterPiecePtr // move hunter to new location
	return true
}
