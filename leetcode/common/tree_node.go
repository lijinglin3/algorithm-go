package common

import "encoding/json"

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode 初始化二叉树
func NewTreeNode(data interface{}) *TreeNode {
	switch t := data.(type) {
	case float64:
		return &TreeNode{Val: int(t)}
	default:
		return nil
	}
}

// TreeNodeDecoder 将数组字符串转换成二叉树
func TreeNodeDecoder(str string) *TreeNode {
	list := make([]interface{}, 0)
	if err := json.Unmarshal([]byte(str), &list); err != nil {
		panic(err)
	}

	if len(list) == 0 {
		return nil
	}

	root := NewTreeNode(list[0])
	queue := []*TreeNode{root}
	index, length := 1, len(list)
	for len(queue) != 0 {
		newQueue := make([]*TreeNode, 0)
		for _, n := range queue {
			if index < length {
				left := NewTreeNode(list[index])
				if left != nil {
					n.Left = left
					newQueue = append(newQueue, left)
				}
			}
			index++

			if index < length {
				right := NewTreeNode(list[index])
				if right != nil {
					n.Right = right
					newQueue = append(newQueue, right)
				}
			}
			index++
		}
		queue = newQueue
	}

	return root
}
