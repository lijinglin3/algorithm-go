package leetcode

func isPathCrossing(path string) bool {
	positions, p := map[int]map[int]int{0: {0: 1}}, [2]int{0, 0}

	for i := 0; i < len(path); i++ {
		switch path[i] {
		case 'N':
			p[1]++
		case 'E':
			p[0]++
		case 'S':
			p[1]--
		case 'W':
			p[0]--
		}
		if _, ok := positions[p[0]]; ok {
			if positions[p[0]][p[1]] == 1 {
				return true
			}
		} else {
			positions[p[0]] = make(map[int]int)
		}
		positions[p[0]][p[1]]++
	}
	return false
}
