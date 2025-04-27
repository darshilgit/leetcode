package arraystring

import "runtime/debug"

func majorityElement(nums []int) int {
	count := 1
	wc := nums[0] //can or cannot be majority
	for i := 1; i < len(nums); i++ {
		if nums[i] == wc {
			count++
		} else {
			count--
		}
		if count == 0 {
			wc = nums[i]
			count = 1
		}
	}
	return wc
}

// found in LeetCode discussion
func init() {
	debug.SetMemoryLimit(10)
}
