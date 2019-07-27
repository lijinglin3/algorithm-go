package sort

// 选择排序，a表示数组
func SelectionSort(a []int) {
	length := len(a)
	if length <= 1 {
		return
	}
	for i := 0; i < length; i++ {
		// 查找最小值
		minIndex := i
		for j := i + 1; j < length; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		// 交换
		a[i], a[minIndex] = a[minIndex], a[i]

	}
}
