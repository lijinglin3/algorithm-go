package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode/common"

func LevelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	queue := []*Node{root}
	results := [][]int{{root.Val}}

	for len(queue) != 0 {
		var tmpQueue []*Node
		var ret []int
		for _, n := range queue {
			for i := range n.Children {
				ret = append(ret, n.Children[i].Val)
				tmpQueue = append(tmpQueue, n.Children[i])
			}
		}
		queue = tmpQueue
		if len(ret) != 0 {
			results = append(results, ret)
		}
	}

	return results
}
