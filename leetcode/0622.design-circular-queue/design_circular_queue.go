package leetcode

// MyCircularQueue ...
type MyCircularQueue struct {
	head, count, size int
	queue             []int
}

// Constructor initialize your data structure here. Set the size of the queue to be k
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, k),
		size:  k,
	}
}

// EnQueue insert an element into the circular queue. Return true if the operation is successful
func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}
	q.queue[(q.head+q.count)%q.size] = value
	q.count++
	return true
}

// DeQueue delete an element from the circular queue. Return true if the operation is successful
func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.head = (q.head + 1) % q.size
	q.count--
	return true
}

// Front get the front item from the queue
func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.queue[q.head]
}

// Rear get the last item from the queue
func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.queue[(q.head+q.count-1)%q.size]
}

// IsEmpty checks whether the circular queue is empty or not
func (q *MyCircularQueue) IsEmpty() bool {
	return q.count == 0
}

// IsFull checks whether the circular queue is full or not
func (q *MyCircularQueue) IsFull() bool {
	return q.count == q.size
}
