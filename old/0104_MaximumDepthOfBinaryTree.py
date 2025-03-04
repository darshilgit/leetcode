# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def maxDepth(self, root):

        if root == None:
            return int(0)

        left = self.maxDepth(root.left) + 1
        right = self.maxDepth(root.right) + 1

        if left > right:
            return left
        else:
            return right