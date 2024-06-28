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
	//hunterPieceId := captureReq.GetHunterPieceId()
	//preyPiece := captureReq.GetDetails().GetPreyPieceId()

	return true
}
