class Solution:
    def searchMatrix(self, matrix, target):
        i = 0
        m = len(matrix)
        while i < m and matrix[i][0] <= target:
            i = i + 1
        i = i - 1
        return target in matrix[i]

if __name__ == '__main__':
    matrix = [[1]]
    target = 1
    ans = Solution().searchMatrix(matrix, target)
    print(ans)