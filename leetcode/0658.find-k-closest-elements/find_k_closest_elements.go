package leetcode

func findClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-1
	m := -1
	for {
		m = l + (r-l)/2
		if arr[m] == x {
			break
		} else if arr[m] < x {
			l = m + 1
		} else {
			r = m
		}
		if l == r {
			m = l
			break
		}
	}

	div := abs(arr[m] - x)
	base := m
	if m > 0 && abs(arr[m-1]-x) <= div {
		div = abs(arr[m-1] - x)
		base = m - 1
	}

	left, right := base, base
	for right-left < k-1 {
		leftDiv, rightDiv := 20001, 20001
		if left > 0 {
			leftDiv = abs(arr[left-1] - x)
		}
		if right < len(arr)-1 {
			rightDiv = abs(arr[right+1] - x)
		}
		if leftDiv > rightDiv {
			right++
		} else {
			left--
		}
	}
	return arr[left : right+1]
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
