package leetcode

func CountPrimes(n int) int {
	notPrim := make([]bool, n)
	for i := 2; i*i < n; i++ {
		if !notPrim[i] {
			for j := i * i; j < n; j += i {
				notPrim[j] = true
			}
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if !notPrim[i] {
			count++
		}
	}
	return count
}
