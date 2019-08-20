package tree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	bt := NewBinaryTree(0)
	fmt.Println(bt.root)

	bt.root.left = NewNode(1)
	bt.root.right = NewNode(2)
	bt.root.left.left = NewNode(3)
	bt.root.left.right = NewNode(4)
	bt.root.right.left = NewNode(5)
	bt.root.right.right = NewNode(6)
	bt.root.left.left.left = NewNode(7)
	bt.root.left.left.right = NewNode(8)
	bt.root.left.right.left = NewNode(9)
	bt.root.left.right.right = NewNode(10)
	bt.root.right.left.left = NewNode(11)
	bt.root.right.left.right = NewNode(12)
	bt.root.right.right.left = NewNode(13)
	bt.root.right.right.right = NewNode(14)

	inOrder := []interface{}{7, 3, 8, 1, 9, 4, 10, 0, 11, 5, 12, 2, 13, 6, 14}
	preOrder := []interface{}{0, 1, 3, 7, 8, 4, 9, 10, 2, 5, 11, 12, 6, 13, 14}
	postOrder := []interface{}{7, 8, 3, 9, 10, 4, 1, 11, 12, 5, 13, 14, 6, 2, 0}

	if !reflect.DeepEqual(bt.InOrderTraverse(), inOrder) {
		t.Fail()
	}
	if !reflect.DeepEqual(bt.PreOrderTraverse(), preOrder) {
		t.Fail()
	}
	if !reflect.DeepEqual(bt.PostOrderTraverse(), postOrder) {
		t.Fail()
	}
	if !reflect.DeepEqual(bt.InOrderTraverseByStack(), inOrder) {
		t.Fail()
	}
	if !reflect.DeepEqual(bt.PreOrderTraverseByStack(), preOrder) {
		t.Fail()
	}
	if !reflect.DeepEqual(bt.PostOrderTraverseByStack(), postOrder) {
		t.Fail()
	}
	if !reflect.DeepEqual(bt.PostOrderTraverseByTwoStack(), postOrder) {
		t.Fail()
	}
}

func TestBinarySearchTree(t *testing.T) {
	bst := NewBinarySearchTree(0, func(v, nodeV interface{}) int {
		return v.(int) - nodeV.(int)
	})

	for i := 0; i < 10; i++ {
		bst.Insert(i)
	}
	for i := 0; i > -10; i-- {
		bst.Insert(i)
	}

	if bst.Find(-15) {
		t.Fail()
	}
	if !bst.Find(0) {
		t.Fail()
	}
	if bst.Find(15) {
		t.Fail()
	}

	if bst.Max() != 9 {
		t.Fatal("failed get max")
	}
	if bst.Min() != -9 {
		t.Fatal("failed get min")
	}

	if bst.Delete(0) {
		t.Fatal("failed delete")
	}

	bst.root = nil
	if bst.Insert(0) {
		t.Fatal()
	}
}
