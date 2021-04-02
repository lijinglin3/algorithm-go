package leetcode

// SortedStack ...
type SortedStack struct {
	data []int
}

// Constructor ...
func Constructor() SortedStack {
	return SortedStack{}
}

// Push ...
func (s *SortedStack) Push(val int) {
	tmp := make([]int, 0)
	for len(s.data) != 0 && s.data[len(s.data)-1] < val {
		tmp = append(tmp, s.data[len(s.data)-1])
		s.data = s.data[:len(s.data)-1]
	}
	s.data = append(s.data, val)
	for i := len(tmp) - 1; i >= 0; i-- {
		s.data = append(s.data, tmp[i])
	}
}

// Pop ...
func (s *SortedStack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.data = s.data[:len(s.data)-1]
}

// Peek ...
func (s *SortedStack) Peek() int {
	if s.IsEmpty() {
		return -1
	}
	return s.data[len(s.data)-1]
}

// IsEmpty ...
func (s *SortedStack) IsEmpty() bool {
	return len(s.data) == 0
}
