# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def rightSideView(self, root):
        answer = []
        if root == None:
            return answer
        q = []
        q.append(root)

        while (len(q)):
            # number of nodes at current level
            n = len(q)

            # Traverse all nodes of current level
            for i in range(1, n + 1):
                temp = q[0]
                q.pop(0)

                # Print the right most element
                # at the level
                if (i == n):
                    answer.append(temp.val)

                    # Add left node to queue
                if (temp.left != None):
                    q.append(temp.left)

                    # Add right node to queue
                if (temp.right != None):
                    q.append(temp.right)
        return answer
