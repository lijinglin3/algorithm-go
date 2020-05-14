package queue

import "fmt"

// ArrayQueue 基于数组实现的消息队列
type ArrayQueue struct {
	data []interface{}
	head int
	tail int
}

// NewArrayQueue 初始化队列
func NewArrayQueue(capacity int) *ArrayQueue {
	if capacity <= 0 {
		return nil
	}
	return &ArrayQueue{
		data: make([]interface{}, capacity, capacity),
		head: 0,
		tail: 0,
	}
}

// Length 获取队列长度
func (queue *ArrayQueue) Length() int {
	return queue.tail - queue.head
}

// IsEmpty 判断队列是否为空
func (queue *ArrayQueue) IsEmpty() bool {
	return queue.tail == queue.head
}

// IsFull 判断队列是否已满
func (queue *ArrayQueue) IsFull() bool {
	return (queue.tail - queue.head) == cap(queue.data)
}

// EnQueue 入队列
func (queue *ArrayQueue) EnQueue(value interface{}) bool {
	if queue.IsFull() {
		return false
	}
	if queue.head != 0 {
		length := queue.Length()
		for i := 0; i < length; i++ {
			queue.data[i] = queue.data[i+queue.head]
		}
		queue.head, queue.tail = 0, length
	}
	queue.data[queue.tail] = value
	queue.tail++
	return true
}

// DeQueue 出队列
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
	}
	str := "head"
	for _, v := range queue.data[queue.head:queue.tail] {
		str += fmt.Sprintf(" --> %+v", v)
	}
	return str + " --> tail"
}
