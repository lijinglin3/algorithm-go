package leetcode

type MyLinkedList struct {
	val  int
	next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}

	cur := this
	for index >= 0 {
		index--
		cur = cur.next

		if cur == nil {
			return -1
		}
	}
	return cur.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.next = &MyLinkedList{val: val, next: this.next}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	cur := this
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = &MyLinkedList{val: val}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		return
	}

	cur := this
	for index > 0 {
		index--
		cur = cur.next

		if cur == nil {
			return
		}
	}
	cur.next = &MyLinkedList{val: val, next: cur.next}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}

	cur := this
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
