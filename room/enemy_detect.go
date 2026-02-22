package room

import (
	"checkers-backend/game"
	"checkers-backend/player"
)

// hasExtraTargets returns TRUE if hunter's single Piece at `cell_idx` has EXTRA nearby enemies to capture.
// This should be called only AFTER `processCapturePiece` by hunter is TRUE
func hasExtraTargets(hunter *player.Player, cellIdx int32, gameMap map[int32]*game.Piece) bool {
	piecePtr, exists := gameMap[cellIdx]
	if !exists || !hunter.HasThisPiece(piecePtr.Id) {
		return false
	}

	if collectFrontLHS(hunter, cellIdx, gameMap) || collectFrontRHS(hunter, cellIdx, gameMap) {
		return true
	}

	if piecePtr.IsKing {
		if collectBehindLHS(hunter, cellIdx, gameMap) || collectBehindRHS(hunter, cellIdx, gameMap) {
			return true
		}
	}
	return false
}

// collectFrontLHS returns true ONLY IF there is an enemy on NorthWest of  player `p` at `cellIdx`
func collectFrontLHS(p *player.Player, cellIdx int32, gameMap map[int32]*game.Piece) bool {
	piecePtr := gameMap[cellIdx]
	//check LHS (north west)
	if p.Name == game.TeamColor_TEAM_RED.String() && piecePtr.Pos.X == 0 {
		return false
	}
	if p.Name == game.TeamColor_TEAM_BLACK.String() && piecePtr.Pos.X >= 7*game.SIZE_CELL {
		return false
	}
	var deltaForward int32 = 5
	var deltaBehindEnemy int32 = 4

	var hasEnemyAhead = false
	var enemyOpenBehind = false // have EMPTY cell behind enemy?

	if game.IsEvenCellRow(cellIdx) {
		deltaForward, deltaBehindEnemy = deltaBehindEnemy, deltaForward // swap values
	}
	var direction int32 = +1 // up +1, down -1

	// if player piece is Black (PLAYER 2), swap values
	if p.Name == game.TeamColor_TEAM_BLACK.String() {
		direction = -1
		deltaBehindEnemy, deltaForward = deltaForward, deltaBehindEnemy
	}
	var cellAheadIdx int32 = cellIdx + (deltaForward * direction)
	if cellAheadIdx > 32 || cellAheadIdx < 1 {
		return false
	}

	pieceAhead, existFront := gameMap[cellAheadIdx] // north-west (of hunter)
	hasEnemyAhead = existFront && !p.HasThisPiece(pieceAhead.Id)
	if existFront && !game.IsAwayFromEdge(pieceAhead.Pos) {
		return false
	}

	cellBehindEnemy := cellIdx + (deltaBehindEnemy * direction) + (deltaForward * direction) // south-east (of enemy)
	if cellBehindEnemy > 32 || cellBehindEnemy < 1 {
		return false
	}
	// does enemy piece have EMPTY cell behind?
	_, existBack := gameMap[cellBehindEnemy]
	enemyOpenBehind = !existBack
	return hasEnemyAhead && enemyOpenBehind
}

// collectFrontRHS returns true ONLY IF there is an enemy on NorthEast of this player `p` at `cellIdx`
func collectFrontRHS(p *player.Player, cellIdx int32, gameMap map[int32]*game.Piece) bool {
	piecePtr := gameMap[cellIdx]
	if p.Name == game.TeamColor_TEAM_RED.String() && piecePtr.Pos.X >= 7*game.SIZE_CELL {
		return false
	}
	if p.Name == game.TeamColor_TEAM_BLACK.String() && piecePtr.Pos.X == 0 {
		return false
	}
	var deltaForward int32 = 4
	var deltaBehindEnemy int32 = 3

	var hasEnemyAhead = false
	var enemyOpenBehind = false // is there an EMPTY cell behind enemy?

	if game.IsEvenCellRow(cellIdx) {
		deltaBehindEnemy, deltaForward = deltaForward, deltaBehindEnemy // do swap
	}
	var direction int32 = +1 // up +1, down -1

	// if piece is Black (PLAYER 2), swap values
	if p.Name == game.TeamColor_TEAM_BLACK.String() {
		direction = -1
		deltaBehindEnemy, deltaForward = deltaForward, deltaBehindEnemy
	}
	var cellAheadIdx int32 = cellIdx + (deltaForward * direction)
	if cellAheadIdx > 32 || cellAheadIdx < 1 {
		return false
	}

	pieceAhead, existFront := gameMap[cellAheadIdx] // north-east of hunter
	hasEnemyAhead = existFront && !p.HasThisPiece(pieceAhead.Id)
	if existFront && !game.IsAwayFromEdge(pieceAhead.Pos) {
		return false
	}

	cellBehindEnemy := cellIdx + (deltaBehindEnemy * direction) + (deltaForward * direction) // south-west (of enemy)
	if cellBehindEnemy > 32 || cellBehindEnemy < 1 {
		return false
	}
	// does enemy piece have EMPTY cell behind?
	_, existBack := gameMap[cellBehindEnemy]
	enemyOpenBehind = !existBack
	return hasEnemyAhead && enemyOpenBehind
}

