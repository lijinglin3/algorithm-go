package leetcode

func findContinuousSequence(target int) [][]int {
	result := make([][]int, 0)
	i, j, sum := 1, 2, 3
	for i < j && i+j <= target {
		switch {
		case sum < target:
			j++
			sum += j
		case sum > target:
			sum -= i
			i++
		default:
			tmp := make([]int, 0, j-i+1)
			for t := i; t <= j; t++ {
				tmp = append(tmp, t)
			}
			result = append(result, tmp)
			j++
			sum += j
		}
	}
	return result
}
