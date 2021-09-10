package task

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_wiat_group(t *testing.T) {
	var worker = func(id int, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Printf("worker %d staring\n", id)
		time.Sleep(time.Second)
		fmt.Printf("worker %d done\n", id)
	}

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
}