// collectBehindRHS returns true ONLY IF there is an enemy on SOUTH EAST of piece. (Only for KING pieces)
func collectBehindRHS(king *player.Player, cellIdx int32, gameMap map[int32]*game.Piece) bool {
	piecePtr := gameMap[cellIdx]
	if king.Name == game.TeamColor_TEAM_RED.String() && piecePtr.Pos.X == 0 {
		return false
	}
	if king.Name == game.TeamColor_TEAM_BLACK.String() && piecePtr.Pos.X > 7*game.SIZE_CELL {
		return false
	}
	var deltaForward int32 = 4
	var deltaBehindEnemy int32 = 5

	var hasEnemyAhead = false
	var enemyOpenBehind = false // is there an EMPTY cell behind enemy?

	if game.IsEvenCellRow(cellIdx) {
		deltaForward, deltaBehindEnemy = deltaBehindEnemy, deltaForward // do swap
	}
	var direction int32 = +1 // for King, it's reversed

	// if player piece is Black (PLAYER 2), swap values
	if king.Name == game.TeamColor_TEAM_BLACK.String() {
		direction = -1
		deltaForward, deltaBehindEnemy = deltaBehindEnemy, deltaForward
	}
	var cellAheadIdx int32 = cellIdx - (deltaForward * direction)
	if cellAheadIdx > 32 || cellAheadIdx < 1 {
		return false
	}

	pieceAhead, existFront := gameMap[cellAheadIdx] // north-west (when facing behind)
	hasEnemyAhead = existFront && !king.HasThisPiece(pieceAhead.Id)
	if existFront && !game.IsAwayFromEdge(pieceAhead.Pos) {
		return false
	}

	cellBehindEnemy := cellIdx - (deltaBehindEnemy * direction) - (deltaForward * direction) // south-east (of enemy)
	if cellBehindEnemy > 32 || cellBehindEnemy < 1 {
		return false
	}
	// does enemy piece have EMPTY cell behind?
	_, existBack := gameMap[cellBehindEnemy]
	enemyOpenBehind = !existBack
	return hasEnemyAhead && enemyOpenBehind
}

// hasAnyPossibleMoves returns TRUE if player `p` has at least one valid move (simple or capture) across ALL their pieces.
// Call this BEFORE waiting for a player's move to detect a blocked/no-moves situation.
func hasAnyPossibleMoves(p *player.Player, gameMap map[int32]*game.Piece) bool {
	for cellIdx, piece := range gameMap {
		if !p.HasThisPiece(piece.Id) {
			continue
		}
		if canSimpleMoveFrontLHS(p, cellIdx, gameMap) || canSimpleMoveFrontRHS(p, cellIdx, gameMap) {
			return true
		}
		if piece.IsKing {
			if canSimpleMoveBackLHS(p, cellIdx, gameMap) || canSimpleMoveBackRHS(p, cellIdx, gameMap) {
				return true
			}
		}
		if collectFrontLHS(p, cellIdx, gameMap) || collectFrontRHS(p, cellIdx, gameMap) {
			return true
		}
		if piece.IsKing {
			if collectBehindLHS(p, cellIdx, gameMap) || collectBehindRHS(p, cellIdx, gameMap) {
				return true
			}
		}
	}
	return false
}

// canSimpleMoveFrontLHS returns true if piece at cellIdx can make a simple 1-step move to the front-left diagonal
func canSimpleMoveFrontLHS(p *player.Player, cellIdx int32, gameMap map[int32]*game.Piece) bool {
	piecePtr := gameMap[cellIdx]
	if p.Name == game.TeamColor_TEAM_RED.String() && piecePtr.Pos.X == 0 {
		return false
	}
	if p.Name == game.TeamColor_TEAM_BLACK.String() && piecePtr.Pos.X >= 7*game.SIZE_CELL {
		return false
	}
	var deltaForward int32 = 5
	if game.IsEvenCellRow(cellIdx) {
		deltaForward = 4
	}
	direction := int32(+1)
	if p.Name == game.TeamColor_TEAM_BLACK.String() {
		direction = -1
		if game.IsEvenCellRow(cellIdx) {
			deltaForward = 5
		} else {
			deltaForward = 4
		}
	}
	cellAheadIdx := cellIdx + (deltaForward * direction)
	if cellAheadIdx > 32 || cellAheadIdx < 1 {
		return false
	}
	_, exists := gameMap[cellAheadIdx]
	return !exists
}

