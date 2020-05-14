package leetcode

// MyLinkedList my linked list
type MyLinkedList struct {
	val  int
	next *MyLinkedList
}

// Constructor initialize your data structure here.
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

// Get get the value of the index-th node in the linked list. If the index is invalid, return -1.
func (list *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}

	cur := list
	for index >= 0 {
		index--
		cur = cur.next

		if cur == nil {
			return -1
		}
	}
	return cur.val
}

// AddAtHead add a node of value val before the first element of the linked list.
// After the insertion, the new node will be the first node of the linked list.
func (list *MyLinkedList) AddAtHead(val int) {
	list.next = &MyLinkedList{val: val, next: list.next}
}

// AddAtTail append a node of value val to the last element of the linked list.
func (list *MyLinkedList) AddAtTail(val int) {
	cur := list
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = &MyLinkedList{val: val}
}

// AddAtIndex add a node of value val before the index-th node in the linked list.
// If index equals to the length of linked list, the node will be appended to the end of linked list.
// If index is greater than the length, the node will not be inserted.
func (list *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		return
	}

	cur := list
	for index > 0 {
		index--
		cur = cur.next

		if cur == nil {
			return
		}
	}
	cur.next = &MyLinkedList{val: val, next: cur.next}
}

// DeleteAtIndex delete the index-th node in the linked list, if the index is valid.
func (list *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}

	cur := list
	for index > 0 {
		index--
		cur = cur.next

		if cur == nil {
			return
		}

	}
	if cur.next == nil {
		return
	}
	cur.next = cur.next.next
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
