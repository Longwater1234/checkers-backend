package room

import (
	"checkers-backend/game"
	"checkers-backend/player"
)

// processMovePiece made by Player `p` against `opponent`. Returns TRUE only if all is OK.
func processMovePiece(request *game.BasePayload, gameMap map[int32]*game.Piece, p, opponent *player.Player) bool {
	success := validateAndUpdateMap(request.GetMovePayload(), gameMap, p)
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

// validateAndUpdateMap after player `p`'s piece MOVES and update gameMap. Returns TRUE only if successful
func validateAndUpdateMap(request *game.MovePayload, gameMap map[int32]*game.Piece, p *player.Player) bool {
	destination := request.GetDestination()
	if destination == nil {
		return false
	}
	srcCellIdx := request.GetSourceCell()
	movingPieceId := request.GetPieceId()

	piecePtr, exists := gameMap[srcCellIdx]
	if !exists || movingPieceId != piecePtr.Id || !p.HasThisPiece(movingPieceId) {
		return false
	}

	// never trust client's cell_index: derive it from the (x,y) position
	destPos := game.Vec2{X: destination.GetX(), Y: destination.GetY()}
	destCellIdx := game.CellIndex(destPos)
	if destCellIdx == 0 {
		return false // off-board, or not a playable cell
	}
	if _, occupied := gameMap[destCellIdx]; occupied {
		return false
	}

	if !piecePtr.MoveSimple(destPos) {
		return false
	}
	destination.CellIndex = destCellIdx // so the forwarded request can't desync the opponent
	delete(gameMap, srcCellIdx)         // set old location empty!
	gameMap[destCellIdx] = piecePtr     // fill in the new location
	return true
}
