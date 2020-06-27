package leetcode

import "math"

// MinStack MinStack
type MinStack struct {
	stack, min []int
}

// Constructor initialize your data structure here.
func Constructor() MinStack {
	return MinStack{stack: []int{}, min: []int{math.MaxInt64}}
}

// Push 入栈
func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)
	top := s.min[len(s.min)-1]
	if x < top {
		s.min = append(s.min, x)
	} else {
		s.min = append(s.min, top)
	}

}

// Pop 出栈
func (s *MinStack) Pop() {
	if len(s.stack) == 0 {
		return
	}
	s.stack = s.stack[:len(s.stack)-1]
	s.min = s.min[:len(s.min)-1]
}

// Top 查询栈顶元素
func (s *MinStack) Top() int {
	if len(s.stack) == 0 {
		return 0
	}
	return s.stack[len(s.stack)-1]
}

// Min 返回栈中的最小值
func (s *MinStack) Min() int {
	if len(s.stack) == 0 {
		return 0
	}
	return s.min[len(s.min)-1]
}
