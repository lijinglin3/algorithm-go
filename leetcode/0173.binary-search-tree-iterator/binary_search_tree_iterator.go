package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode/common"

type BSTIterator struct {
	stack  []*TreeNode
	length int
}

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

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	n := this.stack[this.length-1]
	this.stack = this.stack[:this.length-1]

	tmp := n.Right
	for tmp != nil {
		this.stack = append(this.stack, tmp)
		tmp = tmp.Left
	}
	this.length = len(this.stack)

	return n.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if this.length != 0 {
		return true
	}
	return false
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
