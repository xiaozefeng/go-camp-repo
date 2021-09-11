package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	isRecursion bool
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("expected run subcommands")
		return
	}
	runCmd := flag.NewFlagSet("run", flag.ExitOnError)
	runCmd.BoolVar(&isRecursion, "r", false, "recursion") //usage
	runCmd.Parse(os.Args[2:])
	fmt.Printf("isRecursion: %v\n", isRecursion)
}
