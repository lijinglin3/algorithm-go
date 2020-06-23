package leetcode

import "sort"

func canAttendMeetings(intervals [][]int) bool {
	array := sortInterval(intervals)
	sort.Sort(array)

	pre := 0
	for _, v := range array {
		if v[0] < pre {
			return false
		}
		pre = v[1]
	}
	return true
}

type sortInterval [][]int

func (si sortInterval) Len() int { return len(si) }

func (si sortInterval) Less(i, j int) bool { return si[i][0] < si[j][0] }

func (si sortInterval) Swap(i, j int) { si[i], si[j] = si[j], si[i] }
