package leetcode

func countDigitOne(n int) int {
	if n < 1 {
		return 0
	}
	count, factor := 0, 1
	for n/factor != 0 {
		digit := (n / factor) % 10
		high := n / (10 * factor)
		if digit == 0 {
			count += high * factor
		} else if digit == 1 {
			count += high * factor
			count += (n % factor) + 1
		} else {
			count += (high + 1) * factor
		}
		factor = factor * 10
	}
	return count
}
