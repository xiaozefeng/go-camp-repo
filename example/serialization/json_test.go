package serialization

import (
	"encoding/json"
	"fmt"
	"testing"
)

type r1 struct {
	Page   int
	Fruits []string
}

type r2 struct {
	Page   int      `json:"page,omitempty"`
	Fruits []string `json:"fruits,omitempty"`
}

func Test_json(t *testing.T) {
	v, _ := json.Marshal(true)
	fmt.Printf("v: %s\n", v)

	v,_=json.Marshal(1)
	fmt.Printf("v: %s\n", v)

	v,_=json.Marshal(3.14)
	fmt.Printf("v: %s\n", v)

	v,_=json.Marshal([]string{"apple", "peach", "pear"})
	fmt.Printf("v: %s\n", v)

	v,_=json.Marshal(map[string]int {"apple":5, "lettuce": 7})
	fmt.Printf("v: %s\n", v)


	



}
