package leetcode

// MyHashMap ...
type MyHashMap struct {
	data map[int]int
}

// Constructor initialize your data structure here
func Constructor() MyHashMap {
	return MyHashMap{data: make(map[int]int)}
}

// Put put a value into map
func (m *MyHashMap) Put(key int, value int) {
	m.data[key] = value
}

// Get returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key
func (m *MyHashMap) Get(key int) int {
	if v, ok := m.data[key]; ok {
		return v
	}
	return -1
}

// Remove removes the mapping of the specified value key if this map contains a mapping for the key
func (m *MyHashMap) Remove(key int) {
	delete(m.data, key)
}
