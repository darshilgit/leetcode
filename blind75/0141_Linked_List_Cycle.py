
# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        if head is None or head.next is None:
            return False
        
        single = head
        double = head

        while double and double.next is not None:
            double = double.next.next
            single = single.next

            if single == double:
                return True
        
        return False
