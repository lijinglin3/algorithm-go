package leetcode

var pick int

func guess(i int) int {
	switch {
	case i > pick:
		return -1
	case i < pick:
		return 1
	default:
		return 0
	}
}

func guessNumber(n int) int {
	start := 1
	end := n

	for start <= end {
		mid := start + (end-start)/2
		switch guess(mid) {
		case -1:
			end = mid - 1
		case 1:
			start = mid + 1
		default:
			return mid
		}
	}

	return -1
}
