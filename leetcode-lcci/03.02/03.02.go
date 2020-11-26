package leetcode

// MinStack ...
type MinStack struct {
	stack []node
}

type node struct {
	val int
	min int
}

// Constructor ...
func Constructor() MinStack {
	return MinStack{stack: make([]node, 0)}
}

// Push ...
func (s *MinStack) Push(x int) {
	min := x
	l := len(s.stack)
	if l > 0 && s.stack[l-1].min < min {
		min = s.stack[l-1].min
	}
	item := node{val: x, min: min}
	s.stack = append(s.stack, item)
}

// Pop ...
func (s *MinStack) Pop() {
	l := len(s.stack)
	s.stack = s.stack[:l-1]
}

// Top ...
func (s *MinStack) Top() int {
	l := len(s.stack)
	return s.stack[l-1].val
}

// GetMin ...
func (s *MinStack) GetMin() int {
	l := len(s.stack)
	return s.stack[l-1].min
}
