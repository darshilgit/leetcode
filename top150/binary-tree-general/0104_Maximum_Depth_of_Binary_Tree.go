// Leetcode: https://leetcode.com/problems/maximum-depth-of-binary-tree/
package binarytreegeneral

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type StackNode struct {
	node  *TreeNode
	depth int
}

func maxDepth(root *TreeNode) int {
	var currDepth int
	currDepth = 0
	if root != nil {
		stack := []StackNode{StackNode{node: root, depth: 0}}
		for len(stack) > 0 {
			//get node
			node := stack[0]

			//pop node
			stack = stack[1:]

			//depth check
			if node.depth+1 > currDepth {
				currDepth = node.depth + 1
			}

			//process node
			if node.node.Left != nil || node.node.Right != nil {
				if node.node.Left != nil {
					stack = append(stack, StackNode{node: node.node.Left, depth: node.depth + 1})
					// if node.depth+1 > currDepth{
					//     currDepth = node.depth+1
					// }
				}
				if node.node.Right != nil {
					stack = append(stack, StackNode{node: node.node.Right, depth: node.depth + 1})
					// if node.depth+1 > currDepth{
					//     currDepth = node.depth+1
					// }
				}
			}
		}
	} else {
		return currDepth
	}

	return currDepth
}
