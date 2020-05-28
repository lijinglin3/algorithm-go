package leetcode

func isHappy(n int) bool {
	exist := map[int]struct{}{n: {}}

	for n != 1 {
		tmp := 0
		for n != 0 {
			tmp += (n % 10) * (n % 10)
			n = n / 10
		}
		if _, ok := exist[tmp]; ok {
			return false
		}
		n = tmp
		exist[n] = struct{}{}
	}
	return true
}
