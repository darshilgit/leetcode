// Leetcode: https://leetcode.com/problems/invert-binary-tree/

package binarytreegeneral

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	swapNodes(root)

	return root
}

func swapNodes(node *TreeNode) {
	if node != nil {
		temp := node.Left
		node.Left = node.Right
		node.Right = temp

		if node.Left != nil {
			swapNodes(node.Left)
		}
		if node.Right != nil {
			swapNodes(node.Right)
		}
	}
}
