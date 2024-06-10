package room

import (
	"checkers-backend/game"
	"checkers-backend/player"
	"crypto/rand"
	"log"
	"math/big"
)

const upperLimit int16 = 0x7FFF //random ID max value (short_max)

func StartMatch(p1 *player.Player, p2 *player.Player, gamOver chan bool) {
	log.Println("🟢 Match has begun!")

	//make random pieceId's for player 1
	for i := 0; i < len(p1.Pieces); i++ {
		val, err := rand.Int(rand.Reader, big.NewInt(int64(upperLimit)))
		if err != nil {
			gamOver <- true
			log.Panic("cannot generate random number", err)
		}
		p1.Pieces[i] = int32(val.Int64())
	}

	//make pieces for player 2
	for i := 0; i < len(p2.Pieces); i++ {
		val, err := rand.Int(rand.Reader, big.NewInt(int64(upperLimit)))
		if err != nil {
			gamOver <- true
			log.Panic("cannot generate random number", err)
		}
		p2.Pieces[i] = int32(val.Int64())
	}

	p1.SendMessage(&game.BasePayload{
		Notice: "Opponent joined. Make your first move!",
		Inner: &game.BasePayload_Start{
			Start: &game.StartPayload{
				PiecesRed:   p1.Pieces,
				PiecesBlack: p2.Pieces,
			},
		},
	})

	p2.SendMessage(&game.BasePayload{
		Notice: "Match has begun. Waiting for RED to move!",
		Inner: &game.BasePayload_Start{
			Start: &game.StartPayload{
				PiecesRed:   p1.Pieces,
				PiecesBlack: p2.Pieces,
			},
		},
	})

	var isPlayerRedTurn = true

	log.Println("playerRedturn", isPlayerRedTurn)

	//CREATE gameMap for this match
	//var gameMap = make(map[int32]int32)
	//START GAME MAIN LOOP

}
