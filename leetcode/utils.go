package leetcode

import "encoding/json"

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// ListNodeDecoder 将数组字符串转换成链表
func ListNodeDecoder(str string) *ListNode {
	list := make([]int, 0)
	if err := json.Unmarshal([]byte(str), &list); err != nil {
		panic(err)
	}

	length := len(list)
	if length == 0 {
		return nil
	}

	listNode := &ListNode{Val: list[0]}
	tmp := listNode
	for i := 1; i < length; i++ {
		tmp.Next = &ListNode{Val: list[i]}
		tmp = tmp.Next
	}
	return listNode
}

func (l *ListNode) Last() *ListNode {
	n := l
	for n != nil {
		if n.Next == nil {
			return n
		}
		n = n.Next
	}
	return nil
}

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

// Node N叉树节点
type Node struct {
	Val      int
	Children []*Node
}

var (
	// NodeExample1 测试用例1
	NodeExample1 = &Node{
		Val: 0,
		Children: []*Node{
			{Val: 1, Children: []*Node{{Val: 4}, {Val: 5}, {Val: 6}}},
			{Val: 2},
			{Val: 3},
		},
	}
)
