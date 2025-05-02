// Leetcode: https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

package binarytreegeneral

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	inorderIdxMap := make(map[int]int, len(inorder))

	for idx, val := range inorder {
		inorderIdxMap[val] = idx
	}
	postorderIdx := len(postorder) - 1

	var helper func(left int, right int) *TreeNode
	helper = func(left int, right int) *TreeNode {
		//base case
		if left > right || len(postorder) == 0 {
			return nil
		}

		// create node
		val := postorder[postorderIdx]
		postorderIdx--

		node := &TreeNode{Val: val}

		node.Right = helper(inorderIdxMap[val]+1, right)
		node.Left = helper(left, inorderIdxMap[val]-1)

		return node
	}
	return helper(0, len(inorder)-1)
}
