package leetcode

var m = map[int]int{0: 0, 1: 1}

func fib(N int) int {
	if _, ok := m[N]; !ok {
		m[N] = (fib(N-1) + fib(N-2)) % 1000000007
	}
	return m[N]
}
