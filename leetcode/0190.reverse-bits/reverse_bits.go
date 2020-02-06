package leetcode

func ReverseBits(num uint32) uint32 {
	binary := make([]uint32, 0)
	for num != 0 {
		binary = append(binary, num%2)
		num = num >> 1
	}

	length := len(binary)
	result := 0
	for i := 0; i < length; i++ {
		result += int(binary[i]) * 1 << (31 - i)
	}
	return uint32(result)
}
