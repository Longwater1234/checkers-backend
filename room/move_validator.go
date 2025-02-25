package room

import (
	"checkers-backend/game"
	"checkers-backend/player"
)

// processMovePiece made by Player `p` against `opponent`. Returns TRUE if all is OK. Else returns FALSE.
func processMovePiece(payload *game.BasePayload, gameMap map[int32]*game.Piece, p, opponent *player.Player) bool {
	success := validateAndUpdateMap(payload.GetMovePayload(), gameMap)
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
	// All is OK, forward the "MOVE" payload to opponent
	opponent.SendMessage(payload)
	return true
}

// validateAndUpdateMap after player's "MOVE" and update gameMap. Returns TRUE if successful, else FALSE
func validateAndUpdateMap(payload *game.MovePayload, gameMap map[int32]*game.Piece) bool {
	destination := payload.GetDestination()
	if destination == nil {
		return false
	}
	srcCellIdx := payload.GetSourceCell()
	movingPieceId := payload.GetPieceId()

	piecePtr, exists := gameMap[srcCellIdx]
	if !exists || movingPieceId != piecePtr.Id {
		return false
	}

	// check whether destCell already has a Piece
	_, hasValue := gameMap[destination.CellIndex]
	if hasValue {
		return false
	}

	success := piecePtr.MoveSimple(&game.Vec2{
		X: destination.GetX(),
		Y: destination.GetY(),
	})
	if !success {
		return false
	}
	delete(gameMap, srcCellIdx)                    // set old location empty!
	gameMap[destination.GetCellIndex()] = piecePtr // fill in the new location
	return true
}
