package leetcode

func shortestDistance(words []string, word1 string, word2 string) int {
	result, i1, i2 := len(words), -1, -1
	for i, v := range words {
		if v == word1 {
			i1 = i
			if i2 != -1 {
				result = min(result, i1, i2)
			}
		}
		if v == word2 {
			i2 = i
			if i1 != -1 {
				result = min(result, i1, i2)
			}
		}
	}
	return result
}

func min(i, i1, i2 int) int {
	delta := 0
	if i1 > i2 {
		delta = i1 - i2
	} else {
		delta = i2 - i1
	}
	if i > delta {
		i = delta
	}
	return i
}
