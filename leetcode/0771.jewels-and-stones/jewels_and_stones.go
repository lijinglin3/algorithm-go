package leetcode

func numJewelsInStones(J string, S string) int {
	tmp := make(map[rune]interface{})
	for _, v := range J {
		tmp[v] = struct{}{}
	}

	count := 0
	for _, v := range S {
		if _, ok := tmp[v]; ok {
			count++
		}
	}
	return count
}
