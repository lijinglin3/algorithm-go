package structure

import "fmt"

type ArrayQueue struct {
	data []interface{}
	head int
	tail int
}

func NewArrayQueue(capacity uint) *ArrayQueue {
	if capacity == 0 {
		return nil
	} else {
		return &ArrayQueue{
			data: make([]interface{}, capacity, capacity),
			head: 0,
			tail: 0,
		}
	}
}

func (queue *ArrayQueue) Length() int {
	return queue.tail - queue.head
}

func (queue *ArrayQueue) IsEmpty() bool {
	return queue.tail == queue.head
}

func (queue *ArrayQueue) IsFull() bool {
	return queue.tail == cap(queue.data)
}

func (queue *ArrayQueue) EnQueue(value interface{}) bool {
	if queue.IsFull() {
		return false
	}
	queue.data[queue.tail] = value
	queue.tail++
	if queue.head != 0 {
		length := queue.Length()
		for i := 0; i < length; i++ {
			queue.data[i] = queue.data[i+queue.head]
		}
		queue.head, queue.tail = 0, length
	}
	return true
}

func (queue *ArrayQueue) DeQueue() interface{} {
	if queue.IsEmpty() {
		return nil
	}
	value := queue.data[queue.head]
	queue.head++
	return value
}

func (queue *ArrayQueue) String() string {
	if queue.IsEmpty() {
		return "empty queue"
	} else {
		str := "head"
		for _, v := range queue.data[queue.head:queue.tail] {
			str += fmt.Sprintf(" --> %+v", v)
		}
		return str + " --> tail"
	}
}
