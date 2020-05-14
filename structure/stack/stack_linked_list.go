package stack

import "fmt"

type node struct {
	next  *node
	value interface{}
}

// LinkedListStack 基于链表实现的栈
type LinkedListStack struct {
	top *node // 栈顶节点
}

// NewLinkedListStack 初始化栈
func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

// IsEmpty 判断栈是否为空
func (stack *LinkedListStack) IsEmpty() bool {
	if nil == stack.top {
		return true
	}
	return false
}

// Push 入栈
func (stack *LinkedListStack) Push(value interface{}) {
	stack.top = &node{stack.top, value}
}

// Pop 出栈
func (stack *LinkedListStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	value := stack.top.value
	stack.top = stack.top.next
	return value
}

// Top 查询栈顶元素
func (stack *LinkedListStack) Top() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.top.value
}

// Flush 清空栈
func (stack *LinkedListStack) Flush() {
	stack.top = nil
}

func (stack *LinkedListStack) String() string {
	if stack.IsEmpty() {
		return "empty stack"
	}
	str := "top"
	for cur := stack.top; cur != nil; cur = cur.next {
		str += fmt.Sprintf(" --> %+v", cur.value)
	}
	return str + " --> tail"

}
