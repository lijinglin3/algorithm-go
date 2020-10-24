package leetcode

func oneEditAway(first string, second string) bool {
	if first == second {
		return true
	}

	l1, l2 := len(first), len(second)
	if l1-l2 > 1 {
		return false
	}
	if l1 < l2 {
		return oneEditAway(second, first)
	}

	for i := range second {
		if first[i] != second[i] {
			if l1 == l2 {
				return first[i+1:] == second[i+1:]
			}
			return first[i+1:] == second[i:]
		}
	}
	return true
}
