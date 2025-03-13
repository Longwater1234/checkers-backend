package room

import (
	"checkers-backend/game"
	"checkers-backend/player"
)

// hasExtraTargets returns TRUE if hunter's single Piece at `cellIdx` has EXTRA nearby enemies to capture.
// This should be called only AFTER `handleCapture` by player `hunter` return TRUE
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
	if existFront && !game.IsAwayFromEdge(&pieceAhead.Pos) {
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

	pieceAhead, existFront := gameMap[cellAheadIdx] // north-east
	hasEnemyAhead = existFront && !p.HasThisPiece(pieceAhead.Id)
	if existFront && !game.IsAwayFromEdge(&pieceAhead.Pos) {
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

// collectBehindRHS returns true ONLY IF there is an enemy on (SOUTH EAST) of piece. (Only for KING pieces)
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
	if existFront && !game.IsAwayFromEdge(&pieceAhead.Pos) {
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

// collectBehindLHS returns true ONLY IF there is an enemy on (SOUTH WEST) of player. Only for KING pieces
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
	if existFront && !game.IsAwayFromEdge(&pieceAhead.Pos) {
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
