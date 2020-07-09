package leetcode

// MyStack ...
type MyStack struct {
	data []int
}

// Constructor initialize your data structure here
func Constructor() MyStack {
	return MyStack{}
}

// Push push element x onto stack
func (s *MyStack) Push(x int) {
	s.data = append(s.data, x)
}

// Pop removes the element on top of the stack and returns that element
func (s *MyStack) Pop() int {
	if len(s.data) == 0 {
		return -1
	}
	n := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return n
}

// Top get the top element
func (s *MyStack) Top() int {
	if len(s.data) == 0 {
		return -1
	}
	return s.data[len(s.data)-1]
}

// Empty returns whether the stack is empty
func (s *MyStack) Empty() bool {
	return len(s.data) == 0
}
