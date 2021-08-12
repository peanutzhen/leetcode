package main

type MinStack struct {
	stack []int
	min   []int // 记录当前最小元素的栈
}

/** initialize your data structure here. */
func Constructor() MinStack {
	minStack := MinStack{
		[]int{},
		[]int{},
	}
	return minStack
}

func (this *MinStack) Push(val int) {
	if len(this.min) == 0 || val <= this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	if this.min[len(this.min)-1] == this.stack[len(this.stack)-1] {
		this.min = this.min[:len(this.min)-1] // pop min
	}
	this.stack = this.stack[:len(this.stack)-1] // pop stack
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
