package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	n := root
	for n != nil {
		switch {
		case (p.Val >= n.Val && q.Val <= n.Val) || (p.Val <= n.Val && q.Val >= n.Val):
			return n
		case p.Val > n.Val && q.Val > n.Val:
			n = n.Right
		case p.Val < n.Val && q.Val < n.Val:
			n = n.Left
		}
	}
	return nil
}
