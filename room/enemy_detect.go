package room

import (
	"checkers-backend/game"
	"checkers-backend/player"
	"log"
)

// isEvenCellRow determines wheter given cell_index is on even Row
func isEvenCellRow(cellIdx int32) bool {
	rowNumber := 9 - (cellIdx-1)/4
	return (rowNumber%2 == 0)
}

// awayFromEdge returns TRUE if given position is NOT on any edge of board
func awayFromEdge(pos *game.Vec2) bool {
	return pos.X > 0 && pos.X < 7*game.SIZE_CELL && pos.Y > 0 && pos.Y < 7*game.SIZE_CELL
}

// hasExtraTargets returns TRUE if `p` has extra nearby targets (4 sides of currentCell) to capture. (NOT entire board)
// This should be called ONLY AFTER `handleCapture` by player `p` succeeds
func hasExtraTargets(p *player.Player, currCell int32, gameMap map[int32]*game.Piece) bool {
	//var hasExtra = false
	piecePtr, exists := gameMap[currCell]
	if !exists || !p.HasThisPiece(piecePtr.Id) {
		return false
	}
	if collectFrontLHS(p, currCell, gameMap) || collectFrontRHS(p, currCell, gameMap) {
		return true
	}

	if piecePtr.IsKing {
		if collectBehindLHS(p, currCell, gameMap) || collectBehindRHS(p, currCell, gameMap) {
			return true
		}
	}
	return false
}

// collectFrontLHS checks whether there is an enemy on NorthWest of this player `p`
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

	if isEvenCellRow(cellIdx) {
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

	pieceAhead, existFront := gameMap[cellAheadIdx] // north west
	hasEnemyAhead = existFront && !p.HasThisPiece(pieceAhead.Id)
	if existFront && !awayFromEdge(&pieceAhead.Pos) {
		return false
	}

	cellBehindEnemy := cellIdx + (deltaBehindEnemy * mSign) + (deltaForward * mSign) // south east
	if cellBehindEnemy > 32 || cellBehindEnemy < 1 {
		return false
	}

	// does enemy piece have EMPTY cell behind it?
	_, existBack := gameMap[cellBehindEnemy]
	enemyOpenBehind = !existBack
	log.Println("hasEnemyAhead", hasEnemyAhead, "enemyOpenBehind", enemyOpenBehind)
	return hasEnemyAhead && enemyOpenBehind
}

