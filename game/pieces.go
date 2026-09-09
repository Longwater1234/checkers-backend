package game

import (
	"checkers-backend/player"
	"log"
	"math"
)

const SIZE_CELL = 75.0 // length of single square cell

type Vec2 struct {
	X float32 // x position
	Y float32 // y position
}

type PieceType int16

const (
	Piece_Red PieceType = iota + 1
	Piece_Black
)

type Piece struct {
	Id         int32     // random unique piece id
	IsKing     bool      // whether this piece is King
	Pos        Vec2      // piece current position
	PieceColor PieceType // either RED or BLACK
}

// MoveSimple moves this piece diagonally to given destination by 1 cell. Returns TRUE if successful
func (p *Piece) MoveSimple(dest Vec2) bool {
	var deltaX = float64(dest.X - p.Pos.X)
	var deltaY = float64(dest.Y - p.Pos.Y)

	if math.Abs(deltaX) != SIZE_CELL || math.Abs(deltaY) != SIZE_CELL {
		return false
	}
	if p.PieceColor == Piece_Red && deltaY > 0 && !p.IsKing {
		return false
	}
	if p.PieceColor == Piece_Black && deltaY < 0 && !p.IsKing {
		return false
	}

	p.Pos.X = dest.X
	p.Pos.Y = dest.Y
	// activate King if reached opposite end
	if (p.PieceColor == Piece_Red && dest.Y == 0) ||
		(p.PieceColor == Piece_Black && dest.Y == 7*SIZE_CELL) {
		p.IsKing = true
	}
	return true
}

// MoveCapture (when attacking) moves this piece diagonally by 2 cells to the given `destination`. Returns TRUE if success
func (p *Piece) MoveCapture(dest Vec2) bool {
	var deltaX = float64(dest.X - p.Pos.X)
	var deltaY = float64(dest.Y - p.Pos.Y)

	if math.Abs(deltaX) != 2*SIZE_CELL || math.Abs(deltaY) != 2*SIZE_CELL {
		return false
	}
	if p.PieceColor == Piece_Red && deltaY > 0 && !p.IsKing {
		return false
	}
	if p.PieceColor == Piece_Black && deltaY < 0 && !p.IsKing {
		return false
	}

	p.Pos.X = dest.X
	p.Pos.Y = dest.Y
	// activate King if reached opposite end
	if (p.PieceColor == Piece_Red && dest.Y == 0) ||
		(p.PieceColor == Piece_Black && dest.Y == 7*SIZE_CELL) {
		p.IsKing = true
	}
	return true
}

// canMoveLegally returns TRUE if given piece has at least 1 valid move or capture available
func (p *Piece) canMoveLegally() bool {
	if p == nil {
		return false
	}
	// check all 4 diagonal directions for valid moves
	directions := []Vec2{
		{X: -SIZE_CELL, Y: -SIZE_CELL}, // up-left
		{X: SIZE_CELL, Y: -SIZE_CELL},  // up-right
		{X: -SIZE_CELL, Y: SIZE_CELL},  // down-left
		{X: SIZE_CELL, Y: SIZE_CELL},   // down-right
	}

	for _, dir := range directions {
		dest := Vec2{X: p.Pos.X + dir.X, Y: p.Pos.Y + dir.Y}
		if p.MoveSimple(dest) {
			return true
		}
		destCapture := Vec2{X: p.Pos.X + 2*dir.X, Y: p.Pos.Y + 2*dir.Y}
		if p.MoveCapture(destCapture) {
			return true
		}
	}
	return false
}

// IsEvenCellRow determines whether the CELL with given Index is on EVEN Row on the board
func IsEvenCellRow(cellIdx int32) bool {
	rowNum := (32 - cellIdx) / 4
	return rowNum%2 == 0
}

// IsAwayFromEdge returns TRUE if given position is NOT on any edge of board
func IsAwayFromEdge(pos Vec2) bool {
	return pos.X > 0 && pos.X < 7*SIZE_CELL && pos.Y > 0 && pos.Y < 7*SIZE_CELL
}

// HasWinner returns TRUE if player `p` has beaten `opponent`, and then notifies both players if so.
// This must be called after the current player has completed their turn, and the gameMap is updated with the latest piece positions.
func HasWinner(p *player.Player, opponent *player.Player, gameMap map[int32]*Piece) bool {
	if len(opponent.Pieces) == 0 {
		// Meaning `opponent` has lost, `p` has won! Game over
		notifyLoserWinner(p, opponent)
		return true
	} else if len(opponent.Pieces) <= 2 && !nextPlayerHasValidMoves(opponent, gameMap) {
		notifyLoserWinner(p, opponent)
		return true
	}
	return false
}

// notifyLoserWinner sends WIN/LOSE message to both players and logs the winner
func notifyLoserWinner(p *player.Player, opponent *player.Player) {
	var winner TeamColor = TeamColor_TEAM_RED
	if p.Name == TeamColor_TEAM_BLACK.String() {
		winner = TeamColor_TEAM_BLACK
	}
	p.SendMessage(&BasePayload{
		Notice: "Congrats! You won! GAME OVER",
		Inner: &BasePayload_WinlosePayload{
			WinlosePayload: &WinLosePayload{
				Winner: winner,
			},
		},
	})
	opponent.SendMessage(&BasePayload{
		Notice: "Sorry! You lost! GAME OVER",
		Inner: &BasePayload_WinlosePayload{
			WinlosePayload: &WinLosePayload{
				Winner: winner,
			},
		},
	})
	log.Println("🏆 We got a winner!", p.Name, " has won!")
}

// nextPlayerHasValidMoves returns TRUE if the next player has at least 1 valid move available.
func nextPlayerHasValidMoves(p *player.Player, gameMap map[int32]*Piece) bool {
	if p == nil || len(p.Pieces) == 0 {
		return false
	}

	// check if player has 1 or 2 pieces left, if not then they have valid moves available
	if len(p.Pieces) > 2 {
		return false
	}

	for _, piece := range gameMap {
		if piece.Id == p.Pieces[0] || (len(p.Pieces) > 1 && piece.Id == p.Pieces[1]) {
			if ok := piece.canMoveLegally(); ok {
				return true
			}
		}
	}
	return false
}
