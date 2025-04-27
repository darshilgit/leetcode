package intervals

import "sort"

func merge(intervals [][]int) [][]int {

	var ans [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	l, r := intervals[0][0], intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= r {
			if r <= intervals[i][1] {
				r = intervals[i][1]
			}
			continue
		} else {
			ans = append(ans, []int{l, r})
			l = intervals[i][0]
			r = intervals[i][1]
		}
	}

	ans = append(ans, []int{l, r})

	return ans
}
