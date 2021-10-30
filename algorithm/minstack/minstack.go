package minstack

import "math"

type Node struct {
	val  int
	min  int
	next *Node
}

type MinStack struct {
	head *Node
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.head == nil {
		this.head = &Node{val: val, min: val, next: nil}
	} else {
		this.head = &Node{val: val, min: int(math.Max(float64(val), float64(this.head.min))), next: this.head}
	}
}

func (this *MinStack) Pop() {
	this.head = this.head.next
}

func (this *MinStack) Top() int {
	return this.head.val
}

func (this *MinStack) GetMin() int {
	return this.head.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
