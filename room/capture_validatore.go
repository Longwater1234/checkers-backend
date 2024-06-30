package room

import (
	"checkers-backend/game"
	"checkers-backend/player"
)

// handleCapturePiece made by Player `p` against `opponent`. Returns TRUE if all is OK. Else returns FALSE.
func handleCapturePiece(basePayload *game.BasePayload, gameMap map[int32]*game.Piece, p, opponent *player.Player) bool {
	validateCatptureUpdateMap(basePayload.GetCapturePayload(), gameMap)
	return true
}

func validateCatptureUpdateMap(captureReq *game.CapturePayload, gameMap map[int32]*game.Piece) bool {
	if captureReq.GetDetails() == nil {
		return false
	}
	hunterPieceId := captureReq.GetHunterPieceId()
	hunterSrc := captureReq.GetDetails().HunterSrcCell

	//check hunter params
	hunterPiecePtr, exists := gameMap[hunterSrc]
	if !exists || hunterPieceId != hunterPiecePtr.Id {
		return false
	}

	preyPieceId := captureReq.GetDetails().GetPreyPieceId()
	preyCell := captureReq.GetDetails().GetPreyCellIdx()

	//check Prey params
	preyPiecePtr, exists := gameMap[preyCell]
	if !exists || preyPieceId != preyPiecePtr.Id {
		return false
	}

	return true
}
