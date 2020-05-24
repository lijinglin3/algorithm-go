package leetcode

// Node ...
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Child != nil {
		next, child := root.Next, root.Child
		root.Next = root.Child
		root.Next.Prev = root
		root.Child = nil
		for child.Next != nil {
			child = child.Next
		}
		child.Next = next
		if next != nil {
			next.Prev = child
		}
	}
	flatten(root.Next)
	return root
}
