# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def levelOrder(self, root):
        answer = []
        if root == None:
            return answer
        q = [root]

        while len(q) != 0:
            answer.append([ele.val for ele in q])
            tempq = []
            for ele in q:
                if ele.left:
                    tempq.append(ele.left)
                if ele.right:
                    tempq.append(ele.right)
            q = tempq
        return answer
