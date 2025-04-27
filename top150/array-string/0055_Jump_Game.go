package arraystring

func canJump(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return true
	}
	target := len(nums) - 1
	for i := n - 2; i >= 0; i-- {
		if i+nums[i] >= target {
			target = i
		}
	}
	if target == 0 {
		return true
	}
	return false
}
