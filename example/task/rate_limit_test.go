package task

import (
	"fmt"
	"testing"
	"time"
)

func Test_rate_limit_1(t *testing.T) {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Microsecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

}

func Test_rate_limit_2(t *testing.T) {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := make(chan time.Time)
	// for i := 0; i < 3; i++ {
	// 	limiter <- time.Now()
	// }

	go func() {
		for t := range time.Tick(200 * time.Microsecond) {
			limiter <- t
		}
	}()

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

}
