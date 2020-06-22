package leetcode

import (
	"strconv"
)

func getFolderNames(names []string) []string {
	tmp := make(map[string]int64)
	result := make([]string, len(names))
	for i := 0; i < len(names); i++ {
		if v, ok := tmp[names[i]]; ok {
			count := v
			for {
				name := names[i] + "(" + strconv.FormatInt(count, 10) + ")"
				if _, exist := tmp[name]; exist {
					count++
					continue
				}
				result[i] = name
				tmp[name] = 1
				tmp[names[i]] = tmp[names[i]] + 1
				break
			}
		} else {
			result[i] = names[i]
			tmp[names[i]] = 1
		}
	}
	return result
}
