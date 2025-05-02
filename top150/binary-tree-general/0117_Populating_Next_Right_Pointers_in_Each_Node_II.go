// Leetcode: https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/

package binarytreegeneral

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	var buildNext func(node *Node)
	buildNext = func(node *Node) {
		if node == nil {
			return
		}
		nextNode := node.Next
		for nextNode != nil {
			if nextNode.Left != nil {
				nextNode = nextNode.Left
				break
			}
			if nextNode.Right != nil {
				nextNode = nextNode.Right
				break
			}
			nextNode = nextNode.Next
		}
		if node.Right != nil {
			node.Right.Next = nextNode
		}

		if node.Left != nil {
			if node.Right != nil {
				node.Left.Next = node.Right
			} else {
				node.Left.Next = nextNode
			}
		}
		buildNext(node.Right)
		buildNext(node.Left)
	}
	buildNext(root)
	return root
}
