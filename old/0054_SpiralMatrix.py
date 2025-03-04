class Solution:
    def spiralOrder(self, matrix):
        if not matrix or not matrix[0]:
            return []
        ans = []
        m, n = len(matrix), len(matrix[0])
        upper_row, lower_row, left_colum, right_column = 0, m - 1, 0, n - 1
        while left_colum < right_column and upper_row < lower_row:
            ans.append([matrix[upper_row][j] for j in range(left_colum, right_column)])
            ans.extend([matrix[i][right_column] for i in range(upper_row, lower_row)])
            ans.extend([matrix[lower_row][j] for j in range(right_column, left_colum, -1)])
            ans.extend([matrix[i][left_colum] for i in range(lower_row, upper_row, -1)])
            upper_row, lower_row, left_colum, right_column = upper_row + 1, lower_row - 1, left_colum + 1, right_column - 1
        if left_colum == right_column:
            ans.extend([matrix[i][right_column] for i in range(upper_row, lower_row + 1)])
        elif upper_row == lower_row:
            ans.extend([matrix[upper_row][j] for j in range(left_colum, right_column + 1)])
        return ans