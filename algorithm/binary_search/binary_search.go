package binarysearch

// Search 二分查找
func Search(nums []int, v int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	low, high := 0, length-1
	for low <= high {
		mid := (low + high) / 2
		switch {
		case nums[mid] == v:
			return mid
		case nums[mid] > v:
			high = mid - 1
		case nums[mid] < v:
			low = mid + 1
		}
	}

	return -1
}

// SearchRecursive 使用递归方式的二分查找
func SearchRecursive(nums []int, v int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	return bs(nums, v, 0, length-1)
}

func bs(nums []int, v int, low, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2

	switch {
	case nums[mid] > v:
		return bs(nums, v, low, mid-1)
	case nums[mid] < v:
		return bs(nums, v, mid+1, high)
	default:
		return mid
	}
}

// SearchFirst 查找第一个等于给定值的元素
func SearchFirst(nums []int, v int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	low, high := 0, length-1
	for low <= high {
		mid := (low + high) >> 1
		switch {
		case nums[mid] == v:
			if mid == 0 || nums[mid-1] != v {
				return mid
			}
			high = mid - 1
		case nums[mid] > v:
			high = mid - 1
		case nums[mid] < v:
			low = mid + 1
		}
	}
	return -1
}

// SearchLast 查找最后一个值等于给定值的元素
func SearchLast(nums []int, v int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	low, high := 0, length-1
	for low <= high {
		mid := (low + high) >> 1

		switch {
		case nums[mid] == v:
			if mid == length-1 || nums[mid+1] != v {
				return mid
			}
			low = mid + 1
		case nums[mid] > v:
			high = mid - 1
		case nums[mid] < v:
			low = mid + 1
		}
	}

	return -1
}

// TODO 查找第一个大于等于给定值的元素
// TODO 查找最后一个小于等于给定值的元素
