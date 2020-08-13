package leetcode

import (
	"math"
	"strconv"
)

func toHex(num int) string {
	if num < 0 {
		num += math.MaxUint32 + 1
	}
	return strconv.FormatInt(int64(num), 16)
}
