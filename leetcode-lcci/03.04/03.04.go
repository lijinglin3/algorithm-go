package leetcode

// MyQueue ...
type MyQueue struct {
	nums []int
}

// Constructor initialize your data structure here.
func Constructor() MyQueue {
	return MyQueue{
		nums: make([]int, 0, 10),
	}
}

// Push push element x to the back of queue.
func (q *MyQueue) Push(x int) {
	q.nums = append(q.nums, x)
}

// Pop removes the element from in front of queue and returns that element.
func (q *MyQueue) Pop() int {
	res := q.nums[0]
	q.nums = q.nums[1:]
	return res
}

// Peek get the front element. */
func (q *MyQueue) Peek() int {
	return q.nums[0]
}

// Empty returns whether the queue is empty.
func (q *MyQueue) Empty() bool {
	return len(q.nums) <= 0
}
