package leetcode

func kthGrammar(N int, K int) int {
	if K == 1 {
		return 0
	}

	G := kthGrammar(N-1, (K+1)/2)
	if K%2 == 1 {
		return G
	}
	return 1 - G
}
