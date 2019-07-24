package queue

import "testing"

func TestNewArrayQueue(t *testing.T) {
	if q := NewArrayQueue(0); q != nil {
		t.Fail()
	}
}

func TestArrayQueue_EnQueue(t *testing.T) {
	q := NewArrayQueue(5)
	t.Log(q)

	if !q.EnQueue(1) {
		t.Fatal("en queue failed")
	}

	if !q.EnQueue(2) {
		t.Fatal("en queue failed")
	}

	if !q.EnQueue(3) {
		t.Fatal("en queue failed")
	}

	if !q.EnQueue(4) {
		t.Fatal("en queue failed")
	}

	if !q.EnQueue(5) {
		t.Fatal("en queue failed")
	}
	t.Log(q)

	if q.EnQueue(6) {
		t.Fatal("en queue failed")
	}
	t.Log(q)
}

func TestArrayQueue_DeQueue(t *testing.T) {
	q := NewArrayQueue(3)
	for i := 0; i < 10; i++ {
		_ = q.EnQueue(i)
	}

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

	_ = q.EnQueue(0)
	_ = q.EnQueue(1)

	if q.DeQueue() != 0 {
		t.Fatal("de queue failed")
	}

	_ = q.EnQueue(2)

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
