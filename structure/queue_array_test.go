package structure

import "testing"

func TestArrayQueue_EnQueue(t *testing.T) {
	q := NewArrayQueue(5)
	q.EnQueue(1)
	q.Print()
	q.EnQueue(2)
	q.Print()
	q.EnQueue(3)
	q.Print()
	q.EnQueue(4)
	q.Print()
	q.EnQueue(5)
	q.Print()
	q.EnQueue(6)
	q.Print()
}

func TestArrayQueue_DeQueue(t *testing.T) {
	q := NewArrayQueue(5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	q.Print()
	q.DeQueue()
	q.Print()
	q.DeQueue()
	q.Print()
	q.DeQueue()
	q.Print()
	q.DeQueue()
	q.Print()
	q.DeQueue()
	q.Print()
}
