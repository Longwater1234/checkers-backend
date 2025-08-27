package room

import (
	"checkers-backend/game"
	"checkers-backend/player"
	"context"
	"log"
	"time"
)

const serverVersion = "1.0.13"

// Keep Listening for new players joining lobby. Then forward a pair to new match room
func ListenForJoins(lobby <-chan *player.Player) {
	for {
		//welcome 1st player
		p1 := <-lobby
		var msgOne = &game.BasePayload{
			Notice: "Connected. You are Team RED. Waiting for opponent...",
			Inner: &game.BasePayload_Welcome{
				Welcome: &game.WelcomePayload{
					MyTeam:        game.TeamColor_TEAM_RED,
					ServerVersion: serverVersion,
				},
			},
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		go p1.StartHeartbeat(ctx)
		p1.SendMessage(msgOne)

		//waiting for 2nd player to join (TIMEOUT at 30 seconds)
		select {
		case p2 := <-lobby:
			cancel()
			// welcome 2nd player
			var msgTwo = &game.BasePayload{
				Notice: "Connected. You are Team BLACK. Match is starting!",
				Inner: &game.BasePayload_Welcome{
					Welcome: &game.WelcomePayload{
						MyTeam: game.TeamColor_TEAM_BLACK,
					},
				},
			}
			p2.SendMessage(msgTwo)

			// start the match in new goroutine
			go func(p1, p2 *player.Player) {
				// sleep REQUIRED for [p2] client to process prev message
				time.Sleep(200 * time.Millisecond)
				gameOver := make(chan bool, 1)
				StartMatch(p1, p2, gameOver)
				<-gameOver // block until match ends
				log.Println("🔴 GAME OVER!")
				p1.Dead <- true
				p2.Dead <- true
			}(p1, p2)

		case <-ctx.Done():
			// timeout reached. No other player joined! Goodbye p1!
			p1.SendMessage(&game.BasePayload{
				Notice: "No other players at this moment. Try again later!",
				Inner: &game.BasePayload_ExitPayload{
					ExitPayload: &game.ExitPayload{
						FromTeam: game.TeamColor_TEAM_UNSPECIFIED,
					},
				},
			})
			p1.Dead <- true

		case <-p1.Quit:
			//[p1] has quit before match started
			cancel()
			p1.Dead <- true
		}
		// goto TOP... wait for another pair to join.
	}

}
