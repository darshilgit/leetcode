// Leetcode: https://leetcode.com/problems/same-tree/

package binarytreegeneral

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if (p != nil && q == nil) || (q != nil && p == nil) {
		return false
	}

	return (p.Val == q.Val) && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
