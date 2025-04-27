package arraystring

func removeDuplicates(nums []int) int {
	k := 1
	dup := nums[0]
	for i, num := range nums {
		if i > 0 {
			if num != dup {
				nums[k] = num
				dup = num
				k++
			}
		}
	}
	return k
}
