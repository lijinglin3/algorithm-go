package leetcode

// MyHashSet ...
type MyHashSet struct {
	data map[int]interface{}
}

// Constructor initialize your data structure here
func Constructor() MyHashSet {
	return MyHashSet{data: make(map[int]interface{})}
}

// Add ...
func (s *MyHashSet) Add(key int) {
	s.data[key] = struct{}{}
}

// Remove ...
func (s *MyHashSet) Remove(key int) {
	delete(s.data, key)
}

// Contains returns true if this set contains the specified element
func (s *MyHashSet) Contains(key int) bool {
	_, ok := s.data[key]
	return ok
}
