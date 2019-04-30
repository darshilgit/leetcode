class Solution:
    def generate(self, numRows):
        answer = []

        for i in range(numRows):
            row = [None for ele in range(i + 1)]
            row[0], row[-1] = 1, 1

            # leave out first and last ele as they have been already added
            for j in range(1, len(row) - 1):
                row[j] = answer[i - 1][j - 1] + answer[i - 1][j]

            answer.append(row)

        return answer
