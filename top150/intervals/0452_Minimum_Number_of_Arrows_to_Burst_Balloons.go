package intervals

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i int, j int) bool {
		if points[i][0] <= points[j][0] {
			if points[i][0] == points[j][0] {
				if points[i][1] < points[j][1] {
					return true
				} else {
					return false
				}
			}
			return true
		}
		return false
	})
	// fmt.Printf("%v\n",points)

	r := points[0][1]
	var res int
	for i := 1; i < len(points); i++ {
		if points[i][0] <= r {
			r = min(r, points[i][1])
			continue
		} else {
			// fmt.Printf("i:%v %v\n",i,points[i][0])
			res++
			r = points[i][1]
		}
	}
	res++

	return res
}
