package arraystring

func merge(nums1 []int, m int, nums2 []int, n int) {

	p1, p2, p := m-1, n-1, m+n-1

	// Iterate while there are elements in nums2
	for p2 >= 0 {
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
}
