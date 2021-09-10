package cryptos_test

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func Test_base64(t *testing.T) {

	data := `abc123!?*&^%$#@`
	enc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Printf("enc: %v\n", enc)

	dec, _ := base64.StdEncoding.DecodeString(enc)
	fmt.Printf("dec: %s\n", dec)

	URL:=`https://gobyexample.com/base64-encoding`
	uEnc :=base64.URLEncoding.EncodeToString([]byte(URL))
	fmt.Printf("uEnc: %v\n", uEnc)

	uDec,_ :=base64.URLEncoding.DecodeString(uEnc)
	fmt.Printf("uDec: %s\n", uDec)


}
