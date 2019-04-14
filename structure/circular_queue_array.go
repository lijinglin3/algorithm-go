package structure

import "fmt"

type CircularQueue struct {
	data []interface{}
	head int
	tail int
}

func NewCircularQueue(capacity int) *CircularQueue {
	if capacity == 0 {
		return nil
	}
	return &CircularQueue{make([]interface{}, capacity, capacity), 0, 0}
}

func (queue *CircularQueue) IsEmpty() bool {
	if queue.head == queue.tail {
		return true
	}
	return false
}

func (queue *CircularQueue) IsFull() bool {
	if queue.head == (queue.tail+1)%cap(queue.data) {
		return true
	}
	return false
}

func (queue *CircularQueue) EnQueue(v interface{}) bool {
	if queue.IsFull() {
		return false
	}
	queue.data[queue.tail] = v
	queue.tail = (queue.tail + 1) % cap(queue.data)
	return true
}

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
