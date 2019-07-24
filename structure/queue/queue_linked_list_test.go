package queue

import "testing"

func TestListQueue_EnQueue(t *testing.T) {
	q := NewLinkedListQueue()

	t.Log(q)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	t.Log(q)
}

func TestListQueue_DeQueue(t *testing.T) {
	q := NewLinkedListQueue()

	t.Log(q)
	for i := 0; i < 3; i++ {
		q.EnQueue(i)
	}
	t.Log(q)

	if q.DeQueue() != 0 {
		t.Fatal("de queue failed")
	}

	if q.DeQueue() != 1 {
		t.Fatal("de queue failed")
	}

	if q.DeQueue() != 2 {
		t.Fatal("de queue failed")
	}

	if q.DeQueue() != nil {
		t.Fatal("de queue failed")
	}
}
