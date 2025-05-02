package binarytreegeneral

import "fmt"

// Leetcode: https://leetcode.com/problems/sum-root-to-leaf-numbers/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var globalsum int

func sumNumbers(root *TreeNode) int {
	globalsum = 0
	countSum(root, 0)
	return globalsum
}

func countSum(node *TreeNode, sum int) {
	if node != nil {
		//processing node
		sum *= 10
		sum += node.Val

		//found end of one path
		if node.Left == nil && node.Right == nil {
			// sum+=node.Val
			fmt.Printf("sum : %v node.val : %v gsum:%v\n", sum, node.Val, globalsum)
			globalsum += sum
		}

		//adding more recursive calls
		countSum(node.Left, sum)
		countSum(node.Right, sum)
	}
	return
}
