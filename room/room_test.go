package room

import (
	"checkers-backend/game"
	"checkers-backend/player"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/websocket"
	"google.golang.org/protobuf/proto"
)

// setupMatchTest spins up an in-memory websocket server, connects two players (RED and BLACK),
// and launches StartMatch in a background goroutine.
func setupMatchTest(t *testing.T) (redClient, blackClient *websocket.Conn, startPayload *game.StartPayload, gameOver <-chan bool) {
	t.Helper()

	serverConns := make(chan *websocket.Conn, 2)
	done := make(chan struct{})
	t.Cleanup(func() { close(done) })

	srv := httptest.NewServer(websocket.Handler(func(ws *websocket.Conn) {
		serverConns <- ws
		<-done
	}))
	t.Cleanup(srv.Close)

	wsURL := "ws" + strings.TrimPrefix(srv.URL, "http")

	c1, err := websocket.Dial(wsURL, "", "http://localhost")
	if err != nil {
		t.Fatalf("failed to dial red client: %v", err)
	}
	t.Cleanup(func() { _ = c1.Close() })
	s1 := <-serverConns

	c2, err := websocket.Dial(wsURL, "", "http://localhost")
	if err != nil {
		t.Fatalf("failed to dial black client: %v", err)
	}
	t.Cleanup(func() { _ = c2.Close() })
	s2 := <-serverConns

	p1 := &player.Player{
		Conn:   s1,
		Name:   game.TeamColor_TEAM_RED.String(),
		Pieces: make([]int32, 12),
		Dead:   make(chan bool, 1),
	}
	p2 := &player.Player{
		Conn:   s2,
		Name:   game.TeamColor_TEAM_BLACK.String(),
		Pieces: make([]int32, 12),
		Dead:   make(chan bool, 1),
	}

	over := make(chan bool, 1)
	go StartMatch(p1, p2, over)

	// Drain initial StartPayload sent to both players by notifyMatchStart
	msg1 := readPayload(t, c1)
	msg2 := readPayload(t, c2)
	if msg1.GetStart() == nil || msg2.GetStart() == nil {
		t.Fatalf("expected StartPayload for both players, got %v and %v", msg1, msg2)
	}

	return c1, c2, msg1.GetStart(), over
}

func sendPayload(t *testing.T, conn *websocket.Conn, payload *game.BasePayload) {
	t.Helper()
	raw, err := proto.Marshal(payload)
	if err != nil {
		t.Fatalf("failed to marshal payload: %v", err)
	}
	if err := websocket.Message.Send(conn, raw); err != nil {
		t.Fatalf("failed to send websocket message: %v", err)
	}
}

func readPayload(t *testing.T, conn *websocket.Conn) *game.BasePayload {
	t.Helper()
	_ = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	var raw []byte
	if err := websocket.Message.Receive(conn, &raw); err != nil {
		t.Fatalf("failed to receive websocket message: %v", err)
	}
	var payload game.BasePayload
	if err := proto.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("failed to unmarshal payload: %v", err)
	}
	return &payload
}

func TestStartMatch_CheatMoveOpponentPiece(t *testing.T) {
	redClient, blackClient, start, gameOver := setupMatchTest(t)

	// On initial board layout, Black's piece at index 8 sits on cell 24 (col 1, row 2).
	// Red (whose turn it is first) attempts to move Black's piece to cell 20 (col 0, row 3).
	blackPieceOnCell24 := start.GetPiecesBlack()[8]

	cheatMove := &game.BasePayload{
		Inner: &game.BasePayload_MovePayload{
			MovePayload: &game.MovePayload{
				PieceId:    blackPieceOnCell24,
				SourceCell: 24,
				Destination: &game.MovePayload_Detination{
					CellIndex: 20,
					X:         0 * game.SIZE_CELL,
					Y:         3 * game.SIZE_CELL,
				},
			},
		},
	}
	sendPayload(t, redClient, cheatMove)

	redResp := readPayload(t, redClient)
	if redResp.GetNotice() != "Illegal move!" || redResp.GetExitPayload() == nil {
		t.Errorf("expected Red to get kicked with 'Illegal move!' ExitPayload, got: %+v", redResp)
	}

	blackResp := readPayload(t, blackClient)
	if blackResp.GetNotice() != "Your opponent got kicked out!" || blackResp.GetExitPayload() == nil {
		t.Errorf("expected Black to receive opponent kicked ExitPayload, got: %+v", blackResp)
	}

	select {
	case <-gameOver:
	case <-time.After(2 * time.Second):
		t.Fatal("expected match to end via gameOver channel after illegal move")
	}
}

func TestStartMatch_CheatCaptureNonAdjacentEnemy(t *testing.T) {
	redClient, blackClient, start, gameOver := setupMatchTest(t)

	// Red's piece at index 1 sits on cell 11 (col 2, row 5: X=150, Y=375).
	// A 2-cell diagonal jump from cell 11 lands on cell 20 (col 0, row 3: X=0, Y=225),
	// jumping over cell 16 (col 1, row 4).
	// Red cheats by claiming to capture Black's non-adjacent piece on cell 24 (col 1, row 2).
	redPieceOnCell11 := start.GetPiecesRed()[1]
	blackPieceOnCell24 := start.GetPiecesBlack()[8]

	cheatCapture := &game.BasePayload{
		Inner: &game.BasePayload_CapturePayload{
			CapturePayload: &game.CapturePayload{
				HunterPieceId: redPieceOnCell11,
				Details: &game.CapturePayload_TargetDetails{
					HunterSrcCell: 11,
					PreyPieceId:   blackPieceOnCell24,
					PreyCellIdx:   24,
				},
				Destination: &game.CapturePayload_HunterDestination{
					CellIndex: 20,
					X:         0 * game.SIZE_CELL,
					Y:         3 * game.SIZE_CELL,
				},
			},
		},
	}
	sendPayload(t, redClient, cheatCapture)

	redResp := readPayload(t, redClient)
	if redResp.GetNotice() != "Illegal move!" || redResp.GetExitPayload() == nil {
		t.Errorf("expected Red to get kicked with 'Illegal move!' ExitPayload, got: %+v", redResp)
	}

	blackResp := readPayload(t, blackClient)
	if blackResp.GetNotice() != "Your opponent got kicked out!" || blackResp.GetExitPayload() == nil {
		t.Errorf("expected Black to receive opponent kicked ExitPayload, got: %+v", blackResp)
	}

	select {
	case <-gameOver:
	case <-time.After(2 * time.Second):
		t.Fatal("expected match to end via gameOver channel after illegal capture")
	}
}
