package room

import (
	"checkers-backend/game"
	"checkers-backend/player"
)

// handleMovePiece validates "MOVE" made by Player `current` against `opponent`.
// Updates Map and returns TRUE if all is OK. Else returns FALSE.
func handleMovePiece(payload *game.BasePayload, gameMap map[int32]*game.Piece, current, opponent *player.Player) bool {
	success := validateAndUpdateMap(payload.GetMovePayload(), gameMap)
	if !success {
		current.SendMessage(&game.BasePayload{
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
	//Else, forward the "MOVE" payload to opponent
	opponent.SendMessage(payload)
	return true
}

// validateAndUpdateMap validates player's move and update gameMap. Returns TRUE if move success, else FALSE
func validateAndUpdateMap(payload *game.MovePayload, gameMap map[int32]*game.Piece) bool {
	srcCell := payload.GetSourceCell()
	destCell := payload.GetDestCell()
	movingPieceId := payload.GetPieceId()

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
