package room

import "checkers-backend/game"

// isValidMove checks if the piece move is legal or not
func isValidMove(payload *game.MovePayload, gameMap map[int32]*game.Piece) bool {
	destCell := payload.DestCell
	srcCell := payload.SourceCell
	movingPieceId := payload.PieceId

	piecePtr, exists := gameMap[srcCell]
	if !exists || movingPieceId != piecePtr.Id {
		return false
	}

	//check if destCell.cellIndex has piece or not
	_, hasValue := gameMap[destCell.CellIndex]
	if hasValue {
		return false
	}

	return piecePtr.MoveSimple(&game.Vec2{
		X: destCell.GetX(),
		Y: destCell.GetY(),
	})
}
