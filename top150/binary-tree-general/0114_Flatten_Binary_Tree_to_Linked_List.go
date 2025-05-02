// Leetcode: https://leetcode.com/problems/flatten-binary-tree-to-linked-list/

package binarytreegeneral

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	if root.Left != nil {
		r := rightMost(root.Left)
		r.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
}

func rightMost(node *TreeNode) *TreeNode {
	for node.Right != nil {
		node = node.Right
	}

	return node
}
