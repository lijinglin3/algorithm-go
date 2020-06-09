package leetcode

func addDigits(num int) int {
	for num >= 10 {
		num = num/10 + num%10
	}
	return num
}

func addDigits2(num int) int {
	return (num-1)%9 + 1
}
