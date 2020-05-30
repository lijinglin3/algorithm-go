package leetcode

func groupAnagrams(strs []string) [][]string {
	tmp := map[string][]string{}

	for _, v := range strs {
		sv := quickSort(v)
		tmp[sv] = append(tmp[sv], v)
	}
	result := make([][]string, 0, len(tmp))
	for _, v := range tmp {
		result = append(result, v)
	}
	return result
}

func quickSort(str string) string {
	tmp := []byte(str)
	separateSort(tmp, 0, len(tmp)-1)
	return string(tmp)
}

func separateSort(arr []byte, start, end int) {
	if start >= end {
		return
	}
	i := partition(arr, start, end)
	separateSort(arr, start, i-1)
	separateSort(arr, i+1, end)
}

func partition(arr []byte, start, end int) int {
	// 选取最后一位当对比数字
	pivot := arr[end]

	var i = start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if !(i == j) {
				// 交换位置
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}

	arr[i], arr[end] = arr[end], arr[i]

	return i
}
