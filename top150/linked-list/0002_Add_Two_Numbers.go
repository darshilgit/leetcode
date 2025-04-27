package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	headl1 := l1
	prevl1 := l1
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			l1.Val += l2.Val + carry
			if l1.Val >= 10 {
				carry = l1.Val / 10
				l1.Val %= 10
			} else {
				carry = 0
			}
			prevl1 = l1
			l1 = l1.Next
			l2 = l2.Next
		} else {
			//Everything needs to go in l1
			//case 1 : only l1 is remaning
			//case 2 : only l2 is remaining
			if l1 != nil {
				l1.Val += carry
				if l1.Val >= 10 {
					carry = l1.Val / 10
					l1.Val %= 10
				} else {
					carry = 0
				}
				prevl1 = l1
				l1 = l1.Next
			} else if l2 != nil {
				val := l2.Val + carry
				if val >= 10 {
					carry = val / 10
					val %= 10
				} else {
					carry = 0
				}
				tempNode := &ListNode{
					Val:  val,
					Next: nil,
				}
				prevl1.Next = tempNode
				prevl1 = tempNode
				l2 = l2.Next
			}
		}
	}

	if carry > 0 {
		tempNode := &ListNode{
			Val:  carry,
			Next: nil,
		}
		prevl1.Next = tempNode
	}

	return headl1
}
