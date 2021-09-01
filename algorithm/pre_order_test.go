package algorithm

import (
	"reflect"
	"testing"
)

// func Test_preorderTraversal(t *testing.T) {
// 	var root =  TreeNode{
// 		Val: 1,
// 		Left: &TreeNode{Val: 2},
// 	}
// 	root.Left.Right = &TreeNode{Val: 3}
// 	var res = preorderTraversal(&root)
// 	fmt.Println(res)

// }

func Test_preorderTraversal(t *testing.T) {
	var root = TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
	}

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "v1", args: args{root: &root}, want: []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
