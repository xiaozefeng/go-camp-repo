package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func Test_work(t *testing.T) {
	work("direct")

	go work("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done!")
}

func Test_channel(t *testing.T) {
	msgChan := make(chan string)
	go func() {
		msgChan <- "ping"
	}()

	msg := <-msgChan
	fmt.Println(msg)
}

func Test_buffered_channel(t *testing.T) {
	msgChan := make(chan string, 1)
	msgChan <- "buffered"
	msgChan <- "channel"

	fmt.Println(<-msgChan)
	fmt.Println(<-msgChan)
}

func dosomething(done chan bool) {
	fmt.Println("working")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func Test_channel_sync(t *testing.T) {
	doneChan := make(chan bool, 1)
	go dosomething(doneChan)
	<-doneChan
}

func Test_select(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Microsecond)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(1 * time.Microsecond)
		c1 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case v1 := <-c1:
			fmt.Println("received", v1)
		case v2 := <-c2:
			fmt.Println("received", v2)
		}
	}
}

func Test_timeout(t *testing.T) {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Printf("res: %v\n", res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Printf("res: %v\n", res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}

func Test_non_blocking_channel_ops(t *testing.T) {
	msgChan := make(chan string)
	sinalChan := make(chan bool)

	select {
	case msg := <-msgChan:
		fmt.Printf("received msg: %v\n", msg)
	default:
		fmt.Println("no message received")
	}

	var msg = "hi"
	select {
	case msgChan <- msg:
		fmt.Println("send message", msg)
	default:
		fmt.Println("not messsage sent")
	}

	select {
	case msg := <-msgChan:
		fmt.Printf("received msg: %v\n", msg)
	case sig := <-sinalChan:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

func Test_closing_channel(t *testing.T) {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			if j, ok := <-jobs; ok {
				fmt.Printf("received job: %v\n", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("sent job", i)
	}
	close(jobs)

	fmt.Println("send all jobs")

	<-done
}

func Test_range_over_channel(t *testing.T) {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for ele := range queue {
		fmt.Printf("ele: %v\n", ele)
	}
}
