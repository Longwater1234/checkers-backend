package room

import (
	"checkers-backend/game"
	"checkers-backend/player"
	"crypto/rand"
	"log"
	"math/big"
)

const (
	numRows          = 8      // checker cells rows
	numCols          = 8      // checker cell columns
	upperLimit int16 = 0x7FFF //random ID max value (short_max)
)

// generateGameMap makes the hashmap of cell_index --> player piece
func generateGameMap(p1 *player.Player, p2 *player.Player) map[int32]*game.Piece {
	var gameMap = make(map[int32]*game.Piece)
	var counter int32 = 32
	var iterRed = 0
	var iterBlack = 0

	// create pieces objects, and position them on Board
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if (row+col)%2 != 0 {
				pos := game.Vec2{
					X: float32(col%numRows) * game.SIZE_CELL,
					Y: float32(row) * game.SIZE_CELL,
				}
				if row < 3 {
					//BLACK PIECES
					gameMap[counter] = &game.Piece{
						Id:         p2.Pieces[iterBlack],
						IsKing:     false,
						Pos:        pos,
						PieceColor: game.Piece_Black,
					}
					iterBlack++
				} else if row > 4 {
					//RED PIECES
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
	}
	return gameMap
}

// GeneratePlayerPieces for both player 1 and player 2
func generatePlayerPieces(p1 *player.Player, p2 *player.Player, gamOver chan bool) {
	for i := 0; i < len(p1.Pieces); i++ {
		val, err := rand.Int(rand.Reader, big.NewInt(int64(upperLimit)))
		if err != nil {
			gamOver <- true
			log.Panic("cannot generate random number", err)
		}
		p1.Pieces[i] = int32(val.Int64())
	}

	for i := 0; i < len(p2.Pieces); i++ {
		val, err := rand.Int(rand.Reader, big.NewInt(int64(upperLimit)))
		if err != nil {
			gamOver <- true
			log.Panic("cannot generate random number", err)
		}
		p2.Pieces[i] = int32(val.Int64())
	}
}