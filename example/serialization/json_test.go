package serialization

import (
	"bytes"
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

	v, _ = json.Marshal(1)
	fmt.Printf("v: %s\n", v)

	v, _ = json.Marshal(3.14)
	fmt.Printf("v: %s\n", v)

	fruits := []string{"apple", "peach", "pear"}
	v, _ = json.Marshal(fruits)
	fmt.Printf("v: %s\n", v)

	v, _ = json.Marshal(map[string]int{"apple": 5, "lettuce": 7})
	fmt.Printf("v: %s\n", v)

	v, _ = json.Marshal(&r1{
		Page:   1,
		Fruits: fruits,
	})
	fmt.Printf("v: %s\n", v)

	v, _ = json.Marshal(&r2{
		Page:   2,
		Fruits: fruits,
	})
	fmt.Printf("v: %s\n", v)

	data := ` {"page":2,"fruits":["apple","peach","pear"]}`

	var r1 r1
	_ = json.Unmarshal([]byte(data), &r1)
	fmt.Printf("r1: %+v\n", r1)

	var r2 r2
	decoder := json.NewDecoder(bytes.NewBufferString(data))
	_ = decoder.Decode(&r2)
	fmt.Printf("r2: %+v\n", r2)

}
