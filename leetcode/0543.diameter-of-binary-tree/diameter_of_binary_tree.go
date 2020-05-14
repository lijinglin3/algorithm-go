package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

var result int

func diameterOfBinaryTree(root *TreeNode) int {
	result = 0
	depth(root)
	return result
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := depth(node.Left)
	right := depth(node.Right)
	if left+right > result {
		result = left + right
	}
	if left > right {
		return left + 1
	}
	return right + 1

}
