package structure

import "testing"

func TestArrayQueue_EnQueue(t *testing.T) {
	q := NewArrayQueue(5)
	q.EnQueue(1)
	t.Log(q)
	q.EnQueue(2)
	t.Log(q)
	q.EnQueue(3)
	t.Log(q)
	q.EnQueue(4)
	t.Log(q)
	q.EnQueue(5)
	t.Log(q)
	q.EnQueue(6)
	t.Log(q)
}

func TestArrayQueue_DeQueue(t *testing.T) {
	q := NewArrayQueue(5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	t.Log(q)
	q.DeQueue()
	t.Log(q)
	q.DeQueue()
	t.Log(q)
	q.DeQueue()
	t.Log(q)
	q.DeQueue()
	t.Log(q)
	q.DeQueue()
	t.Log(q)
}
