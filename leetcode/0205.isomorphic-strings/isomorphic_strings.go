package leetcode

func isIsomorphic(s string, t string) bool {
	lens, lent := len(s), len(t)
	if lens != lent {
		return false
	}

	smap, tmap := make(map[uint8]int), make(map[uint8]int)
	for i := 0; i < lens; i++ {
		sv, oks := smap[s[i]]
		tv, okv := tmap[t[i]]
		if oks != okv {
			return false
		}
		if oks && okv && sv != tv {
			return false
		}
		if !oks && !okv {
			smap[s[i]], tmap[t[i]] = i, i
		}
	}
	return true
}
