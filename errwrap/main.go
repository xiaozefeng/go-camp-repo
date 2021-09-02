package main

import (
	"errors"
	"fmt"
	"io/fs"
)

func main() {
	e := errors.New("origin error")
	wrappered := fmt.Errorf("addtional message, %w", e)
	if errors.Is(wrappered, e) {
		fmt.Println("wrappered error is origin error")
	} else {
		fmt.Println("wrappered error is  not  origin error")
	}
	// if w := errors.Unwrap(wrappered); w == e {
	// 	fmt.Println("decode err :", w)
	// } else {
	// 	fmt.Println("can not decode err")
	// }

	err := getPathError()
	var pathError *fs.PathError
	if errors.As(err, &pathError) {
		fmt.Println("occurred path error, the path:", pathError.Path)
	}

}

func getPathError() error {
	return &fs.PathError{Path: "/usr/local/xx_path"}
}
