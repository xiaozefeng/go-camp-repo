package bfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res = make([][]int, 0)
	if root == nil {
		return res
	}
	var Q = []*TreeNode{root}
	for len(Q) > 0 {
		var level = make([]int, 0)
		var size = len(Q)
		for i := 0; i < size; i++ {
			var node = Q[0]
			Q = Q[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				Q = append(Q, node.Left)
			}
			if node.Right != nil {
				Q = append(Q, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
