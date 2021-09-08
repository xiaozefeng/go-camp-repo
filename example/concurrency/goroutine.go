package concurrency

import "fmt"

func work(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}


