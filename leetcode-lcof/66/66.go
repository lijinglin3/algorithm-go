package leetcode

func constructArr(a []int) []int {
	length := len(a)
	if length == 0 {
		return a
	}

	b := make([]int, length)
	b[0] = 1
	sum := 1
	for i := 1; i < length; i++ {
		sum *= a[i-1]
		b[i] = sum
	}
	sum = 1
	for i := length - 2; i >= 0; i-- {
		sum *= a[i+1]
		b[i] *= sum
	}
	return b
}
