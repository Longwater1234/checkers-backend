package game

import "math"

const SIZE_CELL float32 = 75.0

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
	PieceColor PieceType // either red or black
}

// When a piece is moved to given destPos. Returns TRUE if move is valid, else FALSE
func (pc *Piece) MoveSimple(destPos *Vec2) bool {
	var deltaX = float64(destPos.X - pc.Pos.X)
	var deltaY = float64(destPos.Y - pc.Pos.Y)

	if math.Abs(deltaX) > float64(SIZE_CELL) || math.Abs(deltaY) > float64(SIZE_CELL) {
		return false
	}
	if pc.PieceColor == Piece_Red && deltaY > 0.0 && !pc.IsKing {
		return false
	}
	if pc.PieceColor == Piece_Black && deltaY < 0.0 && !pc.IsKing {
		return false
	}

	pc.Pos.X = destPos.X
	pc.Pos.Y = destPos.Y
	if (pc.PieceColor == Piece_Red && destPos.Y == 0) ||
		(pc.PieceColor == Piece_Black && destPos.Y == 7*SIZE_CELL) {
		pc.IsKing = true
	}
	return true
}

// When capturing opponent, Move this piece by 2 cells diagonally to the given `destPos`. Returns TRUE
// if success, else FALSE
func (pc *Piece) MoveCapture(destPos *Vec2) bool {
	var deltaX = float64(destPos.X - pc.Pos.X)
	var deltaY = float64(destPos.Y - pc.Pos.Y)

	if math.Abs(deltaX) != float64(2*SIZE_CELL) || math.Abs(deltaY) != float64(2*SIZE_CELL) {
		return false
	}
	if pc.PieceColor == Piece_Red && deltaY > 0.0 && !pc.IsKing {
		return false
	}
	if pc.PieceColor == Piece_Black && deltaY < 0.0 && !pc.IsKing {
		return false
	}

	pc.Pos.X = destPos.X
	pc.Pos.Y = destPos.Y
	if (pc.PieceColor == Piece_Red && destPos.Y == 0) ||
		(pc.PieceColor == Piece_Black && destPos.Y == 7*SIZE_CELL) {
		pc.IsKing = true
	}
	return true
}
