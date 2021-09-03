package concurrency

import (
	"fmt"
	"sync"
)

func JustOnce() {
	var once sync.Once
	var c = make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(doSomeThing)
			c <- false
		}()
	}

	for i := 0; i < 10; i++ {
		<-c
	}
}

func doSomeThing() {
	fmt.Println("da da da da ...")
}
