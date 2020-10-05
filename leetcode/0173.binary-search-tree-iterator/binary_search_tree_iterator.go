package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

// BSTIterator BSTIterator
type BSTIterator struct {
	stack  []*TreeNode
	length int
}

// Constructor initialize your BSTIterator
func Constructor(root *TreeNode) BSTIterator {
	if root == nil {
		return BSTIterator{}
	}

	stack := []*TreeNode{root}
	n := root
	for n.Left != nil {
		stack = append(stack, n.Left)
		n = n.Left
	}

	return BSTIterator{
		stack:  stack,
		length: len(stack),
	}
}

// Next return the next smallest number
func (iterator *BSTIterator) Next() int {
	n := iterator.stack[iterator.length-1]
	iterator.stack = iterator.stack[:iterator.length-1]

	tmp := n.Right
	for tmp != nil {
		iterator.stack = append(iterator.stack, tmp)
		tmp = tmp.Left
	}
	iterator.length = len(iterator.stack)

	return n.Val
}

// HasNext return whether we have a next smallest number
func (iterator *BSTIterator) HasNext() bool {
	return iterator.length != 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
