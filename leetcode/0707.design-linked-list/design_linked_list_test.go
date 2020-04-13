package leetcode

import (
	"testing"
)

func TestMyLinkedList(t *testing.T) {
	list := Constructor()

	list.AddAtHead(2)
	list.AddAtTail(4)
	list.AddAtIndex(-1, 0)
	list.AddAtIndex(1, 3)
	list.AddAtIndex(0, 1)
	list.AddAtIndex(4, 5)
	list.AddAtIndex(10, 10)

	list.Get(-1)
	list.Get(0)
	list.Get(2)
	list.Get(4)
	list.Get(5)

	list.DeleteAtIndex(-1)
	list.DeleteAtIndex(5)
	list.DeleteAtIndex(6)
	list.DeleteAtIndex(4)
	list.DeleteAtIndex(0)
	list.DeleteAtIndex(2)
}
