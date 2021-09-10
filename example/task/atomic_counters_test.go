package task

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func Test_atomic_counters(t *testing.T) {
	var ops uint64

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 1000; i++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("ops:", ops)
}
