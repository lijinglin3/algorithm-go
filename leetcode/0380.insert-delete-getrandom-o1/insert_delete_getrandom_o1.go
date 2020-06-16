package leetcode

import "math/rand"

// RandomizedSet ...
type RandomizedSet struct {
	m map[int]int
	s []int
}

// Constructor initialize your data structure here.
func Constructor() RandomizedSet {
	return RandomizedSet{
		m: make(map[int]int),
		s: make([]int, 0),
	}
}

// Insert inserts a value to the set. Returns true if the set did not already contain the specified element.
func (set *RandomizedSet) Insert(val int) bool {
	if _, ok := set.m[val]; ok {
		return false
	}
	set.s = append(set.s, val)
	set.m[val] = len(set.s) - 1
	return true
}

// Remove removes a value from the set. Returns true if the set contained the specified element.
func (set *RandomizedSet) Remove(val int) bool {
	if index, ok := set.m[val]; ok {
		set.s[index] = set.s[len(set.s)-1]
		set.m[set.s[index]] = index
		set.s = set.s[:len(set.s)-1]
		delete(set.m, val)
		return true
	}
	return false
}

// GetRandom get a random element from the set.
func (set *RandomizedSet) GetRandom() int {
	i := rand.Int63n(int64(len(set.s)))
	return set.s[i]
}
