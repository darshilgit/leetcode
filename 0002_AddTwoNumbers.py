# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        carry = 0

        if l1 is None:
            return l2
        if l2 is None:
            return l1

        temp = ListNode(0)
        curr = temp
        p = l1
        q = l2

        while p is not None or q is not None:
            x = p.val if (p is not None) else 0
            y = q.val if (q is not None) else 0

            total = carry + x + y
            carry = int(total / 10)

            curr.next = ListNode(int(total % 10))
            curr = curr.next

            if p is not None:
                p = p.next
            if q is not None:
                q = q.next

        if carry > 0:
            curr.next = ListNode(1)

        return temp.next