package leetcode

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if denominator == 0 {
		return "NAN"
	}
	var tmp string
	if numerator*denominator < 0 {
		tmp += "-"
	}
	numerator, denominator = abs(numerator), abs(denominator)
	tmp += strconv.Itoa(numerator / denominator)
	num := numerator % denominator
	if num == 0 {
		return tmp
	}
	tmp += "."
	hm := make(map[int]int)
	rptPos := -1
	for {
		num *= 10
		if pos, ok := hm[num]; ok {
			rptPos = pos
			break
		} else {
			hm[num] = len(tmp)
		}
		tmp += strconv.Itoa(num / denominator)
		num %= denominator
		if num == 0 {
			break
		}
	}
	if rptPos == -1 {
		return tmp
	}
	return fmt.Sprintf("%s(%s)", tmp[:rptPos], tmp[rptPos:])
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