// collectFrontRHS checks whether there is an enemy on NorthEast of this player `p`
func collectFrontRHS(p *player.Player, cellIdx int32, gameMap map[int32]*game.Piece) bool {
	piecePtr := gameMap[cellIdx]
	if p.Name == game.TeamColor_TEAM_RED.String() && piecePtr.Pos.X == 0 {
		return false
	}
	if p.Name == game.TeamColor_TEAM_BLACK.String() && piecePtr.Pos.X >= 7*game.SIZE_CELL {
		return false
	}
	var deltaForward int32 = 3
	var deltaBehindEnemy int32 = 4

	hasEnemyAhead := false
	enemyOpenBehind := false // does enemy piece have EMPTY cell behind it?

	if isEvenCellRow(cellIdx) {
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

	pieceAhead, existFront := gameMap[cellAheadIdx] // north east
	hasEnemyAhead = existFront && !p.HasThisPiece(pieceAhead.Id)
	if existFront && !awayFromEdge(&pieceAhead.Pos) {
		return false
	}

	cellBehindEnemy := cellIdx + (deltaBehindEnemy * mSign) + (deltaForward * mSign) // south west
	if cellBehindEnemy > 32 || cellBehindEnemy < 1 {
		return false
	}
	// does enemy piece have EMPTY cell behind it?
	_, existBack := gameMap[cellBehindEnemy]
	enemyOpenBehind = !existBack
	return hasEnemyAhead && enemyOpenBehind
}

// collectBehindRHS checks whether there is an enemy on (SOUTH EAST). Only for KING pieces
func collectBehindRHS(p *player.Player, cellIdx int32, gameMap map[int32]*game.Piece) bool {
	piecePtr := gameMap[cellIdx]
	if p.Name == game.TeamColor_TEAM_RED.String() && piecePtr.Pos.X > 7*game.SIZE_CELL {
		return false
	}
	if p.Name == game.TeamColor_TEAM_BLACK.String() && piecePtr.Pos.X == 0 {
		return false
	}
	var deltaForward int32 = 5
	var deltaBehindEnemy int32 = 4

	hasEnemyAhead := false
	enemyOpenBehind := false // does enemy piece have EMPTY cell behind it?

	if isEvenCellRow(cellIdx) {
		deltaForward, deltaBehindEnemy = deltaBehindEnemy, deltaForward
	}
	var mSign int32 = +1 // direction. up +1, down -1

	// if player piece is Black (PLAYER 2)
	if p.Name == game.TeamColor_TEAM_BLACK.String() {
		mSign = -1
		deltaForward, deltaBehindEnemy = deltaBehindEnemy, deltaForward
	}
	var cellAheadIdx int32 = cellIdx - (deltaForward * mSign)
	if cellAheadIdx > 32 || cellAheadIdx < 1 {
		return false
	}

	pieceAhead, existFront := gameMap[cellAheadIdx] // north east
	hasEnemyAhead = existFront && !p.HasThisPiece(pieceAhead.Id)
	if existFront && !awayFromEdge(&pieceAhead.Pos) {
		return false
	}

	cellBehindEnemy := cellIdx - (deltaBehindEnemy * mSign) - (deltaForward * mSign) // south west
	if cellBehindEnemy > 32 || cellBehindEnemy < 1 {
		return false
	}
	// does enemy piece have EMPTY cell behind it?
	_, existBack := gameMap[cellBehindEnemy]
	enemyOpenBehind = !existBack
	return hasEnemyAhead && enemyOpenBehind
}

// collectBehindLHS checks whether there is an enemy on (SOUTH WEST). Only for KING pieces
func collectBehindLHS(p *player.Player, cellIdx int32, gameMap map[int32]*game.Piece) bool {
	piecePtr := gameMap[cellIdx]
	if p.Name == game.TeamColor_TEAM_RED.String() && piecePtr.Pos.X > 7*game.SIZE_CELL {
		return false
	}
	if p.Name == game.TeamColor_TEAM_BLACK.String() && piecePtr.Pos.X == 0 {
		return false
	}
	var deltaForward int32 = 4
	var deltaBehindEnemy int32 = 3

	hasEnemyAhead := false
	enemyOpenBehind := false // does enemy piece have EMPTY cell behind it?

	if isEvenCellRow(cellIdx) {
		deltaForward, deltaBehindEnemy = deltaBehindEnemy, deltaForward
	}
	var mSign int32 = +1 // direction. up +1, down -1

	// if player piece is Black (PLAYER 2)
	if p.Name == game.TeamColor_TEAM_BLACK.String() {
		mSign = -1
		deltaForward, deltaBehindEnemy = deltaBehindEnemy, deltaForward
	}
	var cellAheadIdx int32 = cellIdx - (deltaForward * mSign)
	if cellAheadIdx > 32 || cellAheadIdx < 1 {
		return false
	}

	pieceAhead, existFront := gameMap[cellAheadIdx] // north east
	hasEnemyAhead = existFront && !p.HasThisPiece(pieceAhead.Id)
	if existFront && !awayFromEdge(&pieceAhead.Pos) {
		return false
	}

	cellBehindEnemy := cellIdx - (deltaBehindEnemy * mSign) - (deltaForward * mSign) // south west
	if cellBehindEnemy > 32 || cellBehindEnemy < 1 {
		return false
	}
	// does enemy piece have EMPTY cell behind it?
	_, existBack := gameMap[cellBehindEnemy]
	enemyOpenBehind = !existBack
	return hasEnemyAhead && enemyOpenBehind
}
