package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Edge cases when one or both lists are nil
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// Initialize the merged list
	head := &ListNode{}
	current := head

	// Traverse both lists and merge them
	for list1 != nil && list2 != nil {
		// Compare values and add the smaller one to the merged list
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// At this point, one of the lists is nil, append the rest of the other list
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	// Return the merged list, starting from the node after the dummy head
	return head.Next
}
