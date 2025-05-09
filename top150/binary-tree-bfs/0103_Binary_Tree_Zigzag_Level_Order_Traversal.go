// Leetcode: https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
package binarytreebfs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	zig := true
	for len(queue) > 0 {
		n := len(queue)
		sub_res := []int{}
		for i := 0; i < n; i++ {
			node := queue[i]
			sub_res = append(sub_res, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if !zig {
			reverse_arr(sub_res)
			zig = true
		} else {
			zig = false
		}
		res = append(res, sub_res)
		queue = queue[n:]
	}
	return res
}

func reverse_arr(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
