package leetcode

// TripleInOne ...
type TripleInOne struct {
	size  int
	stack []int
}

// Constructor ...
func Constructor(stackSize int) TripleInOne {
	return TripleInOne{
		size:  stackSize,
		stack: make([]int, (stackSize+1)*3),
	}
}

// Push ...
func (o *TripleInOne) Push(stackNum int, value int) {
	sidx := (o.size + 1) * stackNum
	num := o.stack[sidx]
	if num >= o.size {
		return
	}
	o.stack[sidx+num+1] = value
	o.stack[sidx]++
}

// Pop ...
func (o *TripleInOne) Pop(stackNum int) int {
	sidx := (o.size + 1) * stackNum
	num := o.stack[sidx]
	if num == 0 {
		return -1
	}
	v := o.stack[sidx+num]
	o.stack[sidx]--
	o.stack[sidx+num] = 0
	return v
}

// Peek ...
func (o *TripleInOne) Peek(stackNum int) int {
	sidx := (o.size + 1) * stackNum
	num := o.stack[sidx]
	if num == 0 {
		return -1
	}
	return o.stack[sidx+num]
}

// IsEmpty ...
func (o *TripleInOne) IsEmpty(stackNum int) bool {
	sidx := (o.size + 1) * stackNum
	num := o.stack[sidx]
	return num == 0
}
