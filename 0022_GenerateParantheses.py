class Solution:
    def generateParenthesis(self, n):
        answer = []

        self.generateParenthesisHelper(n, n, '', answer)
        return answer

    def generateParenthesisHelper(self, leftP, rightP, currString, answer):
        if leftP == 0 and rightP == 0:
            answer.append(currString)
            return
        if leftP > 0:
            self.generateParenthesisHelper(leftP - 1, rightP, currString + '(', answer)

        if leftP < rightP:
            self.generateParenthesisHelper(leftP, rightP - 1, currString + ')', answer)


