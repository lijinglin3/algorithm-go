package leetcode

import "strconv"

func groupStrings(strings []string) [][]string {
	tmp := make(map[string][]string)

	for _, v := range strings {
		h := hash(v)
		if _, ok := tmp[h]; ok {
			tmp[h] = append(tmp[h], v)
		} else {
			tmp[h] = []string{v}
		}
	}

	result := make([][]string, 0, len(tmp))
	for _, v := range tmp {
		result = append(result, v)
	}
	return result
}

func hash(str string) string {
	result := ""
	delta := int64(str[0])
	for _, v := range str {
		result += "#" + strconv.FormatInt((int64(v)-delta+26)%26, 10)
	}
	return result
}
