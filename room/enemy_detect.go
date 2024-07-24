package room

import (
	"checkers-backend/game"
	"checkers-backend/player"
	"log"
)

// hasExtraTargets returns TRUE if `hunter` has EXTRA nearby targets (ANY 4 sides) to capture (NOT entire board).
// This should be called only AFTER `handleCapture` by player `hunter` is valid
func hasExtraTargets(hunter *player.Player, currCell int32, gameMap map[int32]*game.Piece) bool {
	piecePtr, exists := gameMap[currCell]
	if !exists || !hunter.HasThisPiece(piecePtr.Id) {
		return false
	}
	if collectFrontLHS(hunter, currCell, gameMap) || collectFrontRHS(hunter, currCell, gameMap) {
		return true
	}

	if piecePtr.IsKing {
		if collectBehindLHS(hunter, currCell, gameMap) || collectBehindRHS(hunter, currCell, gameMap) {
			return true
		}
	}
	return false
}

// collectFrontLHS returns true ONLY IF there is an enemy on NorthWest of  player `p`
func collectFrontLHS(p *player.Player, cellIdx int32, gameMap map[int32]*game.Piece) bool {
	piecePtr := gameMap[cellIdx]
	//check LHS (north west)
	if p.Name == game.TeamColor_TEAM_RED.String() && piecePtr.Pos.X == 0 {
		return false
	}
	if p.Name == game.TeamColor_TEAM_BLACK.String() && piecePtr.Pos.X >= 7*game.SIZE_CELL {
		return false
	}
	var deltaForward int32 = 4
	var deltaBehindEnemy int32 = 5

	hasEnemyAhead := false
	enemyOpenBehind := false // does enemy piece have EMPTY cell behind it?

	if game.IsEvenCellRow(cellIdx) {
		deltaForward, deltaBehindEnemy = deltaBehindEnemy, deltaForward
	}
	var mSign int32 = +1 // direction. up +1, down -1

	// if player piece is Black (PLAYER 2)
	if p.Name == game.TeamColor_TEAM_BLACK.String() {
		mSign = -1
		deltaForward, deltaBehindEnemy = deltaBehindEnemy, deltaForward
	}
	var cellAheadIdx int32 = cellIdx + (deltaForward * mSign)
	if cellAheadIdx > 32 || cellAheadIdx < 1 {
		return false
	}

	pieceAhead, existFront := gameMap[cellAheadIdx] // north west (of hunter)
	hasEnemyAhead = existFront && !p.HasThisPiece(pieceAhead.Id)
	if existFront {
		log.Printf("pieceAhead %+v", pieceAhead)
	}
	if existFront && !game.IsAwayFromEdge(&pieceAhead.Pos) {
		return false
	}

	cellBehindEnemy := cellIdx + (deltaBehindEnemy * mSign) + (deltaForward * mSign) // south east (of enemy)
	if cellBehindEnemy > 32 || cellBehindEnemy < 1 {
		return false
	}

	// does enemy piece have EMPTY cell behind it?
	_, existBack := gameMap[cellBehindEnemy]
	enemyOpenBehind = !existBack
	log.Println("collectFrontLHS", "hasEnemyAhead", hasEnemyAhead, "enemyOpenBehind", enemyOpenBehind)
	return hasEnemyAhead && enemyOpenBehind
}

