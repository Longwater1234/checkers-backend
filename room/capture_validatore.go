package room

import (
	"checkers-backend/game"
	"checkers-backend/player"
)

// handleCapturePiece made by Player `p` against `opponent`. Returns TRUE if all is OK. Else returns FALSE.
func handleCapturePiece(basePayload *game.BasePayload, gameMap map[int32]*game.Piece, p, opponent *player.Player) bool {
	ok := validateCatptureUpdateMap(basePayload.GetCapturePayload(), gameMap, opponent)
	if !ok {
		return false
	}

	return true
}

func validateCatptureUpdateMap(captureReq *game.CapturePayload, gameMap map[int32]*game.Piece, opponent *player.Player) bool {
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

	//check if destCell already has a Piece or not
	_, hasValue := gameMap[captureReq.GetHunterDestCell().CellIndex]
	if hasValue {
		return false
	}

	destCell := captureReq.GetHunterDestCell()
	success := hunterPiecePtr.MoveCapture(&game.Vec2{
		X: destCell.GetX(),
		Y: destCell.GetY(),
	})

	if !success {
		return false
	}

	delete(gameMap, hunterSrc)                        // set hunter's old location empty!
	delete(gameMap, preyCell)                         // set Prey's old location empty!
	gameMap[destCell.GetCellIndex()] = hunterPiecePtr // fill in hunter new location
	opponent.LosePiece(preyPieceId)                   // the defending player loses 1 piece
	return true
}
