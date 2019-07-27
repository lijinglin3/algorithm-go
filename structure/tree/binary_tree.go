package tree

import (
	"fmt"
)

type Node struct {
	data  interface{}
	left  *Node
	right *Node
}

func NewNode(data interface{}) *Node {
	return &Node{data: data}
}

func (n *Node) String() string {
	return fmt.Sprintf("v:%+v, left:%+v, right:%+v", n.data, n.left, n.right)
}

type BinaryTree struct {
	root *Node
}

func NewBinaryTree(rootV interface{}) *BinaryTree {
	return &BinaryTree{root: NewNode(rootV)}
}

func (bt *BinaryTree) InOrderTraverse() (result []interface{}) {
	if left := bt.root.left; nil != left {
		result = append(result, left.data)
	}

	result = append(result, bt.root.data)

	if right := bt.root.right; nil != right {
		newBT := &BinaryTree{root: right}
		result = append(result, newBT.InOrderTraverse()...)
	}

	return
}

func (bt *BinaryTree) PreOrderTraverse() (result []interface{}) {
	//stack := make([]*Node, 0)
	//stack = append(stack, bt.root)
	//for len(stack) != 0 {
	//	n := stack[len(stack)-1]
	//	stack = stack[:len(stack)-1]
	//	result = append(result, n.data)
	//	if n.left != nil {
	//		stack = append(stack, n.left)
	//	}
	//	if n.right != nil {
	//		stack = append(stack, n.right)
	//	}
	//}
	//return

	result = append(result, bt.root.data)

	if left := bt.root.left; nil != left {
		newBT := &BinaryTree{root: left}
		newBT.PreOrderTraverse()
	}

	if right := bt.root.right; nil != right {
		newBT := &BinaryTree{root: right}
		newBT.PreOrderTraverse()
	}

	return
}

func (bt *BinaryTree) PostOrderTraverse() (result []interface{}) {
	if left := bt.root.left; nil != left {
		newBT := &BinaryTree{root: left}
		newBT.PostOrderTraverse()
	}

	if right := bt.root.right; nil != right {
		newBT := &BinaryTree{root: right}
		newBT.PostOrderTraverse()
	}

	result = append(result, bt.root.data)
	return
}

func (bt *BinaryTree) InOrderTraverseByStack() (result []interface{}) {
	stack := make([]*Node, 0)
	n := bt.root

	for len(stack) != 0 || nil != n {
		if nil != n {
			stack = append(stack)
			n = n.left
		} else {
			tmp := stack[len(stack)-1]
			result = append(result, tmp.data)
			n = tmp.right
		}
	}

	return
}

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
		}
	}
	return
}

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
			if nil != n.left {
				stack = append(stack, n.left)
			}
			if nil != n.right {
				stack = append(stack, n.right)
			}
		}
	}
	return
}

func (bt *BinaryTree) PostOrderTraverseByTwoStack() (result []interface{}) {
	n := bt.root
	stack1 := []*Node{n}
	stack2 := make([]*Node, 0)
	for len(stack1) != 0 {
		n = stack1[len(stack1)-1]
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

type BinarySearchTree struct {
	*BinaryTree
	CompareFunc func(v, nodeV interface{}) int
}

func NewBinarySearchTree(rootV interface{}, compareFunc func(v, nodeV interface{}) int) *BinarySearchTree {
	return &BinarySearchTree{
		BinaryTree:  NewBinaryTree(rootV),
		CompareFunc: compareFunc,
	}
}

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
	return true
}

func (bst *BinarySearchTree) Delete(v interface{}) bool {
	// TODO
	return false
}

func (bst *BinarySearchTree) Find(v interface{}) interface{} {
	n := bst.root
	for nil != n {
		compareResult := bst.CompareFunc(v, n.data)
		switch {
		case compareResult == 0:
			return n
		case compareResult < 0:
			n = n.left
		case compareResult > 0:
			n = n.right
		}
	}

	return nil
}

func (bst *BinarySearchTree) Min() interface{} {
	n := bst.root
	for nil != n {
		if nil == n.left {
			return n.data
		}
		n = n.left
	}
	return nil
}

func (bst *BinarySearchTree) Max() interface{} {
	n := bst.root
	for nil != n {
		if nil == n.right {
			return n.data
		}
		n = n.right
	}
	return nil
}
