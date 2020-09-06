package queue

import "fmt"

// CircularQueue 循环队列
type CircularQueue struct {
	data []interface{}
	head int
	tail int
}

// NewCircularQueue 初始化循环队列
func NewCircularQueue(capacity int) *CircularQueue {
	if capacity == 0 {
		return nil
	}
	return &CircularQueue{make([]interface{}, capacity), 0, 0}
}

// IsEmpty 判断队列是否为空
func (queue *CircularQueue) IsEmpty() bool {
	return queue.head == queue.tail
}

// IsFull 判断队列是否已满
func (queue *CircularQueue) IsFull() bool {
	return queue.head == (queue.tail+1)%cap(queue.data)
}

// EnQueue 入队列
func (queue *CircularQueue) EnQueue(v interface{}) bool {
	if queue.IsFull() {
		return false
	}
	queue.data[queue.tail] = v
	queue.tail = (queue.tail + 1) % cap(queue.data)
	return true
}

// DeQueue 出队列
func (queue *CircularQueue) DeQueue() interface{} {
	if queue.IsEmpty() {
		return nil
	}
	v := queue.data[queue.head]
	queue.head = (queue.head + 1) % cap(queue.data)
	return v
}

func (queue *CircularQueue) String() string {
	if queue.IsEmpty() {
		return "empty queue"
	}
	result := "head"
	for i := queue.head; true; {
		result += fmt.Sprintf(" --> %+v", queue.data[i])
		i = (i + 1) % cap(queue.data)
		if i == queue.tail {
			break
		}
	}
	return result + " --> tail"
}
