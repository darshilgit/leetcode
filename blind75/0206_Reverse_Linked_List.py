# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:
    def helper (self, prev, curr):
        if curr is None:
            return prev
        next_node = curr.next
        curr.next = prev
        new_head = self.helper(curr, next_node)
        return new_head

    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # Iterative
        # prev, curr = None, head

        # while curr:
        #     next_node = curr.next
        #     curr.next = prev
        #     prev = curr
        #     curr = next_node

        # return prev
        
        # Recursive
        return self.helper(None, head)
