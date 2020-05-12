package leetcode

type MinStack struct {
	stack, sort []int
	length      int
}

// Constructor initialize your data structure here.
func Constructor() MinStack {
	return MinStack{stack: make([]int, 0), sort: make([]int, 0), length: 0}
}

func (stack *MinStack) Push(x int) {
	stack.length = stack.length + 1
	stack.stack = append(stack.stack, x)
	stack.sort = append(stack.sort, x)
	for i := stack.length - 2; i >= 0; i-- {
		if stack.sort[i+1] > stack.sort[i] {
			stack.sort[i+1] = stack.sort[i]
			stack.sort[i] = x
		} else {
			break
		}
	}
}

func (stack *MinStack) Pop() {
	if stack.length == 0 {
		return
	}
	val := stack.stack[stack.length-1]
	for i, v := range stack.sort {
		if v == val {
			stack.sort = append(stack.sort[:i], stack.sort[i+1:]...)
			break
		}
	}
	stack.stack = stack.stack[:stack.length-1]
	stack.length = stack.length - 1
}

func (stack *MinStack) Top() int {
	if stack.length == 0 {
		return 0
	}
	return stack.stack[stack.length-1]
}

func (stack *MinStack) GetMin() int {
	if stack.length == 0 {
		return 0
	}
	return stack.sort[stack.length-1]
}
