package formatter

import (
	"fmt"
	"testing"
)

type point struct {
	x, y int
}

func Test_format(t *testing.T) {
	p := point{1, 2}
	fmt.Printf("v1: %v\n", p)
	fmt.Printf("v2: %+v\n", p)
	fmt.Printf("v3: %#v\n", p)

	fmt.Printf("type:%T\n", p)

	fmt.Printf("bin:%b\n", 14)
	fmt.Printf("char:%c\n", 34)

	fmt.Printf("hex:%x\n", 456)
	fmt.Printf("hex:%x\n", "hello")
}
