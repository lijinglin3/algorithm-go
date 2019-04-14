package structure

import (
	"testing"
)

func TestInsert(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity-2; i++ {

		if !arr.Insert(uint(i), i+1) {
			t.Fatal("insert fail")
		}
	}
	t.Log(arr)

	_ = arr.Insert(uint(6), 999)
	t.Log(arr)

	_ = arr.InsertToTail(666)
	t.Log(arr)
}

func TestDelete(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		if !arr.Insert(uint(i), i+1) {
			t.Fatal("insert fail")
		}
	}
	t.Log(arr)

	for i := 9; i >= 0; i-- {
		if !arr.Delete(uint(i)) {
			t.Fatal("delete fail")
		}
		t.Log(arr)
	}
}

func TestFind(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		if !arr.Insert(uint(i), i+1) {
			t.Fatal("insert fail")
		}
	}
	t.Log(arr)

	t.Log(arr.Find(0))
	t.Log(arr.Find(9))
	t.Log(arr.Find(11))
}
