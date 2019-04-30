# Definition for a Node.
class Node:
    def __init__(self, val, next, random):
        self.val = val
        self.next = next
        self.random = random

class Solution:
    def copyRandomList(self, head):
        hashmap = {}
        if head == None:
            return None
        temp = head
        while temp is not None:
            hashmap[temp] = Node(temp.val, None, None)
            temp = temp.next
        temp = head
        while temp is not None:
            hashmap[temp].next = hashmap.get(temp.next)
            hashmap[temp].random = hashmap.get(temp.random)
            temp = temp.next
        return hashmap.get(head)