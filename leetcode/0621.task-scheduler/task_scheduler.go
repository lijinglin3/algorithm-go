package leetcode

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	charSlice := [26]int{}
	max := 0
	count := 0
	for i := 0; i < len(tasks); i++ {
		charSlice[tasks[i]-'A']++
		if max < charSlice[tasks[i]-'A'] {
			max = charSlice[tasks[i]-'A']
			count = 1
		} else if charSlice[tasks[i]-'A'] == max {
			count++
		}
	}
	if (max-1)*(n+1)+count < len(tasks) {
		return len(tasks)
	}
	return (max-1)*(n+1) + count
}
