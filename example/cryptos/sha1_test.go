package cryptos_test

import (
	"crypto/sha1"
	"fmt"
	"testing"
)

func Test_sha1(t *testing.T){ 
	h:=sha1.New()
	h.Write([]byte("sha1 this string"))

	bs:=h.Sum(nil)
	fmt.Printf("bs: %x\n", bs)
}