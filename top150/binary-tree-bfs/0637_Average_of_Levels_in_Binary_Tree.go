// Leetcode: https://leetcode.com/problems/average-of-levels-in-binary-tree/
package binarytreebfs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {

	var res []float64

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		curr_res := 0.0
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			curr_res += float64(node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, curr_res/float64(n))
	}

	return res
}
