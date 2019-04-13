package structure

import "testing"

func TestInsertToHead(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		_ = l.InsertToHead(i + 1)
	}
	l.Print()
}

func TestInsertToTail(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		_ = l.InsertToTail(i + 1)
	}
	l.Print()
}

func TestFindByIndex(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		_ = l.InsertToTail(i + 1)
	}
	t.Log(l.FindByIndex(0))
	t.Log(l.FindByIndex(9))
	t.Log(l.FindByIndex(5))
	t.Log(l.FindByIndex(11))
}

func TestDeleteNode(t *testing.T) {
	list := NewLinkedList()
	for i := 0; i < 3; i++ {
		_ = list.InsertToTail(i + 1)
	}
	list.Print()

	t.Log(list.Delete(list.head.next))
	list.Print()

	t.Log(list.Delete(list.head.next.next))
	list.Print()
}
