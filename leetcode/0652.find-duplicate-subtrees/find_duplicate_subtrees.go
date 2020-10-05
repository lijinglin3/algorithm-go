package leetcode

import (
	"strconv"

	. "github.com/lijinglin3/algorithm-go/leetcode"
)

var (
	count   map[string]int
	results []*TreeNode
)

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	count, results = make(map[string]int), make([]*TreeNode, 0)
	traverse(root)
	return results
}

func traverse(root *TreeNode) string {
	if root == nil {
		return ""
	}

	uid := strconv.FormatInt(int64(root.Val), 10) + "," + traverse(root.Left) + "," + traverse(root.Right)
	count[uid]++
	if count[uid] == 2 {
		results = append(results, root)
	}
	return uid
}
