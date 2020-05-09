package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode/common"

func Preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	results := []int{root.Val}

	for _, n := range root.Children {
		results = append(results, Preorder(n)...)
	}

	return results
}

func PreorderByStack(root *Node) []int {
	if root == nil {
		return nil
	}

	var results []int
	stack := []*Node{root}

	for len(stack) != 0 {
		last := len(stack) - 1
		n := stack[last]
		stack = stack[:last]
		results = append(results, n.Val)

		length := len(n.Children)
		if length != 0 {
			for i := length - 1; i >= 0; i-- {
				stack = append(stack, n.Children[i])
			}
		}
	}

	return results
}
