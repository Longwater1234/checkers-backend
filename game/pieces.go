package game

import (
	"checkers-backend/player"
	"log"
	"math"
)

const SIZE_CELL float32 = 75.0 //length of single square cell

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

// When a piece is moved diagonally to given `destPos`. Returns TRUE if move is successful, else FALSE
func (p *Piece) MoveSimple(destPos *Vec2) bool {
	var deltaX = float64(destPos.X - p.Pos.X)
	var deltaY = float64(destPos.Y - p.Pos.Y)

	if math.Abs(deltaX) > float64(SIZE_CELL) || math.Abs(deltaY) > float64(SIZE_CELL) {
		return false
	}
	if p.PieceColor == Piece_Red && deltaY > 0 && !p.IsKing {
		return false
	}
	if p.PieceColor == Piece_Black && deltaY < 0 && !p.IsKing {
		return false
	}

	p.Pos.X = destPos.X
	p.Pos.Y = destPos.Y
	if (p.PieceColor == Piece_Red && destPos.Y == 0) ||
		(p.PieceColor == Piece_Black && destPos.Y == 7*SIZE_CELL) {
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
	if p.PieceColor == Piece_Red && deltaY > 0 && !p.IsKing {
		return false
	}
	if p.PieceColor == Piece_Black && deltaY < 0 && !p.IsKing {
		return false
	}

	p.Pos.X = destPos.X
	p.Pos.Y = destPos.Y
	if (p.PieceColor == Piece_Red && destPos.Y == 0) ||
		(p.PieceColor == Piece_Black && destPos.Y == 7*SIZE_CELL) {
		p.IsKing = true
	}
	return true
}

// IsEvenCellRow determines wheter given cell_index is on even Row
func IsEvenCellRow(cellIdx int32) bool {
	rowNumber := 9 - (cellIdx-1)/4
	return (rowNumber%2 == 0)
}

// AwayFromEdge returns TRUE if given position is NOT on any edge of board
func AwayFromEdge(pos *Vec2) bool {
	return pos.X > 0 && pos.X < 7*SIZE_CELL && pos.Y > 0 && pos.Y < 7*SIZE_CELL
}

// HasWinner determines if `p` has won the match, and notifies both players if TRUE.
func HasWinner(p *player.Player, opponent *player.Player) bool {
	if len(opponent.Pieces) == 0 {
		//`opponent` has lost, `p` has won! game over
		p.SendMessage(&BasePayload{
			Notice: "Congrats! You won! GAME OVER",
			Inner: &BasePayload_WinlosePayload{
				WinlosePayload: &WinLosePayload{
					Winner: TeamColor_TEAM_UNSPECIFIED, //TODO fix me
				},
			},
		})
		opponent.SendMessage(&BasePayload{
			Notice: "Sorry! You lost! GAME OVER",
			Inner: &BasePayload_WinlosePayload{
				WinlosePayload: &WinLosePayload{
					Winner: TeamColor_TEAM_UNSPECIFIED, //TODO fix me
				},
			},
		})
		log.Println("🏆 We got a winner!", p.Name, " has won!")
		return true
	}
	return false
}
