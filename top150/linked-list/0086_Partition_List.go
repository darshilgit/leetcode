package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {

	smallDummy := &ListNode{}
	smallPrev := smallDummy

	largeDummy := &ListNode{}
	largePrev := largeDummy

	for head != nil {
		if head.Val < x {
			smallPrev.Next = head
			smallPrev = head
		} else {
			largePrev.Next = head
			largePrev = head
		}
		head = head.Next
	}

	smallPrev.Next = largeDummy.Next
	largePrev.Next = nil

	return smallDummy.Next

}
