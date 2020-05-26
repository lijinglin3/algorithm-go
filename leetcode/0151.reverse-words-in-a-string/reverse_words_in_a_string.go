package leetcode

import "strings"

func reverseWords(s string) string {
	result := make([]string, 0)
	tmp := ""
	for _, v := range s {
		if v == ' ' {
			if tmp != "" {
				result = append(result, tmp)
				tmp = ""
			}
		} else {
			tmp += string(v)
		}
	}
	if tmp != "" {
		result = append(result, tmp)
		tmp = ""
	}

	if len(result) == 0 {
		return ""
	}
	if len(result) == 1 {
		return result[0]
	}
	for i := len(result) - 1; i > 0; i-- {
		tmp += result[i] + " "
	}
	return tmp + result[0]
}

func reverseWords2(s string) string {
	list := strings.Split(s, " ")
	var res []string
	for i := len(list) - 1; i >= 0; i-- {
		if len(list[i]) > 0 {
			res = append(res, list[i])
		}
	}
	s = strings.Join(res, " ")
	return s
}
