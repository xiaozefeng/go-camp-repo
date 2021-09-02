package composite

import "testing"

func Test_ParentAndSub(t *testing.T) {
	var s = Sub{
		Parent{},
	}
	s.SayHello()
}