// canSimpleMoveFrontRHS returns true if piece at cellIdx can make a simple 1-step move to the front-right diagonal
func canSimpleMoveFrontRHS(p *player.Player, cellIdx int32, gameMap map[int32]*game.Piece) bool {
	piecePtr := gameMap[cellIdx]
	if p.Name == game.TeamColor_TEAM_RED.String() && piecePtr.Pos.X >= 7*game.SIZE_CELL {
		return false
	}
	if p.Name == game.TeamColor_TEAM_BLACK.String() && piecePtr.Pos.X == 0 {
		return false
	}
	var deltaForward int32 = 4
	if game.IsEvenCellRow(cellIdx) {
		deltaForward = 3
	}
	direction := int32(+1)
	if p.Name == game.TeamColor_TEAM_BLACK.String() {
		direction = -1
		if game.IsEvenCellRow(cellIdx) {
			deltaForward = 4
		} else {
			deltaForward = 3
		}
	}
	cellAheadIdx := cellIdx + (deltaForward * direction)
	if cellAheadIdx > 32 || cellAheadIdx < 1 {
		return false
	}
	_, exists := gameMap[cellAheadIdx]
	return !exists
}

// canSimpleMoveBackLHS returns true if King piece at cellIdx can make a simple 1-step backward move to back-left diagonal
func canSimpleMoveBackLHS(king *player.Player, cellIdx int32, gameMap map[int32]*game.Piece) bool {
	piecePtr := gameMap[cellIdx]
	if king.Name == game.TeamColor_TEAM_RED.String() && piecePtr.Pos.X > 7*game.SIZE_CELL {
		return false
	}
	if king.Name == game.TeamColor_TEAM_BLACK.String() && piecePtr.Pos.X == 0 {
		return false
	}
	var deltaForward int32 = 3
	if game.IsEvenCellRow(cellIdx) {
		deltaForward = 4
	}
	direction := int32(+1)
	if king.Name == game.TeamColor_TEAM_BLACK.String() {
		direction = -1
		if game.IsEvenCellRow(cellIdx) {
			deltaForward = 3
		} else {
			deltaForward = 4
		}
	}
	cellBehindIdx := cellIdx - (deltaForward * direction)
	if cellBehindIdx > 32 || cellBehindIdx < 1 {
		return false
	}
	_, exists := gameMap[cellBehindIdx]
	return !exists
}

// canSimpleMoveBackRHS returns true if King piece at cellIdx can make a simple 1-step backward move to back-right diagonal
func canSimpleMoveBackRHS(king *player.Player, cellIdx int32, gameMap map[int32]*game.Piece) bool {
	piecePtr := gameMap[cellIdx]
	if king.Name == game.TeamColor_TEAM_RED.String() && piecePtr.Pos.X == 0 {
		return false
	}
	if king.Name == game.TeamColor_TEAM_BLACK.String() && piecePtr.Pos.X > 7*game.SIZE_CELL {
		return false
	}
	var deltaForward int32 = 4
	if game.IsEvenCellRow(cellIdx) {
		deltaForward = 5
	}
	direction := int32(+1)
	if king.Name == game.TeamColor_TEAM_BLACK.String() {
		direction = -1
		if game.IsEvenCellRow(cellIdx) {
			deltaForward = 4
		} else {
			deltaForward = 5
		}
	}
	cellBehindIdx := cellIdx - (deltaForward * direction)
	if cellBehindIdx > 32 || cellBehindIdx < 1 {
		return false
	}
	_, exists := gameMap[cellBehindIdx]
	return !exists
}

// collectBehindLHS returns true ONLY IF there is an enemy on SOUTH WEST of player. Only for KING pieces
func collectBehindLHS(king *player.Player, cellIdx int32, gameMap map[int32]*game.Piece) bool {
	piecePtr := gameMap[cellIdx]
	if king.Name == game.TeamColor_TEAM_RED.String() && piecePtr.Pos.X > 7*game.SIZE_CELL {
		return false
	}
	if king.Name == game.TeamColor_TEAM_BLACK.String() && piecePtr.Pos.X == 0 {
		return false
	}
	var deltaForward int32 = 3
	var deltaBehindEnemy int32 = 4

	var hasEnemyAhead = false
	var enemyOpenBehind = false // have EMPTY cell behind enemy?

	if game.IsEvenCellRow(cellIdx) {
		deltaForward, deltaBehindEnemy = deltaBehindEnemy, deltaForward
	}
	var direction int32 = +1 // for King it's reversed

	// if player piece is Black (PLAYER 2), do swap
	if king.Name == game.TeamColor_TEAM_BLACK.String() {
		direction = -1
		deltaForward, deltaBehindEnemy = deltaBehindEnemy, deltaForward
	}
	var cellAheadIdx int32 = cellIdx - (deltaForward * direction)
	if cellAheadIdx > 32 || cellAheadIdx < 1 {
		return false
	}

	pieceAhead, existFront := gameMap[cellAheadIdx] // north-east (when facing behind)
	hasEnemyAhead = existFront && !king.HasThisPiece(pieceAhead.Id)
	if existFront && !game.IsAwayFromEdge(pieceAhead.Pos) {
		return false
	}

	cellBehindEnemy := cellIdx - (deltaBehindEnemy * direction) - (deltaForward * direction) // south-west of enemy
	if cellBehindEnemy > 32 || cellBehindEnemy < 1 {
		return false
	}
	// does enemy piece have EMPTY cell behind?
	_, existBack := gameMap[cellBehindEnemy]
	enemyOpenBehind = !existBack
	return hasEnemyAhead && enemyOpenBehind
}
