package player

import (
	"sync"
	"testing"
	"time"
)

func TestKillClosesDoneOnce(t *testing.T) {
	p := New(nil)
	done := p.Done()

	var wg sync.WaitGroup
	for range 10 {
		wg.Go(func() {
			p.Kill()
		})
	}
	wg.Wait()

	select {
	case _, ok := <-done:
		if ok {
			t.Fatal("expected Done channel to be closed")
		}
	case <-time.After(time.Second):
		t.Fatal("expected Kill to close Done channel")
	}
}
