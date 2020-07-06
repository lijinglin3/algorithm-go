package leetcode

// MaxQueue ...
type MaxQueue struct {
	data, max []int
}

// Constructor ...
func Constructor() MaxQueue {
	return MaxQueue{}
}

// MaxValue ...
func (q *MaxQueue) MaxValue() int {
	if len(q.max) == 0 {
		return -1
	}
	return q.max[0]
}

// PushBack ...
func (q *MaxQueue) PushBack(value int) {
	for i := 0; i < len(q.max); i++ {
		if q.max[i] < value {
			q.max = q.max[:i]
			break
		}
	}
	q.max = append(q.max, value)
	q.data = append(q.data, value)
}

// PopFront ...
func (q *MaxQueue) PopFront() int {
	if len(q.data) > 0 {
		n := q.data[0]
		if n == q.max[0] {
			q.max = q.max[1:]
		}
		q.data = q.data[1:]
		return n
	}
	return -1
}
