package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{}
	prev := dummy
	curr := head
	found := false
	for curr != nil {
		//move curr ptr forward until you see the same val
		for curr.Next != nil && curr.Next.Val == curr.Val {
			curr = curr.Next
			found = true
		}
		if found {
			curr = curr.Next
		}
		if curr != nil {
			// fmt.Printf("%v\n",curr.Val)
			if (curr.Next == nil) || (curr.Next != nil && curr.Val != curr.Next.Val) {
				prev.Next = curr
				prev = curr
			}
		} else {
			prev.Next = curr
			prev = curr
		}
		if !found {
			curr = curr.Next
			found = false
		}
	}

	return dummy.Next
}
