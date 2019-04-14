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
	return &LinkedList{&listNode{nil, 0}, 0}
}

func (list *LinkedList) InsertAfter(node *listNode, value interface{}) error {
	if nil == node {
		return ErrorNilPointerDereference
	}
	newNode := &listNode{nil, value}
	newNode.next, node.next = node.next, newNode
	list.length++
	return nil
}

func (list *LinkedList) InsertBefore(node *listNode, value interface{}) error {
	if nil == node || list.head == node {
		return ErrorNilPointerDereference
	}
	cur := list.head.next
	pre := list.head
	for ; cur != nil; cur, pre = cur.next, pre.next {
		if cur == node {
			break
		}
	}
	if nil == cur {
		return ErrorNotFound
	}
	if err := list.InsertAfter(pre, value); err != nil {
		return err
	}
	return nil
}

func (list *LinkedList) InsertToHead(value interface{}) error {
	return list.InsertAfter(list.head, value)
}

func (list *LinkedList) InsertToTail(value interface{}) error {
	cur := list.head
	for cur.next != nil {
		cur = cur.next
	}
	return list.InsertAfter(cur, value)
}

func (list *LinkedList) FindByIndex(index uint) *listNode {
	if index >= list.length {
		return nil
	}
	cur := list.head.next
	for i := uint(0); i < index; i++ {
		cur = cur.next
	}
	return cur
}

func (list *LinkedList) Delete(node *listNode) error {
	if nil == node {
		return ErrorNilPointerDereference
	}
	cur := list.head.next
	pre := list.head
	for ; cur != nil; cur, pre = cur.next, pre.next {
		if cur == node {
			break
		}
	}
	if nil == cur {
		return ErrorNotFound
	}
	pre.next = cur.next
	node = nil
	list.length--
	return nil
}

func (list *LinkedList) String() string {
	if nil == list.head.next {
		return ""
	} else {
		str := "top"
		for cur := list.head.next; cur != nil; cur = cur.next {
			str += fmt.Sprintf(" --> %+v", cur.value)
		}
		return str
	}
}
