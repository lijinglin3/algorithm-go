package leetcode

var cache = make(map[int]bool)
var max int

func isPerfectSquare(num int) bool {
	m := max * max
	if num <= m {
		return cache[num]
	}

	for i := max + 1; ; i++ {
		tmp := i * i
		cache[tmp] = true
		max = i
		if tmp == num {
			return true
		}
		if tmp < num {
			continue
		}
		if tmp > num {
			return false
		}
	}
}
