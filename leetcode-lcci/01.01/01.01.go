package leetcode

const a = 'a'

func isUnique(astr string) bool {
	num := 0
	for _, v := range astr {
		move := 1 << (v - a)
		if num&move != 0 {
			return false
		}
		num = num | move
	}
	return true
}
