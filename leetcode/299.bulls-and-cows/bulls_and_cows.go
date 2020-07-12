package leetcode

import "fmt"

func getHint(secret string, guess string) string {
	length := len(secret)
	A, B := 0, 0
	tmp := make(map[uint8]int)
	for i := 0; i < length; i++ {
		if secret[i] == guess[i] {
			A++
		} else {
			tmp[secret[i]]++
		}
	}
	for i := 0; i < length; i++ {
		if secret[i] != guess[i] && tmp[guess[i]] > 0 {
			B++
			tmp[guess[i]]--
		}
	}
	return fmt.Sprintf("%dA%dB", A, B)
}
