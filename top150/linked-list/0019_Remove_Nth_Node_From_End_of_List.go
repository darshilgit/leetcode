package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	nminusone, start := head, head

	for n > 0 {
		start = start.Next
		n--
	}

	if start == nil {
		return head.Next
	}

	for start.Next != nil {
		start = start.Next
		nminusone = nminusone.Next
	}

	nminusone.Next = nminusone.Next.Next

	return head

}
