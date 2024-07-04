## TODO

- [x] Add server version inside Welcome message from SERVER-->CLIENT
- [ ] Client needs to check if server version is not NEWER than its own SPECIFIED
- [ ] Extra endpoint for live server metrics
- [ ] Add timeout after every move. 20 seconds or 10 seconds


## timer example 

```go
func runMatch() {
	// Initialize timers for player 1 and player 2
	player1Timer := time.NewTimer(20 * time.Second)
	player2Timer := time.NewTimer(20 * time.Second)

	// START GAME MAIN LOOP
	for {
		select {
		case <-player1Timer.C:
			// Player 1 timeout
			log.Println("Player 1 timeout. Ending match.")
			// Notify both players
			// End the match logic here

		case <-player2Timer.C:
			// Player 2 timeout
			log.Println("Player 2 timeout. Ending match.")
			// Notify both players
			// End the match logic here

		default:
			if isPlayerRedTurn {
				// IT'S PLAYER 1 (red's) TURN
				var rawBytes []byte
				if err := websocket.Message.Receive(p1.Conn, &rawBytes); err != nil {
					log.Println(p1.Name, "disconnected. Cause:", err)
					p2.SendMessage("p1 left")
				}
				// validate & process rawBytes ... not shown
				// ...
				p2.SendMessage(rawBytes) // forward to player 2
				// Reset player 1 timer
				player1Timer.Reset(20 * time.Second)
			} else {
				// IT'S PLAYER 2 (black) TURN
				// .. same as above, but now receive from P2.
				// validate & process rawBytes ... not shown
				// ...
				p1.SendMessage(rawBytes) // forward to player 1
				// Reset player 2 timer
				player2Timer.Reset(20 * time.Second)
			}
		}
	}
}


```