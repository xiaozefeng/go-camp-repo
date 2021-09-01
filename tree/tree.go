package tree

// practice: design tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func New(val int) *Node {
	return &Node{
		Val: val,
	}
}

func (n *Node) PreOrder(f func(*Node)) {
	if n == nil {
		return
	}
	f(n)
	n.Left.PreOrder(f)
	n.Right.PreOrder(f)
}

func (n *Node) PostOrder(f func(*Node)) {
	if n == nil {
		return
	}
	n.Left.PostOrder(f)
	f(n)
	n.Right.PostOrder(f)
}

func (n *Node) InOrder(f func(*Node)) {
	if n == nil {
		return
	}
	n.Right.InOrder(f)
	n.Left.InOrder(f)
	f(n)
}
