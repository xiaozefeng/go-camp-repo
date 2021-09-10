package defers

import (
	"fmt"
	"os"
	"testing"
)

func Test_create_file(t *testing.T) {
	f := createFile("./defer.txt")
	defer closeFile(f)
	appendContent(f, "some data")
}


func createFile(path string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return f
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	if err := f.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "error:%v\n", err)
		os.Exit(1)
	}
}

func appendContent(f *os.File, content string) {
	fmt.Println("append content")
	fmt.Fprintln(f, content)
}
