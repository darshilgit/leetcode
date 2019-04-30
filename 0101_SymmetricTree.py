# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def isSymmetric(self, root):
        if root == None:
            return True
        return self.sym(root.left, root.right)

    def sym(self, left, right):
        if left == None or right == None:
            return left == right

        if left.val != right.val:
            return False

        return self.sym(left.left, right.right) and self.sym(left.right, right.left)
