package game

import (
	"math"
)

const SIZE_CELL float32 = 75.0 //length of single square cell

type Vec2 struct {
	X float32 // x position
	Y float32 // y position
}

// type PieceType int16

// const (
// 	Piece_Red PieceType = iota + 1
// 	Piece_Black
// )

type Piece struct {
	Id         int32     // unique piece id
	IsKing     bool      // whether this piece is King
	Pos        Vec2      // current piece position
	PieceColor TeamColor // either red or black
}

// When a piece is moved to given destPos. Returns TRUE if move is valid, else FALSE
func (p *Piece) MoveSimple(destPos *Vec2) bool {
	var deltaX = float64(destPos.X - p.Pos.X)
	var deltaY = float64(destPos.Y - p.Pos.Y)

	if math.Abs(deltaX) > float64(SIZE_CELL) || math.Abs(deltaY) > float64(SIZE_CELL) {
		return false
	}
	if p.PieceColor == TeamColor_TEAM_RED && deltaY > 0 && !p.IsKing {
		return false
	}
	if p.PieceColor == TeamColor_TEAM_BLACK && deltaY < 0 && !p.IsKing {
		return false
	}

	p.Pos.X = destPos.X
	p.Pos.Y = destPos.Y
	if (p.PieceColor == TeamColor_TEAM_RED && destPos.Y == 0) ||
		(p.PieceColor == TeamColor_TEAM_BLACK && destPos.Y == 7*SIZE_CELL) {
		p.IsKing = true
	}
	return true
}

// When capturing opponent, Move this piece by 2 cells diagonally to the given `destPos`. Returns TRUE
// if success, else FALSE
func (p *Piece) MoveCapture(destPos *Vec2) bool {
	var deltaX = float64(destPos.X - p.Pos.X)
	var deltaY = float64(destPos.Y - p.Pos.Y)

	if math.Abs(deltaX) != float64(2*SIZE_CELL) || math.Abs(deltaY) != float64(2*SIZE_CELL) {
		return false
	}
	if p.PieceColor == TeamColor_TEAM_RED && deltaY > 0 && !p.IsKing {
		return false
	}
	if p.PieceColor == TeamColor_TEAM_BLACK && deltaY < 0 && !p.IsKing {
		return false
	}

	p.Pos.X = destPos.X
	p.Pos.Y = destPos.Y
	if (p.PieceColor == TeamColor_TEAM_RED && destPos.Y == 0) ||
		(p.PieceColor == TeamColor_TEAM_BLACK && destPos.Y == 7*SIZE_CELL) {
		p.IsKing = true
	}
	return true
}
