package leetcode

func twoSum(numbers []int, target int) []int {
	length := len(numbers)
	p1, p2 := 0, 1
	for {
		if p2 >= length {
			p1, p2 = p1+1, p1+2
		} else {
			if numbers[p1]+numbers[p2] < target {
				p2++
			} else if numbers[p1]+numbers[p2] > target {
				p1, p2 = p1+1, p1+2
			} else {
				return []int{p1 + 1, p2 + 1}
			}
		}
	}
}
