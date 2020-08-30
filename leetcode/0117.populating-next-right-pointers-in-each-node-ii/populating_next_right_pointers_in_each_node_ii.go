package leetcode

// Node ...
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := []*Node{root}
	for len(queue) != 0 {
		length := len(queue)
		for i, n := range queue {
			if i != length-1 {
				n.Next = queue[i+1]
			}
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
		queue = queue[length:]
	}
	return root
}
