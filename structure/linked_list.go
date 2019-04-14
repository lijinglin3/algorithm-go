package structure

import (
	"fmt"
)

type ListNode struct {
	next  *ListNode
	value interface{}
}

type LinkedList struct {
	head   *ListNode
	length uint
}

func NewListNode(value interface{}) *ListNode {
	return &ListNode{nil, value}
}

func (node *ListNode) Next() *ListNode {
	return node.next
}

func (node *ListNode) Value() interface{} {
	return node.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

func (list *LinkedList) InsertAfter(node *ListNode, value interface{}) error {
	if nil == node {
		return ErrorNilPointerDereference
	}
	newNode := NewListNode(value)
	newNode.next, node.next = node.Next(), newNode
	list.length++
	return nil
}

func (list *LinkedList) InsertBefore(node *ListNode, value interface{}) error {
	if nil == node || list.head == node {
		return ErrorNilPointerDereference
	}
	cur := list.head.next
	preNode := list.head
	for cur != nil {
		if cur == node {
			break
		} else {
			cur = cur.next
			preNode = preNode.next
		}
	}
	if nil == cur {
		return ErrorNotFound
	}
	if err := list.InsertAfter(preNode, value); err != nil {
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

func (list *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= list.length {
		return nil
	}
	cur := list.head.next
	for i := uint(0); i < index; i++ {
		cur = cur.next
	}
	return cur
}

func (list *LinkedList) Delete(node *ListNode) error {
	if nil == node {
		return ErrorNilPointerDereference
	}
	cur := list.head.next
	pre := list.head
	for cur != nil {
		if cur == node {
			break
		}
		cur = cur.next
		pre = pre.next
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
		cur := list.head.next
		for cur != nil {
			str += fmt.Sprintf(" --> %+v", cur.value)
			cur = cur.next
		}
		return str
	}
}
