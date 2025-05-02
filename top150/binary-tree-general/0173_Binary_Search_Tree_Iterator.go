package binarytreegeneral

// Leetcode: https://leetcode.com/problems/binary-search-tree-iterator/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	minstack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	bst := BSTIterator{
		minstack: []*TreeNode{},
	}

	curr := root
	for curr != nil {
		bst.minstack = append(bst.minstack, curr)
		curr = curr.Left
	}

	return bst
}

func (this *BSTIterator) Next() int {
	retNode := this.minstack[len(this.minstack)-1]
	this.minstack = this.minstack[:len(this.minstack)-1]
	if retNode.Right != nil {
		curr := retNode.Right
		for curr != nil {
			this.minstack = append(this.minstack, curr)
			curr = curr.Left
		}
	}

	return retNode.Val

}

func (this *BSTIterator) HasNext() bool {
	return len(this.minstack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
