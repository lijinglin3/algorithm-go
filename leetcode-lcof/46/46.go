package leetcode

func translateNum(num int) int {
	if num < 10 {
		return 1
	}
	if num%100 < 26 && num%100 > 9 {
		return translateNum(num/10) + translateNum(num/100)
	}
	return translateNum(num / 10)
}
