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

// implement with loop
func preorderTraversal2(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return res
}

// implement with morris

func preorderTraversal3(root *TreeNode) []int {
	var res []int
	var p1, p2 *TreeNode = root, nil
	for p1 != nil {
		p2 = p1.Left
		if p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				res = append(res, p1.Val)
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			p2.Right = nil
		} else {
			res = append(res, p1.Val)
		}
		p1 = p1.Right
	}
	return res
}
