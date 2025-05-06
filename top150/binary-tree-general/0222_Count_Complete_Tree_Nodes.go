// Leetcode: https://leetcode.com/problems/count-complete-tree-nodes/
package binarytreegeneral

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := getLeftHeight(root)
	rightHeight := getRightHeight(root)

	if leftHeight == rightHeight {
		return (1 << leftHeight) - 1
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func getLeftHeight(node *TreeNode) int {
	height := 0
	for node != nil {
		height++
		node = node.Left
	}
	return height
}

func getRightHeight(node *TreeNode) int {
	height := 0
	for node != nil {
		height++
		node = node.Right
	}
	return height
}
