package leetcode

func romanToInt(s string) int {
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	sum := make([]int, len(s))
	for i := range s {
		if i != len(s)-1 {
			if m[string(s[i])] < m[string(s[i+1])] {
				sum = append(sum, -m[string(s[i])])
			} else {
				sum = append(sum, m[string(s[i])])
			}
		} else {
			sum = append(sum, m[string(s[i])])
		}
	}

	ss := 0
	for _, i := range sum {
		ss += i
	}

	return ss
}
