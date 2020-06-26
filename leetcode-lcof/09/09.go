package leetcode

// CQueue ...
type CQueue struct {
	head, tail []int
}

// Constructor ...
func Constructor() CQueue {
	return CQueue{
		head: make([]int, 0),
		tail: make([]int, 0),
	}
}

// AppendTail ...
func (q *CQueue) AppendTail(value int) {
	q.tail = append(q.tail, value)
}

// DeleteHead ...
func (q *CQueue) DeleteHead() int {
	if len(q.head) == 0 {
		if len(q.tail) == 0 {
			return -1
		}
		for i := len(q.tail) - 1; i >= 0; i-- {
			q.head = append(q.head, q.tail[i])
		}
		q.tail = make([]int, 0)
	}

	n := q.head[len(q.head)-1]
	q.head = q.head[:len(q.head)-1]
	return n
}
