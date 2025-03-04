class Solution:
    def searchMatrix(self, matrix, target):
        if (len(matrix) == 0 or len(matrix[0]) == 0 or matrix == None):
            return False
        row = 0
        col = len(matrix[0]) - 1

        while (row < len(matrix) and col >= 0):
            if (matrix[row][col] == target):
                return True
            if (matrix[row][col] < target):
                row+=1
            else:
                col-=1

        return False