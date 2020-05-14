package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode/common"

func minDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}

	depth := 0
	depthMap := map[*TreeNode]int{root: 1}
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		n := queue[0]
		queue = queue[1:]
		depth = depthMap[n]
		if n.Left == nil && n.Right == nil {
			break
		}
		if n.Left != nil {
			queue = append(queue, n.Left)
			depthMap[n.Left] = depth + 1
		}
		if n.Right != nil {
			queue = append(queue, n.Right)
			depthMap[n.Right] = depth + 1
		}
	}
	return depth
}
