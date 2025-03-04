# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def isSametree(self, s, t):
        if not s and not t:
            return True

        if s and t and s.val == t.val:
            return self.isSametree(s.left, t.left) and self.isSametree(s.right, t.right)

        # one empty and one non empty tree
        return False

    def isSubtree(self, root: Optional[TreeNode], subRoot: Optional[TreeNode]) -> bool:
        if not subRoot:
            return True
        if not root:
            return False
        
        # check if tree and subtree are exactly same
        if self.isSametree(root, subRoot):
            return True
        
        # check if subtree
        return self.isSubtree(root.left, subRoot) or self.isSubtree(root.right, subRoot)
