package leetcode

// MinStack MinStack
type MinStack struct {
	stack, min []int
	length     int
}

// Constructor initialize your data structure here.
func Constructor() MinStack {
	return MinStack{stack: make([]int, 0), min: make([]int, 0), length: 0}
}

// Push 入栈
func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)
	if s.length == 0 || s.min[s.length-1] > x {
		s.min = append(s.min, x)
	} else {
		s.min = append(s.min, s.min[s.length-1])
	}
	s.length = s.length + 1
}

// Pop 出栈
func (s *MinStack) Pop() {
	if s.length == 0 {
		return
	}
	s.stack = s.stack[:s.length-1]
	s.min = s.min[:s.length-1]
	s.length = s.length - 1
}

// Top 查询栈顶元素
func (s *MinStack) Top() int {
	if s.length == 0 {
		return 0
	}
	return s.stack[s.length-1]
}

// GetMin 返回栈中的最小值
func (s *MinStack) GetMin() int {
	if s.length == 0 {
		return 0
	}
	return s.min[s.length-1]
}
