package leetcode

// MapSum ...
type MapSum struct {
	val      int
	children map[rune]*MapSum
}

// Constructor initialize your data structure here.
func Constructor() MapSum {
	return MapSum{children: make(map[rune]*MapSum)}
}

// Insert ...
func (s *MapSum) Insert(key string, val int) {
	tmp := s
	for _, v := range key {
		if tmp.children[v] == nil {
			tmp.children[v] = &MapSum{children: make(map[rune]*MapSum)}
		}
		tmp = tmp.children[v]
	}
	tmp.val = val
}

// Sum ...
func (s *MapSum) Sum(prefix string) int {
	tmp := s
	for _, v := range prefix {
		if tmp.children[v] == nil {
			return 0
		}
		tmp = tmp.children[v]
	}

	sum := 0
	queue := []*MapSum{tmp}
	for len(queue) != 0 {
		nq := make([]*MapSum, 0)
		for _, v := range queue {
			sum += v.val
			for _, c := range v.children {
				nq = append(nq, c)
			}
		}
		queue = nq
	}
	return sum
}
