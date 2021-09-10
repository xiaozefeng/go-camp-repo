package task

import (
	"fmt"
	"testing"
	"time"
)

func Test_timer(t *testing.T) {
	t1 := time.NewTimer(2 * time.Microsecond)

	<-t1.C
	fmt.Println("timer 1 fired")

	t2 := time.NewTimer(1 * time.Microsecond)

	go func() {
		<-t2.C
		fmt.Println("timer 2 fired")
	}()

	time.Sleep(2 * time.Microsecond)
	if stop2 := t2.Stop(); stop2 {
		fmt.Println("timer 2 stopped")
	}

	time.Sleep(2 * time.Microsecond)

}

func Test_ticker(t *testing.T) {
	ticker := time.NewTicker(500 * time.Microsecond)
	done := make(chan bool, 0)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("done!")
				return
			case t := <-ticker.C:
				fmt.Println("tick at ", t)
			}
		}
	}()

	time.Sleep(1600 * time.Microsecond)
	ticker.Stop()
	done <- true
	fmt.Println("ticker stopped")
}
