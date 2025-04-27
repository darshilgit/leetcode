package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	curr := head
	n := lenOfLL(curr)
	if k%n == 0 {
		return curr
	}

	k %= n //keep remainder incase k>n
	firstnode, _ := reverseList(curr, n)
	// fmt.Printf("%v %v\n",firstnode.Val,lastnode.Val)
	//5 4 3 2 1
	kthnode := getKthNode(firstnode, k)
	kfirst, klast := reverseList(firstnode, k)
	// fmt.Printf("%v %v\n",kfirst.Val,klast.Val)
	//4 5 | 3 2 1
	kthnodefirst, _ := reverseList(kthnode, n-k)
	// fmt.Printf("%v %v\n",kthnodefirst.Val,kthnodelast.Val)
	//4 5 | 1 2 3

	klast.Next = kthnodefirst

	return kfirst
}

func getKthNode(head *ListNode, k int) *ListNode {
	for k > 0 {
		head = head.Next
		k--
	}
	return head
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

func lenOfLL(head *ListNode) int {
	count := 0
	for head != nil {
		count++
		head = head.Next
	}

	return count
}
