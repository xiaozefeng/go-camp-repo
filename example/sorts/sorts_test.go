package sorts

import (
	"fmt"
	"sort"
	"testing"
)

func Test_sort1(t *testing.T) {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("strs", strs)

	ints := []int{7, 2, 4}
	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted", s)
	sort.Ints(ints)
	fmt.Println("ints", ints)
	s = sort.IntsAreSorted(ints)
	fmt.Println("sorted", s)
}


func Test_sort2(t *testing.T) {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println("fruits", fruits)
}

type byLength []string

func (a byLength) Len() int           { return len(a) }
func (a byLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byLength) Less(i, j int) bool { return a[i] < a[j] }
