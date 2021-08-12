class Solution:
    def isValidSudoku(self, board):
        # 行限制
        row_lim = [[0]*9 for _ in range(9)]
        # 列限制
        col_lim = [[0]*9 for _ in range(9)]
        # 区块限制
        block_lim = [[0]*9 for _ in range(9)]
        # 遍历并设置限制
        for row in range(9):
            for col in range(9):
                # 获取当前位置数字
                if board[row][col] != ".":
                    num = int(board[row][col])
                else:
                    continue
                # 检测是否冗余
                if row_lim[num-1][row] == 1:
                    return False
                if col_lim[num-1][col] == 1:
                    return False
                block_count = int((col+3)/3) + 3 * int(row/3)
                if block_lim[num-1][block_count-1] == 1:
                    return False
                # 设置限制
                row_lim[num-1][row] = 1
                col_lim[num-1][col] = 1
                block_lim[num-1][block_count-1] = 1
        return True

if __name__ == '__main__':
    board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
    main = Solution()
    print(main.isValidSudoku(board))