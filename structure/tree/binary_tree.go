package tree

import (
	"fmt"
)

// Node 二叉树的节点
type Node struct {
	data  interface{}
	left  *Node
	right *Node
}

// NewNode 初始化一个节点
func NewNode(data interface{}) *Node {
	return &Node{data: data}
}

func (n *Node) String() string {
	return fmt.Sprintf("v:%+v, left:%+v, right:%+v", n.data, n.left, n.right)
}

// BinaryTree 二叉树的结构体
type BinaryTree struct {
	root *Node
}

// NewBinaryTree 初始化一个二叉树
func NewBinaryTree(rootV interface{}) *BinaryTree {
	return &BinaryTree{root: NewNode(rootV)}
}

// InOrderTraverse 使用递归方法的中序遍历
func (bt *BinaryTree) InOrderTraverse() (result []interface{}) {
	if left := bt.root.left; nil != left {
		newBT := &BinaryTree{root: left}
		result = append(result, newBT.InOrderTraverse()...)
	}

	result = append(result, bt.root.data)

	if right := bt.root.right; nil != right {
		newBT := &BinaryTree{root: right}
		result = append(result, newBT.InOrderTraverse()...)
	}

	return
}

// PreOrderTraverse 使用递归方法的前序遍历
func (bt *BinaryTree) PreOrderTraverse() (result []interface{}) {
	result = append(result, bt.root.data)

	if left := bt.root.left; nil != left {
		newBT := &BinaryTree{root: left}
		result = append(result, newBT.PreOrderTraverse()...)
	}

	if right := bt.root.right; nil != right {
		newBT := &BinaryTree{root: right}
		result = append(result, newBT.PreOrderTraverse()...)
	}

	return
}

// PostOrderTraverse 使用递归方法的后序遍历
func (bt *BinaryTree) PostOrderTraverse() (result []interface{}) {
	if left := bt.root.left; nil != left {
		newBT := &BinaryTree{root: left}
		result = append(result, newBT.PostOrderTraverse()...)
	}

	if right := bt.root.right; nil != right {
		newBT := &BinaryTree{root: right}
		result = append(result, newBT.PostOrderTraverse()...)
	}

	result = append(result, bt.root.data)
	return
}

// InOrderTraverseByStack 使用stack的中序遍历
func (bt *BinaryTree) InOrderTraverseByStack() (result []interface{}) {
	stack := make([]*Node, 0)
	n := bt.root

	for len(stack) != 0 || nil != n {
		if nil != n {
			stack = append(stack, n)
			n = n.left
		} else {
			tmp := stack[len(stack)-1]
			result = append(result, tmp.data)
			n = tmp.right
			stack = stack[:len(stack)-1]
		}
	}

	return
}

// PreOrderTraverseByStack 使用stack的前序遍历
func (bt *BinaryTree) PreOrderTraverseByStack() (result []interface{}) {
	stack := make([]*Node, 0)
	n := bt.root
	for len(stack) != 0 || nil != n {
		if nil != n {
			result = append(result, n.data)
			stack = append(stack, n)
			n = n.left
		} else {
			n = stack[len(stack)-1].right
			stack = stack[:len(stack)-1]
		}
	}
	return
}

// PostOrderTraverseByStack 使用stack的后序遍历
func (bt *BinaryTree) PostOrderTraverseByStack() (result []interface{}) {
	var pre *Node
	n := bt.root
	stack := []*Node{n}
	for len(stack) != 0 {
		n := stack[len(stack)-1]
		if (n.left == nil && n.right == nil) || (pre != nil && (pre == n.left || pre == n.right)) {
			stack = stack[:len(stack)-1]
			result = append(result, n.data)
			pre = n
		} else {
			if nil != n.right {
				stack = append(stack, n.right)
			}
			if nil != n.left {
				stack = append(stack, n.left)
			}
		}
	}
	return
}

// PostOrderTraverseByTwoStack 使用两个stack的后序遍历
func (bt *BinaryTree) PostOrderTraverseByTwoStack() (result []interface{}) {
	n := bt.root
	stack1 := []*Node{n}
	stack2 := make([]*Node, 0)
	for len(stack1) != 0 {
		n = stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = append(stack2, n)
		if nil != n.left {
			stack1 = append(stack1, n.left)
		}
		if nil != n.right {
			stack1 = append(stack1, n.right)
		}
	}

	for i := len(stack2) - 1; i >= 0; i-- {
		result = append(result, stack2[i].data)
	}
	return
}

// BinarySearchTree 二叉搜索树
type BinarySearchTree struct {
	*BinaryTree
	CompareFunc func(v, nodeV interface{}) int
}

// NewBinarySearchTree 初始化二叉搜索树
func NewBinarySearchTree(rootV interface{}, compareFunc func(v, nodeV interface{}) int) *BinarySearchTree {
	return &BinarySearchTree{
		BinaryTree:  NewBinaryTree(rootV),
		CompareFunc: compareFunc,
	}
}

// Insert 插入一个值到二叉搜索树中
func (bst *BinarySearchTree) Insert(v interface{}) bool {
	n := bst.root
	for nil != n {
		compareResult := bst.CompareFunc(v, n.data)
		switch {
		case compareResult == 0:
			return false
		case compareResult < 0:
			if nil == n.left {
				n.left = NewNode(v)
				return true
			}
			n = n.left
		case compareResult > 0:
			if nil == n.right {
				n.right = NewNode(v)
				return true
			}
			n = n.right
		}
	}
	return false
}

// Delete 删除二叉搜索树中的某个值
func (bst *BinarySearchTree) Delete(v interface{}) bool {
	// TODO
	return false
}

// Find 查找二叉搜索树中的某个值
func (bst *BinarySearchTree) Find(v interface{}) bool {
	n := bst.root
	for nil != n {
		compareResult := bst.CompareFunc(v, n.data)
		switch {
		case compareResult == 0:
			return true
		case compareResult < 0:
			n = n.left
		case compareResult > 0:
			n = n.right
		}
	}

	return false
}

// Min 返回二叉搜索树中的最小值
func (bst *BinarySearchTree) Min() interface{} {
	n := bst.root
	for nil != n {
		if nil == n.left {
			break
		}
		n = n.left
	}
	return n.data
}

// Max 返回二叉搜索树中的最大值
func (bst *BinarySearchTree) Max() interface{} {
	n := bst.root
	for nil != n {
		if nil == n.right {
			break
		}
		n = n.right
	}
	return n.data
}
