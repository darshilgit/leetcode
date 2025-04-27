package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	// Create a dummy node to handle edge cases easily
	dummy := &ListNode{Next: head}
	prev := dummy

	// Move `prev` to the node just before the left position
	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	revptr, nextptr := reverseList(prev.Next, right-left+1)
	prev.Next.Next = nextptr
	prev.Next = revptr

	return dummy.Next

}

func reverseList(head *ListNode, diff int) (*ListNode, *ListNode) {

	var prev, tmp *ListNode
	curr := head

	for curr != nil && diff > 0 {
		tmp = curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
		diff--
	}

	return prev, tmp

}
