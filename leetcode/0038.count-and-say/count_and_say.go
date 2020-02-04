package leetcode

import "strconv"

func CountAndSay(n int) string {
	result := "1"
	for n > 1 {
		tag := result[0]
		count := 1
		tmp := ""
		for i := 1; i < len(result); i++ {
			if result[i] == tag {
				count++
			} else {
				tmp = tmp + strconv.FormatInt(int64(count), 10) + string(tag)
				count = 1
				tag = result[i]
			}
		}
		tmp = tmp + strconv.FormatInt(int64(count), 10) + string(tag)
		result = tmp
		n--
	}

	return result
}
