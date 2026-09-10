package game

import (
	"testing"
)

func TestGetCellIndex(t *testing.T) {
	// Verify all 32 playable cells match the expected layout
	// Row 0 (Y=0): cols 1, 3, 5, 7 -> cells 32, 31, 30, 29
	// Row 1 (Y=75): cols 0, 2, 4, 6 -> cells 28, 27, 26, 25
	// Row 2 (Y=150): cols 1, 3, 5, 7 -> cells 24, 23, 22, 21
	// Row 3 (Y=225): cols 0, 2, 4, 6 -> cells 20, 19, 18, 17
	// Row 4 (Y=300): cols 1, 3, 5, 7 -> cells 16, 15, 14, 13
	// Row 5 (Y=375): cols 0, 2, 4, 6 -> cells 12, 11, 10, 9
	// Row 6 (Y=450): cols 1, 3, 5, 7 -> cells 8, 7, 6, 5
	// Row 7 (Y=525): cols 0, 2, 4, 6 -> cells 4, 3, 2, 1
	expectedCells := map[int32]Vec2{
		32: {X: 1 * SIZE_CELL, Y: 0 * SIZE_CELL},
		31: {X: 3 * SIZE_CELL, Y: 0 * SIZE_CELL},
		30: {X: 5 * SIZE_CELL, Y: 0 * SIZE_CELL},
		29: {X: 7 * SIZE_CELL, Y: 0 * SIZE_CELL},

		28: {X: 0 * SIZE_CELL, Y: 1 * SIZE_CELL},
		27: {X: 2 * SIZE_CELL, Y: 1 * SIZE_CELL},
		26: {X: 4 * SIZE_CELL, Y: 1 * SIZE_CELL},
		25: {X: 6 * SIZE_CELL, Y: 1 * SIZE_CELL},

		24: {X: 1 * SIZE_CELL, Y: 2 * SIZE_CELL},
		23: {X: 3 * SIZE_CELL, Y: 2 * SIZE_CELL},
		22: {X: 5 * SIZE_CELL, Y: 2 * SIZE_CELL},
		21: {X: 7 * SIZE_CELL, Y: 2 * SIZE_CELL},

		20: {X: 0 * SIZE_CELL, Y: 3 * SIZE_CELL},
		19: {X: 2 * SIZE_CELL, Y: 3 * SIZE_CELL},
		18: {X: 4 * SIZE_CELL, Y: 3 * SIZE_CELL},
		17: {X: 6 * SIZE_CELL, Y: 3 * SIZE_CELL},

		16: {X: 1 * SIZE_CELL, Y: 4 * SIZE_CELL},
		15: {X: 3 * SIZE_CELL, Y: 4 * SIZE_CELL},
		14: {X: 5 * SIZE_CELL, Y: 4 * SIZE_CELL},
		13: {X: 7 * SIZE_CELL, Y: 4 * SIZE_CELL},

		12: {X: 0 * SIZE_CELL, Y: 5 * SIZE_CELL},
		11: {X: 2 * SIZE_CELL, Y: 5 * SIZE_CELL},
		10: {X: 4 * SIZE_CELL, Y: 5 * SIZE_CELL},
		9:  {X: 6 * SIZE_CELL, Y: 5 * SIZE_CELL},

		8: {X: 1 * SIZE_CELL, Y: 6 * SIZE_CELL},
		7: {X: 3 * SIZE_CELL, Y: 6 * SIZE_CELL},
		6: {X: 5 * SIZE_CELL, Y: 6 * SIZE_CELL},
		5: {X: 7 * SIZE_CELL, Y: 6 * SIZE_CELL},

		4: {X: 0 * SIZE_CELL, Y: 7 * SIZE_CELL},
		3: {X: 2 * SIZE_CELL, Y: 7 * SIZE_CELL},
		2: {X: 4 * SIZE_CELL, Y: 7 * SIZE_CELL},
		1: {X: 6 * SIZE_CELL, Y: 7 * SIZE_CELL},
	}

	for expectedIdx, pos := range expectedCells {
		actualIdx := getCellIndex(pos)
		if actualIdx != expectedIdx {
			t.Errorf("getCellIndex(%+v) = %d; want %d", pos, actualIdx, expectedIdx)
		}
	}

	// Test light (unplayable) cells return 0
	lightCells := []Vec2{
		{X: 0, Y: 0},
		{X: 2 * SIZE_CELL, Y: 0},
		{X: 1 * SIZE_CELL, Y: 1 * SIZE_CELL},
		{X: 7 * SIZE_CELL, Y: 7 * SIZE_CELL},
	}
	for _, pos := range lightCells {
		if idx := getCellIndex(pos); idx != 0 {
			t.Errorf("getCellIndex light cell (%+v) = %d; want 0", pos, idx)
		}
	}

	// Test out of bounds positions return 0
	oobPositions := []Vec2{
		{X: -SIZE_CELL, Y: 0},
		{X: 0, Y: -SIZE_CELL},
		{X: 8 * SIZE_CELL, Y: 0},
		{X: 0, Y: 8 * SIZE_CELL},
		{X: -SIZE_CELL, Y: -SIZE_CELL},
		{X: 8 * SIZE_CELL, Y: 8 * SIZE_CELL},
	}
	for _, pos := range oobPositions {
		if idx := getCellIndex(pos); idx != 0 {
			t.Errorf("getCellIndex OOB position (%+v) = %d; want 0", pos, idx)
		}
	}
}

