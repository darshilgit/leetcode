import heapq
# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def mergeKLists(self, lists):
        if not lists:
            return None

        head = ListNode(0)
        curr = head
        priorityQ = []
        for item in lists:
            if item:
                # print(item.val)
                heapq.heappush(priorityQ, (item.val, item))
        # print(priorityQ)
        while priorityQ:
            pop = heapq.heappop(priorityQ)
            curr.next = ListNode(pop[0])
            curr = curr.next
            if pop[1].next:
                heapq.heappush(priorityQ, (pop[1].next.val, pop[1].next))
        return head.next
