package leetcode

var m = map[int]int{0: 0, 1: 1}

func fib(n int) int {
	if _, ok := m[n]; !ok {
		m[n] = (fib(n-1) + fib(n-2)) % 1000000007
	}
	return m[n]
}
