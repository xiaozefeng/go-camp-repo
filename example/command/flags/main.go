package main

import (
	"flag"
	"fmt"
)

var (
	word string
	num  int
	fork bool
)

func main() {
	flag.StringVar(&word, "word", "foo", "a string")
	flag.IntVar(&num, "num", 42, "a number")
	flag.BoolVar(&fork, "fork", false, "a bool")


	flag.Parse()

	fmt.Printf("word: %v\n", word)
	fmt.Printf("num: %v\n", num)
	fmt.Printf("fork: %v\n", fork)
}
