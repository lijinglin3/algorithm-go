package leetcode

func twoSum1(nums []int, target int) []int {
	for i := range nums {
		for j := range nums {
			if i != j && (nums[i]+nums[j]) == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	lookup := make(map[int]int)
	for i, v := range nums {
		j, ok := lookup[v]
		lookup[target-v] = i
		if ok {
			return []int{j, i}
		}
	}
	return nil
}
