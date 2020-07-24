package leetcode

func findTheDifference(s string, t string) byte {
	sum := uint8(0)
	for i := 0; i < len(s); i++ {
		sum ^= s[i]
	}
	for i := 0; i < len(t); i++ {
		sum ^= t[i]
	}
	return sum
}
