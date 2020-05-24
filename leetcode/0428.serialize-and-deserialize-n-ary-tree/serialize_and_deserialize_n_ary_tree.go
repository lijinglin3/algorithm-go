package leetcode

import "encoding/json"

// Node ...
type Node struct {
	Val      int
	Children []*Node
}

// Codec ...
type Codec struct {
}

// Constructor ...
func Constructor() *Codec {
	return &Codec{}
}

func (c *Codec) serialize(root *Node) string {
	if root == nil {
		return ""
	}

	v, _ := json.Marshal(root)
	return string(v)
}

func (c *Codec) deserialize(data string) *Node {
	if data == "" {
		return nil
	}

	root := &Node{}
	_ = json.Unmarshal([]byte(data), root)
	return root
}
