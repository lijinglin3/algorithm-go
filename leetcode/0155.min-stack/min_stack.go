package leetcode

// MinStack MinStack
type MinStack struct {
	stack  [][2]int
	length int
}

// Constructor initialize your data structure here.
func Constructor() MinStack {
	return MinStack{stack: make([][2]int, 0), length: 0}
}

// Push 入栈
func (s *MinStack) Push(x int) {
	min := x
	if s.length > 0 && s.GetMin() < x {
		min = s.GetMin()
	}
	s.stack = append(s.stack, [2]int{x, min})
	s.length++
}

// Pop 出栈
func (s *MinStack) Pop() {
	if s.length == 0 {
		return
	}
	s.stack = s.stack[:s.length-1]
	s.length = s.length - 1
}

// Top 查询栈顶元素
func (s *MinStack) Top() int {
	if s.length == 0 {
		return 0
	}
	return s.stack[s.length-1][0]
}

// GetMin 返回栈中的最小值
func (s *MinStack) GetMin() int {
	if s.length == 0 {
		return 0
	}
	return s.stack[s.length-1][1]
}
