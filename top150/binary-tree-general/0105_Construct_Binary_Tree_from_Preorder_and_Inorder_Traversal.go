// Leetcode: https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

package binarytreegeneral

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// preorder 3 9 20 10 11 15 7 12 13
// inorder 12 10 13 9 11 3 15 20 7

func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderIdxMap := make(map[int]int, len(inorder))

	for idx, val := range inorder {
		inorderIdxMap[val] = idx
	}
	preorderIdx := 0

	var helper func(left int, right int) *TreeNode
	helper = func(left int, right int) *TreeNode {
		//base case
		if left > right {
			return nil
		}

		// create node
		val := preorder[preorderIdx]
		preorderIdx++

		node := &TreeNode{Val: val}

		node.Left = helper(left, inorderIdxMap[val]-1)
		node.Right = helper(inorderIdxMap[val]+1, right)

		return node
	}

	return helper(0, len(inorder)-1)

}
