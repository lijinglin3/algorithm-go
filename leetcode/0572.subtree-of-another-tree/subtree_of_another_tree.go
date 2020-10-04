package leetcode

import (
	"fmt"
	"strings"

	. "github.com/lijinglin2019/algorithm-go/leetcode"
)

func isSubtree(s *TreeNode, t *TreeNode) bool {
	var s1 strings.Builder
	preOrder(s, &s1, "")
	var s2 strings.Builder
	preOrder(t, &s2, "")
	return strings.Contains(s1.String(), s2.String())
}
func preOrder(root *TreeNode, res *strings.Builder, tag string) {
	if root == nil {
		res.WriteString(tag)
		return
	}
	res.WriteString(fmt.Sprintf("(%d)", root.Val))
	preOrder(root.Left, res, "L")
	preOrder(root.Right, res, "R")
}
