package leetcode

// StackOfPlates ...
type StackOfPlates struct {
	stacks     [][]int
	cap, count int
}

// Constructor ...
func Constructor(cap int) StackOfPlates {
	return StackOfPlates{
		stacks: make([][]int, 0),
		cap:    cap,
		count:  -1,
	}
}

// Push ...
func (p *StackOfPlates) Push(val int) {
	if p.cap <= 0 {
		return
	}
	if p.count == -1 || len(p.stacks[p.count]) == p.cap {
		p.stacks = append(p.stacks, []int{val})
		p.count++
	} else {
		p.stacks[p.count] = append(p.stacks[p.count], val)
	}
}

// Pop ...
func (p *StackOfPlates) Pop() int {
	return p.PopAt(p.count)
}

// PopAt ...
func (p *StackOfPlates) PopAt(index int) int {
	val := -1
	if p.count == -1 || index > p.count {
		return val
	}
	if len(p.stacks[index]) != 0 {
		val = p.stacks[index][len(p.stacks[index])-1]
		p.stacks[index] = p.stacks[index][:len(p.stacks[index])-1]
	}
	if len(p.stacks[index]) == 0 {
		p.stacks = append(p.stacks[:index], p.stacks[index+1:]...)
		p.count--
	}
	return val
}
