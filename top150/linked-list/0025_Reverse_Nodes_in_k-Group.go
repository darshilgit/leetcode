package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 || head == nil {
		return head
	}
	curr := head
	dummy := &ListNode{}
	firstk := false
	prevlast := dummy

	for {
		// fmt.Printf("%v\n",curr.Val)
		//getKthNode
		found, kthnode := checkForKNodes(curr, k)
		if !firstk {
			dummy.Next = kthnode
			firstk = true
		}
		// fmt.Printf("%v %v\n",found,kthnode.Val)
		if found {
			nextcurr := kthnode.Next
			//1 2 3
			first, last := reverseList(curr, k)
			prevlast.Next = first
			prevlast = last
			// fmt.Printf("%v %v\n",first.Val,last.Val)

			curr = nextcurr
		} else {
			prevlast.Next = curr
			break
		}
	}

	return dummy.Next
}

func reverseList(head *ListNode, diff int) (*ListNode, *ListNode) {

	var prev, tmp *ListNode
	curr := head
	start := head

	for curr != nil && diff > 0 {
		tmp = curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
		diff--
	}

	return prev, start

}

func checkForKNodes(head *ListNode, k int) (bool, *ListNode) {
	count := 0
	target := k
	for head != nil && target > 0 {
		count++
		if count == k {
			return true, head
		}
		target--
		head = head.Next
	}

	return false, nil
}
