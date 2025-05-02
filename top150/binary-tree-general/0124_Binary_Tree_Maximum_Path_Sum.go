// Leetcode: https://leetcode.com/problems/binary-tree-maximum-path-sum/

package binarytreegeneral

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var global_sum int

func maxPathSum(root *TreeNode) int {
	global_sum = math.MinInt

	_ = findMaxPathSum(root)

	return global_sum
}

func findMaxPathSum(node *TreeNode) int {
	if node == nil {
		return 0
	}

	lps := findMaxPathSum(node.Left)
	rps := findMaxPathSum(node.Right)

	curr_sum := max(lps+node.Val, rps+node.Val, node.Val, lps+rps+node.Val)

	if curr_sum > global_sum {
		global_sum = curr_sum
	}

	return max(lps+node.Val, rps+node.Val, node.Val)

}

// max function for multiple arguments
func max(values ...int) int {
	maxVal := values[0]
	for _, v := range values[1:] {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}
