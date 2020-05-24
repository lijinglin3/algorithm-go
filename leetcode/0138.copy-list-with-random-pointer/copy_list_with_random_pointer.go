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

	list := &Node{Val: head.Val}
	n, l := head, list
	tmp1, tmp2 := map[*Node]int{n: 0}, map[int]*Node{0: list}
	index := 1

	for n != nil {
		if n.Next != nil {
			if i, ok := tmp1[n.Next]; ok {
				l.Next = tmp2[i]
			} else {
				l.Next = &Node{Val: n.Next.Val}
				tmp1[n.Next], tmp2[index] = index, l.Next
				index++
			}
		}

		if n.Random != nil {
			if i, ok := tmp1[n.Random]; ok {
				l.Random = tmp2[i]
			} else {
				l.Random = &Node{Val: n.Random.Val}
				tmp1[n.Random], tmp2[index] = index, l.Random
				index++
			}
		}

		l, n = l.Next, n.Next
	}

	return list
}
