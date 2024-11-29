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
	Id         int32     // unique piece id
	IsKing     bool      // whether this piece is King
	Pos        Vec2      // current piece position
	PieceColor PieceType // either RED or BLACK
}

// MoveSimple actually moves this piece diagonally to given destination by 1 cell. Returns TRUE if successful
func (p *Piece) MoveSimple(dest *Vec2) bool {
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
	if (p.PieceColor == Piece_Red && dest.Y == 0) ||
		(p.PieceColor == Piece_Black && dest.Y == 7*SIZE_CELL) {
		p.IsKing = true
	}
	return true
}

// MoveCapture (when attacking) moves this piece by 2 cells diagonally to the given `destination`. Returns TRUE if ok
func (p *Piece) MoveCapture(dest *Vec2) bool {
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
	if (p.PieceColor == Piece_Red && dest.Y == 0) ||
		(p.PieceColor == Piece_Black && dest.Y == 7*SIZE_CELL) {
		p.IsKing = true
	}
	return true
}

// IsEvenCellRow determines whether the CELL with given Index is on EVEN Row on the board
func IsEvenCellRow(cellIdx int32) bool {
	rowNumber := 9 - (cellIdx-1)/4
	return rowNumber%2 == 0
}

// IsAwayFromEdge returns TRUE if given position is NOT on any edge of board
func IsAwayFromEdge(pos *Vec2) bool {
	return pos.X > 0 && pos.X < 7*SIZE_CELL && pos.Y > 0 && pos.Y < 7*SIZE_CELL
}

// HasWinner returns TRUE if `p` has won the match against `opponent`, then notifies both players.
func HasWinner(p *player.Player, opponent *player.Player) bool {
	if len(opponent.Pieces) == 0 {
		// Meaning `opponent` has lost, `p` has won! Game over
		p.SendMessage(&BasePayload{
			Notice: "Congrats! You won! GAME OVER",
			Inner: &BasePayload_WinlosePayload{
				WinlosePayload: &WinLosePayload{
					Winner: TeamColor_TEAM_UNSPECIFIED,
				},
			},
		})
		opponent.SendMessage(&BasePayload{
			Notice: "Sorry! You lost! GAME OVER",
			Inner: &BasePayload_WinlosePayload{
				WinlosePayload: &WinLosePayload{
					Winner: TeamColor_TEAM_UNSPECIFIED,
				},
			},
		})
		log.Println("🏆 We got a winner!", p.Name, " has won!")
		return true
	}
	return false
}
