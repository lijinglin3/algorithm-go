package stack

import "testing"

func TestNewArrayStack(t *testing.T) {
	if NewArrayStack(-1) != nil {
		t.Fatal("new arrary stack failed")
	}
}

func TestArrayStackPush(t *testing.T) {
	s := NewArrayStack(3)
	if !s.Push(1) {
		t.Fatal("push failed")
	}
	if !s.Push(2) {
		t.Fatal("push failed")
	}
	if !s.Push(3) {
		t.Fatal("push failed")
	}
	t.Log(s)
}

func TestArrayStackPop(t *testing.T) {
	s := NewArrayStack(3)
	for i := 0; i < 10; i++ {
		s.Push(i)
	}
	t.Log(s)

	if s.Pop() != 2 {
		t.Fatal("pop failed")
	}
	if s.Pop() != 1 {
		t.Fatal("pop failed")
	}
	if s.Pop() != 0 {
		t.Fatal("pop failed")
	}
	if s.Pop() != nil {
		t.Fatal("pop failed")
	}
	t.Log(s)
}

func TestArrayStackTop(t *testing.T) {
	s := NewArrayStack(3)
	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	if s.Top() != 2 {
		t.Fatal("top failed")
	}
	s.Pop()
	if s.Top() != 1 {
		t.Fatal("top failed")
	}
	s.Pop()
	if s.Top() != 0 {
		t.Fatal("top failed")
	}
	s.Pop()
	if s.Top() != nil {
		t.Fatal("top failed")
	}
}

func TestArrayStack_Flush(t *testing.T) {
	s := NewArrayStack(3)
	for i := 0; i < 10; i++ {
		s.Push(i)
	}
	t.Log(s)

	if s.Flush(); !s.IsEmpty() {
		t.Fatal("flush failed")
	}
}
