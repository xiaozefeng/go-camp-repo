package algorithm

import "math"

type Node struct {
	Val      int
	Children []*Node
}

// implement with BFS
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var deepth int
	var Q = []*Node{root}
	for len(Q) >0{
		var size = len(Q)

		for i := 0; i < size; i++ {
			current := Q[0]
			Q = Q[1:]
			for _, v := range current.Children {
				Q = append(Q, v)
			}
		}
		deepth++
	}

	return deepth
}

// implement with DFS
func maxDepth2(root *Node) int {
	if root == nil {
		return 0
	}
	if len(root.Children) ==0 {
		return 1
	}
	heights:= make([]int, 0)
	for _, v := range root.Children {
		heights = append(heights, maxDepth2(v))
	}
	return max(heights)+1
}

func max(arr[]int) int{
	var res float64
	for _, v := range arr {
		res = math.Max(float64(v), res)
	}
	return int(res)
}