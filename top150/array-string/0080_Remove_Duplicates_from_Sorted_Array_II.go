package arraystring

func removeDuplicates2(nums []int) int {
	wptr, occ, num2Count := 0, 0, nums[0]

	for _, num := range nums {
		if num2Count != num {
			loop := min(2, occ)
			for i := 0; i < loop; i++ {
				nums[wptr] = num2Count
				wptr++
			}
			occ = 1
			num2Count = num
		} else {
			occ++
		}
	}

	loop := min(2, occ)
	for i := 0; i < loop; i++ {
		nums[wptr] = num2Count
		wptr++
	}

	return wptr
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
