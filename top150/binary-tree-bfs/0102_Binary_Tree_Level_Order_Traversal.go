// Leetcode: https://leetcode.com/problems/binary-tree-level-order-traversal/
package binarytreebfs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

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
		res = append(res, sub_res)
		queue = queue[n:]
	}

	return res
}
