package leetcode

// MyQueue ...
type MyQueue struct {
	in, out []int
}

// Constructor initialize your data structure here
func Constructor() MyQueue {
	return MyQueue{}
}

// Push push element x to the back of queue
func (q *MyQueue) Push(x int) {
	q.in = append(q.in, x)
}

// Pop removes the element from in front of queue and returns that element
func (q *MyQueue) Pop() int {
	if len(q.out) == 0 {
		if len(q.in) != 0 {
			q.out = make([]int, 0, len(q.in))
			for i := len(q.in) - 1; i >= 0; i-- {
				q.out = append(q.out, q.in[i])
			}
			q.in = make([]int, 0)
		} else {
			return -1
		}
	}
	n := q.out[len(q.out)-1]
	q.out = q.out[:len(q.out)-1]
	return n
}

// Peek get the front element
func (q *MyQueue) Peek() int {
	if len(q.out) != 0 {
		return q.out[len(q.out)-1]
	}
	if len(q.in) != 0 {
		return q.in[0]
	}

	return -1
}

// Empty returns whether the queue is empty
func (q *MyQueue) Empty() bool {
	return len(q.in)+len(q.out) == 0
}
