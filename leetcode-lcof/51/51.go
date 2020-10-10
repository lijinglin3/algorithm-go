package leetcode

func reversePairs(nums []int) int {
	var count int
	sortArray(nums, &count)
	return count
}

func sortArray(nums []int, cnt *int) []int {
	var lens = len(nums)
	if lens <= 1 {
		return nums
	}
	return merge(sortArray(nums[:lens/2], cnt), sortArray(nums[lens/2:], cnt), cnt)
}

func merge(Pre, Post []int, cnt *int) []int {
	var lenPre, lenPost = len(Pre), len(Post)
	var list []int
	var i, j int
	for i < lenPre && j < lenPost {
		if Pre[i] <= Post[j] {
			list = append(list, Pre[i])
			i++
		} else {
			list = append(list, Post[j])
			*cnt = *cnt + lenPre - i
			j++
		}
	}
	if i < lenPre {
		list = append(list, Pre[i:]...)
	}
	if j < lenPost {
		list = append(list, Post[j:]...)
	}
	return list
}
