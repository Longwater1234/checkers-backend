package room

import (
	"checkers-backend/game"
	"checkers-backend/player"
	"crypto/rand"
	"math"
	"math/big"
)

const (
	numRows          = 8             // checker rows
	numCols          = 8             // checker columns
	upperLimit int16 = math.MaxInt16 // piece ID max value (short_max)
)

// generateGameMap creates the hashmap of cell_index --> Piece. This should be called AFTER `generatePieces`
func generateGameMap(p1 *player.Player, p2 *player.Player) map[int32]*game.Piece {
	var gameMap = make(map[int32]*game.Piece, 24) // cell_idx -> piece
	var counter int32 = 32                        // total playable checker cells
	var iterRed = 0                               // red pieces iterator
	var iterBlack = 0                             // black pieces iterator

	// create pieces, and position them on checkerboard (from top -> down)
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if (row+col)%2 == 0 {
				continue // skip even-number cells
			}
			pos := game.Vec2{
				X: float32(col%numRows) * game.SIZE_CELL,
				Y: float32(row) * game.SIZE_CELL,
			}
			if row < 3 && iterBlack != len(p2.Pieces) {
				// BLACK PIECES
				gameMap[counter] = &game.Piece{
					Id:         p2.Pieces[iterBlack],
					IsKing:     false,
					Pos:        pos,
					PieceColor: game.Piece_Black,
				}
				iterBlack++
			} else if row > 4 && iterRed != len(p1.Pieces) {
				// RED PIECES
				gameMap[counter] = &game.Piece{
					Id:         p1.Pieces[iterRed],
					IsKing:     false,
					Pos:        pos,
					PieceColor: game.Piece_Red,
				}
				iterRed++
			}
			counter--
		}
	}
	return gameMap
}

// generatePieces using secure RNG for the two players.
func generatePieces(p1 *player.Player, p2 *player.Player) error {
	bigMax := big.NewInt(int64(upperLimit))
	for i := range p1.Pieces {
		val, err := rand.Int(rand.Reader, bigMax)
		if err != nil {
			return err
		}
		p1.Pieces[i] = int32(val.Int64())
	}

	for i := range p2.Pieces {
		val, err := rand.Int(rand.Reader, bigMax)
		if err != nil {
			return err
		}
		p2.Pieces[i] = int32(val.Int64())
	}
	return nil
}
