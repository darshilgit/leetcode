class Solution:
    def letterCombinations(self, digits):
        answer = []

        if digits == "":
            return answer
        matcher = ["0", "1", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"]

        self.letterCombinationsHelper(digits, answer, 0, "", matcher)
        return answer

    def letterCombinationsHelper(self, digits, answer, index, currString, matcher):
        if len(currString) == len(digits):
            answer.append(currString)
            return
        letters = matcher[int(digits[index])]
        for i in range(len(letters)):
            self.letterCombinationsHelper(digits, answer, index + 1, currString + letters[i], matcher)
