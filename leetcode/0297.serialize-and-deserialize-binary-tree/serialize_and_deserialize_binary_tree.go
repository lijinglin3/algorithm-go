package leetcode

import (
	"encoding/json"

	. "github.com/lijinglin3/algorithm-go/leetcode"
)

// Codec ...
type Codec struct {
}

// Constructor ...
func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	result := make([]interface{}, 0)
	queue := []*TreeNode{root}
	count := 0
	for len(queue) != 0 {
		nq := make([]*TreeNode, 0)
		for i := 0; i < len(queue); i++ {
			if queue[i] == nil {
				count++
				result = append(result, nil)
			} else {
				count = 0
				result = append(result, queue[i].Val)
				nq = append(nq, queue[i].Left, queue[i].Right)
			}
		}
		queue = nq
	}
	result = result[:len(result)-count]
	buf, _ := json.Marshal(result)
	return string(buf)
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	list := make([]interface{}, 0)
	_ = json.Unmarshal([]byte(data), &list)

	if len(list) == 0 {
		return nil
	}

	root := newTreeNode(list[0])
	queue := []*TreeNode{root}
	index, length := 1, len(list)
	for len(queue) != 0 {
		newQueue := make([]*TreeNode, 0)
		for _, n := range queue {
			if index < length {
				left := newTreeNode(list[index])
				if left != nil {
					n.Left = left
					newQueue = append(newQueue, left)
				}
			}
			index++

			if index < length {
				right := newTreeNode(list[index])
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

func newTreeNode(data interface{}) *TreeNode {
	switch t := data.(type) {
	case float64:
		return &TreeNode{Val: int(t)}
	default:
		return nil
	}
}
