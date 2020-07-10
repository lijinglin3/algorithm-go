package leetcode

func isPowerOfFour(num int) bool {
	return num > 0 && num&(num-1) == 0 && num&2863311530 == 0
}
