package structure

import "testing"

func TestNewCircularQueue(t *testing.T) {
	if q := NewCircularQueue(0); q != nil {
		t.Fatal("create new circular queue failed")
	}
}

func TestCircularQueue_EnQueue(t *testing.T) {
	q := NewCircularQueue(2)
	if !q.EnQueue(1) {
		t.Fatal("en queue failed")
	}

	if q.EnQueue(2) {
		t.Fatal("en queue failed")
	}
	t.Log(q)
}

func TestCircularQueue_DeQueue(t *testing.T) {
	q := NewCircularQueue(4)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	if q.DeQueue() != 1 {
		t.Fatal("de queue failed")
	}
	q.EnQueue(10)
	if q.DeQueue() != 2 {
		t.Fatal("de queue failed")
	}
	if q.DeQueue() != 3 {
		t.Fatal("de queue failed")
	}
	if q.DeQueue() != 10 {
		t.Fatal("de queue failed")
	}
	if q.DeQueue() != nil {
		t.Fatal("de queue failed")
	}
	t.Log(q)
}
