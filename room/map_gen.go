package room

import (
	"checkers-backend/game"
	"checkers-backend/player"
)

const numRows = 8

// generateGameMap makes the hashmap of cell index --> player piece
func generateGameMap(p1 *player.Player, p2 *player.Player) map[int32]*game.Piece {
	var gameMap = make(map[int32]*game.Piece)
	var counter int32 = 32
	var iterRed = 0
	var iterBlack = 0

	// create pieces objects, and position them on Board
	for row := 0; row < numRows; row++ {
		for col := 0; col < numRows; col++ {
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
