// Leetcode: https://leetcode.com/problems/minimum-absolute-difference-in-bst/
package binarysearchtree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func getMinimumDifference(root *TreeNode) int {
	// Helper function to perform the in-order traversal and track minimum difference
	var prev *TreeNode
	var abs_diff = math.MaxInt64 // Initialize to a very large value

	var inOrderTraversal func(node *TreeNode)
	inOrderTraversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		// Traverse the left subtree
		inOrderTraversal(node.Left)

		// Update the minimum absolute difference if prev is not nil
		if prev != nil {
			curr_diff := absDiff(node.Val, prev.Val)
			if curr_diff < abs_diff {
				abs_diff = curr_diff
			}
		}

		// Update prev to the current node
		prev = node

		// Traverse the right subtree
		inOrderTraversal(node.Right)
	}

	// Start the in-order traversal
	inOrderTraversal(root)

	return abs_diff
}
