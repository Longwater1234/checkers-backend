package game

import (
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

// directions is fixed array of possible positions a [Piece] can legally move to.
var directions = [4]Vec2{
	{X: -SIZE_CELL, Y: -SIZE_CELL}, // up-left
	{X: SIZE_CELL, Y: -SIZE_CELL},  // up-right
	{X: -SIZE_CELL, Y: SIZE_CELL},  // down-left
	{X: SIZE_CELL, Y: SIZE_CELL},   // down-right
}

// canMoveLegally returns TRUE if given piece has at least 1 valid move or capture available
func (p *Piece) canMoveLegally(gameMap map[int32]*Piece) bool {
	if p == nil || gameMap == nil {
		return false
	}
	dirs := directions[:] // for king pieces
	if !p.IsKing {
		if p.PieceColor == Piece_Red {
			dirs = directions[:2]
		} else {
			dirs = directions[2:]
		}
	}

	for _, dir := range dirs {
		dest := Vec2{
			X: p.Pos.X + dir.X,
			Y: p.Pos.Y + dir.Y,
		}
		destCellIdx := getCellIndex(dest)
		if destCellIdx < 1 || destCellIdx > 32 {
			continue
		}

		prey, occupied := gameMap[destCellIdx]
		if !occupied || prey == nil {
			return true // empty adjacent cell, simple move available
		}

		// If occupied by opponent, check if capture is possible
		if prey.PieceColor != p.PieceColor {
			destCapture := Vec2{
				X: p.Pos.X + 2*dir.X,
				Y: p.Pos.Y + 2*dir.Y,
			}
			captureCellIdx := getCellIndex(destCapture)
			if captureCellIdx >= 1 && captureCellIdx <= 32 {
				landingPiece, landOccupied := gameMap[captureCellIdx]
				if !landOccupied || landingPiece == nil {
					return true // valid jump over enemy into empty cell
				}
			}
		}
	}
	return false
}

// getCellIndex converts a board coordinate (Vec2) into a 1-based playable cell index (1..32).
// Returns 0 if the position is off-board or not a playable cell.
func getCellIndex(dest Vec2) int32 {
	col := int(math.Round(float64(dest.X / SIZE_CELL)))
	row := int(math.Round(float64(dest.Y / SIZE_CELL)))

	if col < 0 || col > 7 || row < 0 || row > 7 {
		return 0
	}
	if (row+col)%2 == 0 {
		return 0
	}
	return int32(32 - (row*4 + col/2))
}
