# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        # Step 1: Initialize two pointers
        left = head
        right = head
        previous = None
        
        # Step 2: Move the `right` pointer n steps ahead
        for _ in range(n):
            right = right.next
        
        # Step 3: Traverse the list, keeping track of the `previous` pointer
        while right is not None:
            previous = left
            left = left.next
            right = right.next
        
        # Step 4: Handle case where the head itself needs to be removed
        if previous is None:
            # The node to remove is the head
            return head.next
        
        # Step 5: Remove the nth node by adjusting the `previous` pointer's next
        previous.next = left.next
        
        # Return the modified list
        return head