package room

import "checkers-backend/game"

// validateAndUpdateMap validates player's move and update gameMap. Returns TRUE if move success, else FALSE
func validateAndUpdateMap(payload *game.MovePayload, gameMap map[int32]*game.Piece) bool {
	srcCell := payload.SourceCell
	destCell := payload.GetDestCell()
	movingPieceId := payload.PieceId

	piecePtr, exists := gameMap[srcCell]
	if !exists || movingPieceId != piecePtr.Id {
		return false
	}

	//check if destCell already has a Piece or not
	_, hasValue := gameMap[destCell.CellIndex]
	if hasValue {
		return false
	}

	success := piecePtr.MoveSimple(&game.Vec2{
		X: destCell.GetX(),
		Y: destCell.GetY(),
	})
	if !success {
		return false
	}
	delete(gameMap, srcCell)                    // set old location empty!
	gameMap[destCell.GetCellIndex()] = piecePtr // fill in the new location
	return true
}
