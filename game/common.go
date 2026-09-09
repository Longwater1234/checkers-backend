package game

import (
	"checkers-backend/player"
	"log"
)

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

	// check if player has 1 or 2 pieces left, if NOT then they definitely have valid moves available
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
