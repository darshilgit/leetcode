package intervals

import "strconv"

func summaryRanges(nums []int) []string {
	l, r := 0, 0
	var res []string
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) && nums[i+1]-nums[i] == 1 {
			r++
		} else {
			if l != r {
				res = append(res, strconv.Itoa(nums[l])+"->"+strconv.Itoa(nums[r]))
				// res = append(res,fmt.Sprintf("%v->%v",nums[l],nums[r]))
			} else {
				res = append(res, strconv.Itoa(nums[l]))
				// res = append(res,fmt.Sprintf("%v",nums[l]))
			}
			//update l & r
			l = i + 1
			r = i + 1
		}
	}

	return res
}
