package arraystring

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	product := 1
	for idx := range nums {
		product *= nums[idx]
		answer[idx] = product
	}

	product = 1
	for idx := len(nums) - 1; idx > 0; idx-- {
		product *= nums[idx]
		nums[idx] = product
	}

	lpro, rpro := 1, 1
	for idx := 0; idx < len(nums); idx++ {
		if idx == 0 {
			lpro = 1
		} else {
			lpro = answer[idx-1]
		}
		if idx == len(nums)-1 {
			rpro = 1
		} else {
			rpro = nums[idx+1]
		}
		nums[idx] = lpro * rpro
	}

	return nums
}
