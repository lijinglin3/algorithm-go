package leetcode

import "strconv"

// ValidWordAbbr ...
type ValidWordAbbr struct {
	data map[string]int
}

// Constructor ...
func Constructor(dictionary []string) ValidWordAbbr {
	abbr := ValidWordAbbr{data: make(map[string]int)}
	for _, v := range dictionary {
		if len(v) <= 2 {
			continue
		}
		abbr.data[v]++
		abbr.data[short(v)]++
	}
	return abbr
}

// IsUnique ...
func (a *ValidWordAbbr) IsUnique(word string) bool {
	if len(word) <= 2 || a.data[short(word)] == 0 {
		return true
	}
	if a.data[short(word)] == 1 && a.data[word] == 1 {
		return true
	}
	return false
}

func short(str string) string {
	return string(str[0]) + strconv.FormatInt(int64(len(str)-2), 10) + string(str[len(str)-1])
}
