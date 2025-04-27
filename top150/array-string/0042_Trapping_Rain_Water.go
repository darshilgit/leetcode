package arraystring

func trap(height []int) int {

	if len(height) == 0 {
		return 0
	}

	total := 0
	l, r := 0, len(height)-1
	leftMax, rightMax := height[l], height[r]

	for l < r {
		if leftMax < rightMax {
			l += 1
			leftMax = max(leftMax, height[l])
			total += leftMax - height[l]
		} else {
			r -= 1
			rightMax = max(rightMax, height[r])
			total += rightMax - height[r]
		}
	}
	return total

}
