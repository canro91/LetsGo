package sync

import (
	"sync"
	"testing"
)

func TestSync(t *testing.T){
	t.Run("Increment", func(t *testing.T){
		counter := Counter{}

		counter.Inc()
		counter.Inc()
		counter.Inc()

		if counter.Value() != 3 {
			t.Errorf("Expected 3, but was %d", counter.Value())
		}
	})

	t.Run("Concurrent", func(t *testing.T){
		iterations := 1000

		var w sync.WaitGroup
		w.Add(iterations)

		counter := Counter{}
		for i := 0; i < iterations; i++ {
			go func(w *sync.WaitGroup){
				counter.Inc()
				w.Done()
			}(&w)
		}
		w.Wait()

		if counter.Value() != 1000 {
			t.Errorf("Expected 1000, but was %d", counter.Value())
		}
	})
}