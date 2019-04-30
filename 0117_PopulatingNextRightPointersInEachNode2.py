"""
# Definition for a Node.
class Node:
    def __init__(self, val, left, right, next):
        self.val = val
        self.left = left
        self.right = right
        self.next = next
"""


class Solution:
    def connect(self, root):
        if root:
            queue = [root]
            while queue:
                tmpq = []
                for node in queue:
                    if node.left:
                        if tmpq: tmpq[-1].next = node.left
                        tmpq.append(node.left)
                    if node.right:
                        if tmpq: tmpq[-1].next = node.right
                        tmpq.append(node.right)

                queue = tmpq
        return root