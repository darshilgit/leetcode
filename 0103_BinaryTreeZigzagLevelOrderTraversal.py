# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def zigzagLevelOrder(self, root):
        answer = []
        if root == None:
            return answer
        q = [root]
        flag= True
        while len(q) != 0:
            a =[ele.val for ele in q]
            if flag:
                answer.append(a)
            else:
                answer.append(a[::-1])
            tempq = []
            for ele in q:
                if ele.left:
                    tempq.append(ele.left)
                if ele.right:
                    tempq.append(ele.right)
            q = tempq
            if flag:
                flag = False
            else:
                flag = True
        return answer