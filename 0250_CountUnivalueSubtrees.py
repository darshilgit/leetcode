class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
class Solution:

    def countSingleRec(self, root, count):
        if root is None:
            return True

        left = self.countSingleRec(root.left, count)
        right = self.countSingleRec(root.right, count)

        if left == False or right == False:
            return False

        if root.left and root.val != root.left.val:
            return False

        if root.right and root.val != root.right.val:
            return False

        count[0] += 1
        return True

    def countUnivalSubtrees(self, root):
        count = [0]
        self.countSingleRec(root, count)
        return count[0]

root = TreeNode(5)
root.left = TreeNode(4)
root.right = TreeNode(5)
root.left.left = TreeNode(4)
root.left.right = TreeNode(4)
root.right.right = TreeNode(5)

a = Solution()
print(a.countUnivalSubtrees(root))