package linked_list

import "testing"

func TestLinkedList_Insert(t *testing.T) {
	l := NewLinkedList()
	if !l.IsEmpty() {
		t.Fatal()
	}
	if !l.Insert(0, 0) {
		t.Fatal("insert failed")
	}
	if l.IsEmpty() {
		t.Fatal()
	}
	if l.Insert(100, 100) {
		t.Fatal("insert failed")
	}
}

func TestLinkedList_InsertToTail(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		if !l.InsertToTail(i + 1) {
			t.Fatal("insert to tail failed")
		}
	}
}

func TestLinkedList_InsertToHead(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		if !l.InsertToHead(i + 1) {
			t.Fatal("insert to head failed")
		}
	}
}

func TestLinkedList_Find(t *testing.T) {
	l := NewLinkedList()
	t.Log(l)
	for i := 0; i < 10; i++ {
		_ = l.InsertToTail(i + 1)
	}
	t.Log(l)
	if 1 != l.Find(0) {
		t.Fatal()
	}
	if 10 != l.Find(9) {
		t.Fatal()
	}
	if nil != l.Find(11) {
		t.Fatal()
	}
}

func TestLinkedList_Delete(t *testing.T) {
	list := NewLinkedList()
	for i := 0; i < 3; i++ {
		list.InsertToTail(i + 1)
	}

	if !list.Delete(1) {
		t.Fatal()
	}
	if !list.Delete(0) {
		t.Fatal()
	}
	if list.Delete(3) {
		t.Fatal()
	}
}
