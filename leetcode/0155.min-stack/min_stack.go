package leetcode

type MinStack struct {
	stack, sort []int
	length      int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{stack: make([]int, 0), sort: make([]int, 0), length: 0}
}

func (this *MinStack) Push(x int) {
	this.length = this.length + 1
	this.stack = append(this.stack, x)
	this.sort = append(this.sort, x)
	for i := this.length - 2; i >= 0; i-- {
		if this.sort[i+1] > this.sort[i] {
			this.sort[i+1] = this.sort[i]
			this.sort[i] = x
		} else {
			break
		}
	}
}

func (this *MinStack) Pop() {
	if this.length == 0 {
		return
	}
	val := this.stack[this.length-1]
	for i, v := range this.sort {
		if v == val {
			this.sort = append(this.sort[:i], this.sort[i+1:]...)
			break
		}
	}
	this.stack = this.stack[:this.length-1]
	this.length = this.length - 1
}

func (this *MinStack) Top() int {
	if this.length == 0 {
		return 0
	}
	return this.stack[this.length-1]
}

func (this *MinStack) GetMin() int {
	if this.length == 0 {
		return 0
	}
	return this.sort[this.length-1]
}
