package algorithm

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var res []int
	preorder(root, func(tn *TreeNode) {
		res = append(res, tn.Val)
	})
	return res
}

func preorder(n *TreeNode, f func(*TreeNode)) {
	if n == nil {
		return
	}
	f(n)
	preorder(n.Left, f)
	preorder(n.Right, f)
}
