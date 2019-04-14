package structure

import "fmt"

type queueNode struct {
	val  interface{}
	next *queueNode
}

type LinkedListQueue struct {
	head   *queueNode
	tail   *queueNode
	length int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{nil, nil, 0}
}

func (this *LinkedListQueue) EnQueue(value interface{}) {
	node := &queueNode{value, nil}
	if nil == this.tail {
		this.tail = node
		this.head = node
	} else {
		this.tail.next = node
		this.tail = node
	}
	this.length++
}

func (this *LinkedListQueue) DeQueue() interface{} {
	if this.head == nil {
		return nil
	}
	v := this.head.val
	this.head = this.head.next
	this.length--
	return v
}

func (this *LinkedListQueue) String() string {
	if this.head == nil {
		return "empty queue"
	} else {
		result := "head"
		for cur := this.head; cur != nil; cur = cur.next {
			result += fmt.Sprintf(" --> %+v", cur.val)
		}
		result += "--> tail"
		return result
	}
}
