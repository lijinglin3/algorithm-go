package leetcode

var dp = map[int]map[int]float64{
	1: {
		1: 1.0 / 6,
		2: 1.0 / 6,
		3: 1.0 / 6,
		4: 1.0 / 6,
		5: 1.0 / 6,
		6: 1.0 / 6,
	},
}

func twoSum(n int) []float64 {
	data := multi(n)
	result := make([]float64, 0, len(data))
	for i := n; i <= 6*n; i++ {
		result = append(result, data[i])
	}
	return result
}

func multi(n int) map[int]float64 {
	if v, ok := dp[n]; ok {
		return v
	}
	tmp := make(map[int]float64)
	if n&1 != 0 {
		for i1, v1 := range dp[1] {
			for i2, v2 := range multi(n - 1) {
				tmp[i1+i2] += v1 * v2
			}
		}
	} else {
		for i1, v1 := range multi(n / 2) {
			for i2, v2 := range multi(n / 2) {
				tmp[i1+i2] += v1 * v2
			}
		}
	}
	dp[n] = tmp
	return tmp
}
