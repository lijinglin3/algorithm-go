package leetcode

// FreqStack ...
type FreqStack struct {
	freq  map[int]int
	group map[int][]int
	max   int
}

// Constructor ...
func Constructor() FreqStack {
	return FreqStack{
		freq:  make(map[int]int),
		group: make(map[int][]int),
		max:   0,
	}
}

// Push ...
func (s *FreqStack) Push(x int) {
	s.freq[x]++
	f := s.freq[x]
	if f > s.max {
		s.max = f
	}
	s.group[f] = append(s.group[f], x)
}

// Pop ...
func (s *FreqStack) Pop() int {
	length := len(s.group[s.max])
	f := s.group[s.max][length-1]
	s.group[s.max] = s.group[s.max][:length-1]
	s.freq[f]--
	if length == 1 {
		s.max--
	}
	return f
}
