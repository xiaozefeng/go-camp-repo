package channel_test

import (
	"fmt"
	"sync"
	"testing"
)

func Test_unbuffered_channel(t *testing.T) {
	var ch = make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		ch <- "content"
	}()

	go func() {
		defer wg.Done()
		fmt.Println("msg:", <-ch)
	}()

	wg.Wait()
}

func Test_test_buffered_channel(t *testing.T){ 
	var ch = make(chan string, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		defer wg.Done()
		ch<- "hello"
		ch<- "world"
	}()

	go func(){
		defer wg.Done()
		fmt.Println("msg:", <- ch)
		fmt.Println("msg:", <- ch)
	}()

	wg.Wait()

}