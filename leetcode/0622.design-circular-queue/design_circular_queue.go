package leetcode

// MyCircularQueue ...
type MyCircularQueue struct {
	queue []int
	size  int
}

// Constructor initialize your data structure here. Set the size of the queue to be k
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, 0, k),
		size:  k,
	}
}

// EnQueue insert an element into the circular queue. Return true if the operation is successful
func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}
	q.queue = append(q.queue, value)
	return true
}

// DeQueue delete an element from the circular queue. Return true if the operation is successful
func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.queue = q.queue[1:]
	return true
}

// Front get the front item from the queue
func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.queue[0]
}

// Rear get the last item from the queue
func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.queue[len(q.queue)-1]
}

// IsEmpty checks whether the circular queue is empty or not
func (q *MyCircularQueue) IsEmpty() bool {
	return len(q.queue) == 0
}

// IsFull checks whether the circular queue is full or not
func (q *MyCircularQueue) IsFull() bool {
	return len(q.queue) >= q.size
}
