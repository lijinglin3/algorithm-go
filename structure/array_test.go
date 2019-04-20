package structure

import (
	"testing"
)

func TestNewArray(t *testing.T) {
	arr := NewArray(0)
	if arr != nil {
		t.Fatal("create new array failed")
	}
}

func TestArray_Insert(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity-1; i++ {
		if !arr.Insert(uint(i), i+1) {
			t.Fatal("insert fail")
		}
	}

	if arr.Insert(100, 100) {
		t.Fail()
	}

	if !arr.Insert(0, 100) {
		t.Fail()
	}
}

func TestArray_InsertToHead(t *testing.T) {
	arr := NewArray(1)
	if !arr.InsertToHead(1) {
		t.Fail()
	}
	if arr.InsertToHead(2) {
		t.Fail()
	}
}

func TestArray_InsertToTail(t *testing.T) {
	arr := NewArray(1)
	if !arr.InsertToTail(1) {
		t.Fatal("insert to tail failed")
	}
	if arr.InsertToTail(2) {
		t.Fatal("insert to tail failed")
	}
}

func TestArray_Delete(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		if !arr.Insert(uint(i), i+1) {
			t.Fatal("insert failed")
		}
	}
	t.Log(arr)

	for i := 0; i < capacity; i++ {
		if !arr.Delete(uint(i)) {
			t.Fatal("delete failed")
		}
	}
	t.Log(arr)

	if arr.Delete(100) {
		t.Fatal("delete failed")
	}
}

func TestArray_Find(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		if !arr.Insert(uint(i), i+1) {
			t.Fatal("insert fail")
		}
	}

	if arr.Find(0) != 1 {
		t.Fatal("find failed")
	}
	if arr.Find(9) != 10 {
		t.Fatal("find failed")
	}
	if arr.Find(11) != nil {
		t.Fatal("find failed")
	}
}
