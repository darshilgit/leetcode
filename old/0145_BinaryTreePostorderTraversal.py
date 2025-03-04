# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

# class Solution:

#     def _postorderTraversal(self, root, result):
#         if root:
#             self._postorderTraversal(root.left, result)
#             self._postorderTraversal(root.right, result)
#             result.append(root.val)
#     def postorderTraversal(self, root):
#         if root == None:
#             return []
#         result = []
#         self._postorderTraversal(root, result)
#         return result

class Solution():
    def postorderTraversal(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        if root is None:
            return []

        stack, result = [root, ], []
        while stack:
            root = stack.pop()
            result.append(root.val)
            if root.left is not None:
                stack.append(root.left)
            if root.right is not None:
                stack.append(root.right)

        return result[::-1]