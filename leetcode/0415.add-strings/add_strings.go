package leetcode

func addStrings(num1 string, num2 string) string {
	nb1, nb2 := []byte(num1), []byte(num2)
	if len(nb1) < len(nb2) {
		nb1, nb2 = nb2, nb1
	}

	sum := byte(0)
	for i, j := len(nb1)-1, len(nb2)-1; i >= 0; i, sum = i-1, sum/10 {
		if j >= 0 {
			sum += nb2[j] - '0'
			j--
		}
		sum += nb1[i] - '0'
		nb1[i] = (sum % 10) + '0'
	}
	if sum != 0 {
		nb1 = append([]byte{'1'}, nb1...)
	}
	return string(nb1)
}
