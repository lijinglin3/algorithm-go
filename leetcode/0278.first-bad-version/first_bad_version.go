package leetcode

var testdata []bool

func isBadVersion(version int) bool { return testdata[version-1] }

func firstBadVersion(n int) int {
	left := 0
	right := n

	for left < right {
		mid := (left + right) / 2
		if !isBadVersion(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
