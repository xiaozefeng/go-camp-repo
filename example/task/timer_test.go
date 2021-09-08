package task

import (
	"fmt"
	"testing"
	"time"
)

func Test_timer(t *testing.T) {
	t1 := time.NewTimer(2 * time.Second)

	<-t1.C
	fmt.Println("timer 1 fired")

	t2 := time.NewTimer(1 * time.Second)

	go func() {
		<-t2.C
		fmt.Println("timer 2 fired")
	}()

	// if stop2 := t2.Stop(); stop2 {
	// 	fmt.Println("timer 2 stopped")
	// }

	time.Sleep(2 * time.Second)

}
