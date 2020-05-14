package stack

import "fmt"

// ArrayStack 基于数组实现的栈
type ArrayStack struct {
	data []interface{} // 数据
	top  int           // 栈顶指针
}

// NewArrayStack 初始化栈
func NewArrayStack(capacity int) *ArrayStack {
	if capacity <= 0 {
		return nil
	}
	return &ArrayStack{
		data: make([]interface{}, capacity, capacity),
		top:  -1,
	}
}

// IsEmpty 判断栈是否为空
func (stack *ArrayStack) IsEmpty() bool {
	if stack.top < 0 {
		return true
	}
	return false
}

// Push 入栈
func (stack *ArrayStack) Push(value interface{}) bool {
	if cap(stack.data) == stack.top+1 {
		return false
	}
	stack.top++
	stack.data[stack.top] = value
	return true
}

// Pop 出栈
func (stack *ArrayStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	value := stack.data[stack.top]
	stack.top--
	return value
}

// Top 查询栈顶元素
func (stack *ArrayStack) Top() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.data[stack.top]
}

// Flush 清空栈
func (stack *ArrayStack) Flush() {
	stack.top = -1
}

func (stack *ArrayStack) String() string {
	if stack.IsEmpty() {
		return fmt.Sprintln([]interface{}{})
	}
	return fmt.Sprintln(stack.data[0 : stack.top+1])

}
