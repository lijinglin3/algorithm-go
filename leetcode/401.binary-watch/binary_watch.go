package leetcode

import (
	"fmt"
	"strings"
	"sync"
)

var (
	cache = make(map[int][]string)
	once  = sync.Once{}
)

func readBinaryWatch(num int) []string {
	once.Do(func() {
		for i := 0; i < 12; i++ {
			for j := 0; j < 60; j++ {
				b1 := fmt.Sprintf("%b", i)
				b2 := fmt.Sprintf("%b", j)
				sumOne := strings.Count(b1, "1") + strings.Count(b2, "1")
				cache[sumOne] = append(cache[sumOne], fmt.Sprintf("%d:%02d", i, j))
			}
		}
	})
	return cache[num]
}
