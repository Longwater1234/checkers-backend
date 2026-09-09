package room

import (
	"checkers-backend/game"
	"checkers-backend/player"
)

// processMovePiece made by Player `p` against `opponent`. Returns TRUE only if all is OK.
func processMovePiece(request *game.BasePayload, gameMap map[int32]*game.Piece, p, opponent *player.Player) bool {
	success := validateAndUpdateMap(request.GetMovePayload(), gameMap)
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
	// All is OK, forward the original request to opponent
	opponent.SendMessage(request)
	return success
}

// validateAndUpdateMap after player's piece MOVES and update gameMap. Returns TRUE only if successful
func validateAndUpdateMap(request *game.MovePayload, gameMap map[int32]*game.Piece) bool {
	destination := request.GetDestination()
	if destination == nil {
		return false
	}
	srcCellIdx := request.GetSourceCell()
	movingPieceId := request.GetPieceId()

	piecePtr, exists := gameMap[srcCellIdx]
	if !exists || movingPieceId != piecePtr.Id {
		return false
	}

	// check whether destCell already has a Piece
	_, hasValue := gameMap[destination.GetCellIndex()]
	if hasValue {
		return false
	}

	success := piecePtr.MoveSimple(game.Vec2{
		X: destination.GetX(),
		Y: destination.GetY(),
	})
	if !success {
		return false
	}
	delete(gameMap, srcCellIdx)                    // set old location empty!
	gameMap[destination.GetCellIndex()] = piecePtr // fill in the new location
	return success
}
