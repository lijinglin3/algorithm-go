package leetcode

func CanThreePartsEqualSum(A []int) bool {
	sum := 0
	for _, v := range A {
		sum += v
	}
	if sum%3 != 0 {
		return false
	}

	found := false
	value := sum / 3
	sum = 0
	for _, v := range A[:len(A)-1] {
		sum += v
		if sum == value {
			if found {
				return true
			}
			found = true
			sum = 0
		}
	}
	return false
}
