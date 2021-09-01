package tree

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{name: "v1", args: args{1}, want: &Node{Val: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