func TestCanMoveLegally_NilCases(t *testing.T) {
	var p *Piece
	gameMap := make(map[int32]*Piece)
	if p.canMoveLegally(gameMap) {
		t.Errorf("canMoveLegally on nil piece should return false")
	}

	p = &Piece{Id: 1, Pos: Vec2{X: 75, Y: 75}, PieceColor: Piece_Red}
	if p.canMoveLegally(nil) {
		t.Errorf("canMoveLegally with nil gameMap should return false")
	}
}

func TestCanMoveLegally_SimpleMoves(t *testing.T) {
	// Red piece at cell 10 (row 5, col 4: X=300, Y=375)
	// Open board -> Red moves up towards Y=0
	redPiece := &Piece{
		Id:         100,
		Pos:        Vec2{X: 4 * SIZE_CELL, Y: 5 * SIZE_CELL},
		PieceColor: Piece_Red,
		IsKing:     false,
	}
	gameMap := map[int32]*Piece{
		10: redPiece,
	}

	if !redPiece.canMoveLegally(gameMap) {
		t.Errorf("Red piece on open board should be able to move legally")
	}

	// Black piece at cell 23 (row 2, col 3: X=225, Y=150)
	// Open board -> Black moves down towards Y=7*SIZE_CELL
	blackPiece := &Piece{
		Id:         200,
		Pos:        Vec2{X: 3 * SIZE_CELL, Y: 2 * SIZE_CELL},
		PieceColor: Piece_Black,
		IsKing:     false,
	}
	gameMap[23] = blackPiece
	if !blackPiece.canMoveLegally(gameMap) {
		t.Errorf("Black piece on open board should be able to move legally")
	}
}

func TestCanMoveLegally_BlockedByFriendly(t *testing.T) {
	// Red piece at cell 10 (row 5, col 4: X=300, Y=375)
	// Block up-left: row 4, col 3 (cell 15)
	// Block up-right: row 4, col 5 (cell 14)
	redPiece := &Piece{
		Id:         100,
		Pos:        Vec2{X: 4 * SIZE_CELL, Y: 5 * SIZE_CELL},
		PieceColor: Piece_Red,
		IsKing:     false,
	}
	blocker1 := &Piece{
		Id:         101,
		Pos:        Vec2{X: 3 * SIZE_CELL, Y: 4 * SIZE_CELL},
		PieceColor: Piece_Red,
	}
	blocker2 := &Piece{
		Id:         102,
		Pos:        Vec2{X: 5 * SIZE_CELL, Y: 4 * SIZE_CELL},
		PieceColor: Piece_Red,
	}

	gameMap := map[int32]*Piece{
		10: redPiece,
		15: blocker1,
		14: blocker2,
	}

	if redPiece.canMoveLegally(gameMap) {
		t.Errorf("Red piece completely blocked by friendly pieces should not be able to move")
	}
}

func TestCanMoveLegally_CaptureEnemy(t *testing.T) {
	// Red piece at cell 10 (row 5, col 4: X=300, Y=375)
	// Enemy Black piece at cell 15 (row 4, col 3: X=225, Y=300)
	// Landing cell for jump: row 3, col 2 (cell 19: X=150, Y=225) is empty
	// Up-right is blocked by friendly Red at cell 14 (row 4, col 5)
	redPiece := &Piece{
		Id:         100,
		Pos:        Vec2{X: 4 * SIZE_CELL, Y: 5 * SIZE_CELL},
		PieceColor: Piece_Red,
		IsKing:     false,
	}
	enemyBlack := &Piece{
		Id:         201,
		Pos:        Vec2{X: 3 * SIZE_CELL, Y: 4 * SIZE_CELL},
		PieceColor: Piece_Black,
	}
	friendlyBlocker := &Piece{
		Id:         102,
		Pos:        Vec2{X: 5 * SIZE_CELL, Y: 4 * SIZE_CELL},
		PieceColor: Piece_Red,
	}

	gameMap := map[int32]*Piece{
		10: redPiece,
		15: enemyBlack,
		14: friendlyBlocker,
	}

	if !redPiece.canMoveLegally(gameMap) {
		t.Errorf("Red piece should be able to capture enemy when landing cell is open")
	}

	// Now block the landing square cell 19
	landingBlocker := &Piece{
		Id:         202,
		Pos:        Vec2{X: 2 * SIZE_CELL, Y: 3 * SIZE_CELL},
		PieceColor: Piece_Black,
	}
	gameMap[19] = landingBlocker

	if redPiece.canMoveLegally(gameMap) {
		t.Errorf("Red piece should not be able to capture when landing cell is occupied")
	}
}

