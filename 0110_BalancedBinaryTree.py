# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def height(self, root):
        if root is None:
            return 0
        left = self.height(root.left) + 1
        right = self.height(root.right) + 1

        return max(left, right)

    def isBalanced(self, root):
        if root is None:
            return True

        lh = self.height(root.left)
        rh = self.height(root.right)

        if (abs(lh - rh) <= 1) and self.isBalanced(root.left) and self.isBalanced(root.right):
            return True

        return False