// collectFrontRHS returns true ONLY IF there is an enemy on NorthEast of this player `p`
func collectFrontRHS(p *player.Player, cellIdx int32, gameMap map[int32]*game.Piece) bool {
	piecePtr := gameMap[cellIdx]
	if p.Name == game.TeamColor_TEAM_RED.String() && piecePtr.Pos.X >= 7*game.SIZE_CELL {
		return false
	}
	if p.Name == game.TeamColor_TEAM_BLACK.String() && piecePtr.Pos.X == 0 {
		return false
	}
	var deltaForward int32 = 3
	var deltaBehindEnemy int32 = 4

	var hasEnemyAhead = false
	var enemyOpenBehind = false //is there an EMPTY cell behind enemy?

	if game.IsEvenCellRow(cellIdx) {
		deltaForward, deltaBehindEnemy = deltaBehindEnemy, deltaForward
	}
	var mSign int32 = +1 // direction. up +1, down -1

	// if piece is Black (PLAYER 2)
	if p.Name == game.TeamColor_TEAM_BLACK.String() {
		mSign = -1
		deltaForward, deltaBehindEnemy = deltaBehindEnemy, deltaForward
	}
	var cellAheadIdx int32 = cellIdx + (deltaForward * mSign)
	if cellAheadIdx > 32 || cellAheadIdx < 1 {
		return false
	}

	pieceAhead, existFront := gameMap[cellAheadIdx] // north east
	hasEnemyAhead = existFront && !p.HasThisPiece(pieceAhead.Id)
	if existFront {
		log.Printf("pieceAhead %+v", pieceAhead)
	}
	if existFront && !game.IsAwayFromEdge(&pieceAhead.Pos) {
		return false
	}

	cellBehindEnemy := cellIdx + (deltaBehindEnemy * mSign) + (deltaForward * mSign) // south west (of enemy)
	if cellBehindEnemy > 32 || cellBehindEnemy < 1 {
		return false
	}
	// does enemy piece have EMPTY cell behind it?
	_, existBack := gameMap[cellBehindEnemy]
	enemyOpenBehind = !existBack
	log.Println("collectFrontRHS", "hasEnemyAhead", hasEnemyAhead, "enemyOpenBehind", enemyOpenBehind)
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
	var deltaForward int32 = 5
	var deltaBehindEnemy int32 = 4

	var hasEnemyAhead = false
	var enemyOpenBehind = false //is there an EMPTY cell behind enemy?

	if game.IsEvenCellRow(cellIdx) {
		deltaForward, deltaBehindEnemy = deltaBehindEnemy, deltaForward
	}
	var mSign int32 = +1 // direction

	// if player piece is Black (PLAYER 2)
	if king.Name == game.TeamColor_TEAM_BLACK.String() {
		mSign = -1
		deltaForward, deltaBehindEnemy = deltaBehindEnemy, deltaForward
	}
	var cellAheadIdx int32 = cellIdx - (deltaForward * mSign)
	if cellAheadIdx > 32 || cellAheadIdx < 1 {
		return false
	}

	pieceAhead, existFront := gameMap[cellAheadIdx] // north west (opposite direction)
	hasEnemyAhead = existFront && !king.HasThisPiece(pieceAhead.Id)
	if existFront {
		log.Printf("pieceAhead %+v", pieceAhead)
	}
	if existFront && !game.IsAwayFromEdge(&pieceAhead.Pos) {
		return false
	}

	cellBehindEnemy := cellIdx - (deltaBehindEnemy * mSign) - (deltaForward * mSign) // south east (of enemy)
	if cellBehindEnemy > 32 || cellBehindEnemy < 1 {
		return false
	}
	// does enemy piece have EMPTY cell behind it?
	_, existBack := gameMap[cellBehindEnemy]
	enemyOpenBehind = !existBack
	log.Println("collectBehindRHS", "hasEnemyAhead", hasEnemyAhead, "enemyOpenBehind", enemyOpenBehind)
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
	var deltaForward int32 = 4
	var deltaBehindEnemy int32 = 3

	var hasEnemyAhead = false
	var enemyOpenBehind = false //is there an EMPTY cell behind enemy?

	if game.IsEvenCellRow(cellIdx) {
		deltaForward, deltaBehindEnemy = deltaBehindEnemy, deltaForward
	}
	var mSign int32 = +1 // direction

	// if player piece is Black (PLAYER 2)
	if king.Name == game.TeamColor_TEAM_BLACK.String() {
		mSign = -1
		deltaForward, deltaBehindEnemy = deltaBehindEnemy, deltaForward
	}
	var cellAheadIdx int32 = cellIdx - (deltaForward * mSign)
	if cellAheadIdx > 32 || cellAheadIdx < 1 {
		return false
	}

	pieceAhead, existFront := gameMap[cellAheadIdx] // north east (opposite direction)
	hasEnemyAhead = existFront && !king.HasThisPiece(pieceAhead.Id)
	if existFront {
		log.Printf("pieceAhead %+v", pieceAhead)
	}
	if existFront && !game.IsAwayFromEdge(&pieceAhead.Pos) {
		return false
	}

	cellBehindEnemy := cellIdx - (deltaBehindEnemy * mSign) - (deltaForward * mSign) // south west of enemy
	if cellBehindEnemy > 32 || cellBehindEnemy < 1 {
		return false
	}
	// does enemy piece have EMPTY cell behind it?
	_, existBack := gameMap[cellBehindEnemy]
	enemyOpenBehind = !existBack
	log.Println("collectBehindLHS", "hasEnemyAhead", hasEnemyAhead, "enemyOpenBehind", enemyOpenBehind)
	return hasEnemyAhead && enemyOpenBehind
}
