package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left > right {
		return left + 1
	}
	return right + 1

}

func maxDepthByBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		depth++
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return depth
}

func maxDepthByDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type myNode struct {
		node  *TreeNode
		depth int
	}

	stack := []*myNode{{node: root, depth: 1}}
	maxDepth := 0

	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node := n.node
		depth := n.depth
		if depth > maxDepth {
			maxDepth = depth
		}

		if node.Right != nil {
			stack = append(stack, &myNode{node: node.Right, depth: depth + 1})
		}
		if node.Left != nil {
			stack = append(stack, &myNode{node: node.Left, depth: depth + 1})
		}
	}

	return maxDepth
}
