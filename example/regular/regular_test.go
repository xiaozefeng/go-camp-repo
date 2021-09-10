package regular

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"
)

func Test_regular(t *testing.T) {
	matched, _ := regexp.MatchString(`p([z-z]+)ch`, "peach")
	fmt.Println("matche string:", matched)

	r, _ := regexp.Compile(`p([z-z]+)ch`)
	fmt.Println("matche string = peach:", r.MatchString("peach"))
	fmt.Println("find string:", r.FindString("peach punch"))
	fmt.Println("find string index:", r.FindStringIndex("peach punch"))
	fmt.Println("find string submatch:", r.FindStringSubmatch("peach punch"))

	fmt.Println("find all string:", r.FindAllString("peach punch pinch", -1))

	fmt.Println("find all string submatch index:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	fmt.Println("find all string:", r.FindAllString("peach punch pinch ", 2))

	fmt.Println("match bytes:", r.Match([]byte("peach")))
	r = regexp.MustCompile(`p([z-z]+)ch`)
	fmt.Println(r)

	fmt.Println("replace all string:", r.ReplaceAllString("a peach", "<fruit>"))
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in , bytes.ToUpper)
	fmt.Println(string(out))	
}
