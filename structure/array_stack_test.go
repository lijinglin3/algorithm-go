package structure

import "testing"

func TestArrayStackPush(t *testing.T) {
	s := NewArrayStack(3)
	t.Log(s.Push(1))
	t.Log(s.Push(2))
	t.Log(s.Pop())
	t.Log(s.Push(3))
	t.Log(s.Push(4))
	t.Log(s.Push(5))
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	s.Print()
}

func TestArrayStackPop(t *testing.T) {
	s := NewArrayStack(3)
	_ = s.Push(1)
	_ = s.Push(2)
	_ = s.Push(3)
	s.Print()

	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	s.Print()
}

func TestArrayStackTop(t *testing.T) {
	s := NewArrayStack(3)
	_ = s.Push(1)
	_ = s.Push(2)
	_ = s.Push(3)

	t.Log(s.Top())
	s.Pop()
	t.Log(s.Top())
	s.Pop()
	t.Log(s.Top())
	s.Pop()
	t.Log(s.Top())
	s.Pop()
}