func TestCanMoveLegally_CaptureOffBoard(t *testing.T) {
	// Enemy is at row 0 (edge). Red piece is at row 1.
	// Red cannot jump over enemy because landing square would be at row -1 (off board).
	// Red piece at cell 27 (row 1, col 2: X=150, Y=75)
	// Enemy Black piece at cell 32 (row 0, col 1: X=75, Y=0)
	// Friendly Red at cell 31 (row 0, col 3: X=225, Y=0)
	redPiece := &Piece{
		Id:         100,
		Pos:        Vec2{X: 2 * SIZE_CELL, Y: 1 * SIZE_CELL},
		PieceColor: Piece_Red,
		IsKing:     false,
	}
	enemyBlack := &Piece{
		Id:         201,
		Pos:        Vec2{X: 1 * SIZE_CELL, Y: 0 * SIZE_CELL},
		PieceColor: Piece_Black,
	}
	friendlyRed := &Piece{
		Id:         101,
		Pos:        Vec2{X: 3 * SIZE_CELL, Y: 0 * SIZE_CELL},
		PieceColor: Piece_Red,
	}

	gameMap := map[int32]*Piece{
		27: redPiece,
		32: enemyBlack,
		31: friendlyRed,
	}

	if redPiece.canMoveLegally(gameMap) {
		t.Errorf("Red piece cannot jump enemy on edge if landing is off-board")
	}
}

func TestCanMoveLegally_KingMovesAndCaptures(t *testing.T) {
	// Red King at cell 19 (row 3, col 2: X=150, Y=225)
	// Block forward directions (row 2, col 1: cell 24; row 2, col 3: cell 23)
	// Open backward direction (row 4, col 1: cell 16)
	king := &Piece{
		Id:         100,
		Pos:        Vec2{X: 2 * SIZE_CELL, Y: 3 * SIZE_CELL},
		PieceColor: Piece_Red,
		IsKing:     true,
	}
	blocker1 := &Piece{
		Id:         101,
		Pos:        Vec2{X: 1 * SIZE_CELL, Y: 2 * SIZE_CELL},
		PieceColor: Piece_Red,
	}
	blocker2 := &Piece{
		Id:         102,
		Pos:        Vec2{X: 3 * SIZE_CELL, Y: 2 * SIZE_CELL},
		PieceColor: Piece_Red,
	}

	gameMap := map[int32]*Piece{
		19: king,
		24: blocker1,
		23: blocker2,
	}

	// King should be able to move backwards into cell 16 or cell 15
	if !king.canMoveLegally(gameMap) {
		t.Errorf("King should be able to move backwards")
	}

	// Block cell 16 with friendly piece
	blocker3 := &Piece{
		Id:         103,
		Pos:        Vec2{X: 1 * SIZE_CELL, Y: 4 * SIZE_CELL},
		PieceColor: Piece_Red,
	}
	gameMap[16] = blocker3

	// Place enemy at cell 15 (row 4, col 3: X=225, Y=300)
	enemy := &Piece{
		Id:         201,
		Pos:        Vec2{X: 3 * SIZE_CELL, Y: 4 * SIZE_CELL},
		PieceColor: Piece_Black,
	}
	gameMap[15] = enemy

	// King can capture enemy backwards into row 5, col 4 (cell 10: X=300, Y=375)
	if !king.canMoveLegally(gameMap) {
		t.Errorf("King should be able to capture backwards into open square")
	}

	// Block cell 10 so jump is blocked
	blocker4 := &Piece{
		Id:         104,
		Pos:        Vec2{X: 4 * SIZE_CELL, Y: 5 * SIZE_CELL},
		PieceColor: Piece_Red,
	}
	gameMap[10] = blocker4

	// Now all forward moves blocked, backward left blocked, backward right enemy jump blocked
	if king.canMoveLegally(gameMap) {
		t.Errorf("King with all moves and captures blocked should return false")
	}
}
