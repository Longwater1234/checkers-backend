package room

import (
	"checkers-backend/game"
	"checkers-backend/player"
)

// processCapturePiece made by Player `p` against `opponent`. Returns TRUE if all is OK. Else returns FALSE.
func processCapturePiece(request *game.BasePayload, gameMap map[int32]*game.Piece, p, opponent *player.Player) bool {
	capturePayload := request.GetCapturePayload()
	success := validateAndDoCapture(capturePayload, gameMap, p)
	if !success {
		p.SendMessage(&game.BasePayload{
			Notice: "Illegal move!",
			Inner: &game.BasePayload_ExitPayload{
				ExitPayload: &game.ExitPayload{
					FromTeam: game.TeamColor_TEAM_UNSPECIFIED,
				},
			},
		})
		opponent.SendMessage(&game.BasePayload{
			Notice: "Your opponent got kicked out!",
			Inner: &game.BasePayload_ExitPayload{
				ExitPayload: &game.ExitPayload{
					FromTeam: game.TeamColor_TEAM_UNSPECIFIED,
				},
			},
		})
		return false
	}
	preyPieceId := capturePayload.GetDetails().GetPreyPieceId()
	opponent.LosePiece(preyPieceId) // all is OK, Opponent loses 1 piece
	opponent.SendMessage(request)   // forward the original request to opponent
	return success
}

// validateAndDoCapture when player `p` attacks opponent's piece. Returns TRUE if successful
func validateAndDoCapture(captureReq *game.CapturePayload, gameMap map[int32]*game.Piece, p *player.Player) bool {
	if captureReq.GetDetails() == nil || captureReq.GetDestination() == nil {
		return false
	}
	hunterPieceId := captureReq.GetHunterPieceId()
	hunterSrcIdx := captureReq.GetDetails().GetHunterSrcCell()

	// check hunter params (must be owned by the player whose turn it is)
	hunterPiecePtr, exists := gameMap[hunterSrcIdx]
	if !exists || hunterPieceId != hunterPiecePtr.Id || !p.HasThisPiece(hunterPieceId) {
		return false
	}

	preyPieceId := captureReq.GetDetails().GetPreyPieceId()
	preyCellIdx := captureReq.GetDetails().GetPreyCellIdx()

	// check Prey params (must be an ENEMY piece)
	preyPiecePtr, exists := gameMap[preyCellIdx]
	if !exists || preyPieceId != preyPiecePtr.Id || preyPiecePtr.PieceColor == hunterPiecePtr.PieceColor {
		return false
	}

	// never trust client's cell_index: derive it from the (x,y) position
	destination := captureReq.GetDestination()
	destPos := game.Vec2{X: destination.GetX(), Y: destination.GetY()}
	destCellIdx := game.CellIndex(destPos)
	if destCellIdx == 0 {
		return false // off-board, or not a playable cell
	}
	if _, occupied := gameMap[destCellIdx]; occupied {
		return false
	}

	// Prey must sit exactly halfway between hunter and destination
	midPos := game.Vec2{
		X: (hunterPiecePtr.Pos.X + destPos.X) / 2,
		Y: (hunterPiecePtr.Pos.Y + destPos.Y) / 2,
	}
	if game.CellIndex(midPos) != preyCellIdx {
		return false
	}

	// perform the capture (also checks the 2-cell diagonal jump and direction)
	if !hunterPiecePtr.MoveCapture(destPos) {
		return false
	}
	destination.CellIndex = destCellIdx // so the forwarded request can't desync the opponent
	delete(gameMap, hunterSrcIdx)       // set hunter's old location empty!
	delete(gameMap, preyCellIdx)        // set Prey's old location empty!
	gameMap[destCellIdx] = hunterPiecePtr
	return true
}
