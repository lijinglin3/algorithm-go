package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	l := len(postorder) - 1
	rootVal := postorder[l]
	root := &TreeNode{Val: rootVal}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			break
		}
	}
	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:l])
	return root
}
