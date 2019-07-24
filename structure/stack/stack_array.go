package stack

import "fmt"

type ArrayStack struct {
	data []interface{} // 数据
	top  int           // 栈顶指针
}

func NewArrayStack(capacity int) *ArrayStack {
	if capacity <= 0 {
		return nil
	}
	return &ArrayStack{
		data: make([]interface{}, capacity, capacity),
		top:  -1,
	}
}

func (stack *ArrayStack) IsEmpty() bool {
	if stack.top < 0 {
		return true
	}
	return false
}

func (stack *ArrayStack) Push(value interface{}) bool {
	if cap(stack.data) == stack.top+1 {
		return false
	}
	stack.top++
	stack.data[stack.top] = value
	return true
}

func (stack *ArrayStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	value := stack.data[stack.top]
	stack.top--
	return value
}

func (stack *ArrayStack) Top() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.data[stack.top]
}

func (stack *ArrayStack) Flush() {
	stack.top = -1
}

func (stack *ArrayStack) String() string {
	if stack.IsEmpty() {
		return fmt.Sprintln([]interface{}{})
	} else {
		return fmt.Sprintln(stack.data[0 : stack.top+1])
	}
}
