package linkedlist

import (
	"fmt"
)

type node struct {
	next  *node
	value interface{}
}

// LinkedList 链表
type LinkedList struct {
	head   *node
	length int
}

// NewLinkedList 初始化一个链表
func NewLinkedList() *LinkedList {
	return &LinkedList{nil, 0}
}

func (list *LinkedList) outOfRange(index int) bool {
	if index >= list.length || index < 0 {
		return true
	}
	return false
}

// IsEmpty 判断链表是否为空
func (list *LinkedList) IsEmpty() bool {
	if nil == list.head {
		return true
	}
	return false
}

// Insert 在指定位置插入一个元素
func (list *LinkedList) Insert(index int, value interface{}) bool {
	if 0 == index {
		list.head = &node{list.head, value}
	} else {
		if list.outOfRange(index - 1) {
			return false
		}
		n := list.findNode(index - 1)
		n.next = &node{n.next, value}
	}
	list.length++
	return true
}

// InsertToHead 在头部插入一个元素
func (list *LinkedList) InsertToHead(value interface{}) bool {
	return list.Insert(0, value)
}

// InsertToTail 在尾部插入一个元素
func (list *LinkedList) InsertToTail(value interface{}) bool {
	return list.Insert(list.length, value)
}

func (list *LinkedList) findNode(index int) *node {
	if list.outOfRange(index) {
		return nil
	}
	cur := list.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur
}

// Find 查找指定位置的元素
func (list *LinkedList) Find(index int) interface{} {
	node := list.findNode(index)
	if nil == node {
		return nil
	}
	return node.value

}

// Delete 删除指定位置的元素
func (list *LinkedList) Delete(index int) bool {
	if 0 == index {
		list.head = list.head.next
	} else {
		if list.outOfRange(index - 1) {
			return false
		}
		node := list.findNode(index - 1)
		node.next = node.next.next
	}
	list.length--
	return true
}

func (list *LinkedList) String() string {
	str := "head"
	if nil != list.head {
		for cur := list.head; cur != nil; cur = cur.next {
			str += fmt.Sprintf(" --> %+v", cur.value)
		}

	}
	return str + " --> tail"
}
