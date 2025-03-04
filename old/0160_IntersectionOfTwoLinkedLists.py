# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def getIntersectionNode(self, headA, headB):
        """
        :type head1, head1: ListNode
        :rtype: ListNode
        """
        if headA == None or headB == None:
            return None
        len_a = 0
        len_b = 0
        curr_a = headA
        curr_b = headB

        while curr_a.next != None:
            len_a += 1
            curr_a = curr_a.next
        while curr_b.next != None:
            len_b += 1
            curr_b = curr_b.next
        curr_a = headA
        curr_b = headB
        d = abs(len_a - len_b)
        if len_a > len_b:
            while d > 0:
                curr_a = curr_a.next
                d -= 1
        else:
            while d > 0:
                curr_b = curr_b.next
                d -= 1
        while curr_a != curr_b:
            curr_a = curr_a.next
            curr_b = curr_b.next
        return curr_a

