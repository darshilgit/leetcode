package arraystring

func rotate(nums []int, k int) {
	n := len(nums)
	if k%n == 0 {
		return
	}
	k %= n //keep remainder incase k>n

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, i int, j int) {
	//reverse the array from index i to j
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

// 1 2 3 4 5 6 7 k=0
// 7 1 2 3 4 5 6 k=1,8
// 6 7 1 2 3 4 5 k=2,9
// 5 6 7 1 2 3 4 k=3,10
// 4 5 6 7 1 2 3 k=4,11
// 3 4 5 6 7 1 2 k=5,12
// 2 3 4 5 6 7 1 k=6,13
// 1 2 3 4 5 6 7 k=7,14

// 7 6 5 4 3 2 1
// 6 7 1 2 3 4 5
