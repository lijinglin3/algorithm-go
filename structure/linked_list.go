package structure

import (
	"fmt"
)

type listNode struct {
	next  *listNode
	value interface{}
}

type LinkedList struct {
	head   *listNode
	length uint
}

func NewLinkedList() *LinkedList {
	return &LinkedList{nil, 0}
}

func (list *LinkedList) isIndexOutOfRange(index uint) bool {
	if index >= list.length {
		return true
	}
	return false
}

func (list *LinkedList) Insert(index uint, value interface{}) bool {
	if 0 == index {
		list.head = &listNode{list.head, value}
	} else {
		if list.isIndexOutOfRange(index - 1) {
			return false
		}
		node := list.findNodeByIndex(index - 1)
		node.next = &listNode{node.next, value}
	}
	list.length++
	return true
}

func (list *LinkedList) IsEmpty() bool {
	if nil == list.head {
		return true
	}
	return false
}

func (list *LinkedList) InsertToHead(value interface{}) bool {
	return list.Insert(0, value)
}

func (list *LinkedList) InsertToTail(value interface{}) bool {
	return list.Insert(list.length, value)
}

func (list *LinkedList) findNodeByIndex(index uint) *listNode {
	if list.isIndexOutOfRange(index) {
		return nil
	}
	cur := list.head
	for i := uint(0); i < index; i++ {
		cur = cur.next
	}
	return cur
}

func (list *LinkedList) Find(index uint) interface{} {
	node := list.findNodeByIndex(index)
	if nil == node {
		return nil
	} else {
		return node.value
	}
}

func (list *LinkedList) Delete(index uint) bool {
	if 0 == index {
		list.head = list.head.next
	} else {
		if list.isIndexOutOfRange(index - 1) {
			return false
		}
		node := list.findNodeByIndex(index - 1)
		node.next = node.next.next
	}
	list.length--
	return true
}

func (list *LinkedList) String() string {
	if nil == list.head {
		return ""
	} else {
		str := "head"
		for cur := list.head; cur != nil; cur = cur.next {
			str += fmt.Sprintf(" --> %+v", cur.value)
		}
		return str + " --> tail"
	}
}
