package binarytreegeneral

// Leetcode: https://leetcode.com/problems/path-sum/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return isSumPath(root, targetSum)
}

func isSumPath(node *TreeNode, sum int) bool {
	if node != nil {
		if node.Left == nil && node.Right == nil {
			if sum-node.Val == 0 {
				return true
			}
			return false
		}
		return isSumPath(node.Left, sum-node.Val) || isSumPath(node.Right, sum-node.Val)
	}

	return false
}
