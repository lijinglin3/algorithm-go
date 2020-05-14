package leetcode

import "strconv"

func fizzBuzz(n int) []string {
	result := make([]string, 0)
	for i := 1; i <= n; i++ {
		str := ""
		if i%3 == 0 {
			str += "Fizz"
		}
		if i%5 == 0 {
			str += "Buzz"
		}
		if str == "" {
			str = strconv.FormatInt(int64(i), 10)
		}
		result = append(result, str)
	}
	return result
}
