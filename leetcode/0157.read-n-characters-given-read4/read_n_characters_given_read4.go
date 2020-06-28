package leetcode

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	return func(buf []byte, n int) int {
		buff := make([]byte, 4)
		ret := 0
		for ret < n {
			count := read4(buff)
			if count <= 0 {
				break
			}

			for i := 0; i < count && ret < n; i++ {
				buf[ret] = buff[i]
				ret++
			}
		}
		return ret
	}
}
