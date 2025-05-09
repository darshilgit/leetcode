// LeetCode: https://leetcode.com/problems/binary-tree-right-side-view/
package binarytreebfs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var dummy *TreeNode

	stack := []*TreeNode{}

	stack = append(stack, root)
	stack = append(stack, dummy)

	var prev *TreeNode

	for len(stack) > 0 {
		node := stack[0]
		if node == dummy && len(stack) == 1 {
			res = append(res, prev.Val)
			break
		}

		if node == dummy {
			res = append(res, prev.Val)
			if len(stack) > 0 {
				stack = append(stack, dummy)
			}
		}

		stack = stack[1:]

		if node != dummy {
			prev = node
			//process the node
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}

	return res
}
