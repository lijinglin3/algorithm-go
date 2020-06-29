package leetcode

// Node ...
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	res := &Node{Val: head.Val}
	tail := res
	mpH := map[*Node]int{head: 0}
	sliR := []*Node{tail}
	index := 1
	for head.Next != nil {
		head = head.Next
		tail.Next = &Node{Val: head.Val}
		tail = tail.Next
		sliR = append(sliR, tail)
		mpH[head] = index
		index++
	}

	for node, index := range mpH {
		if node.Random != nil {
			sliR[index].Random = sliR[mpH[node.Random]]
		}
	}

	return res
}
