# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
import sys


class Solution:
    def isValidBST_util(self, root, _min, _max):
        if root == None:
            return True
        if root.val < _min or root.val > _max:
            return False
        return self.isValidBST_util(root.left, _min, root.val - 1) and self.isValidBST_util(root.right, root.val + 1,
                                                                                            _max)

    def isValidBST(self, root):
        int_max = sys.maxsize
        int_min = -sys.maxsize - 1

        return self.isValidBST_util(root, int_min, int_max)

