package closures

import (
	"fmt"
	"testing"
)

func Test_intSeq(t *testing.T) {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSeq()
	fmt.Println(newInt())
}
