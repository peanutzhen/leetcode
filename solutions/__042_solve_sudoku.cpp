/*
 *      Author: peanutzhen
 *      Created time: 2021/8/18 16:15:52
 */

#include <bits/stdc++.h>

using namespace std;

bool row[9][9];
bool col[9][9];
bool block[9][9];
vector<pair<int, int>> spaces;

bool dfs(vector<vector<char>> &board, int pos) {
    if (pos == spaces.size())
        return true;
    int i = spaces[pos].first, j = spaces[pos].second;
    for (int num = 1; num <= 9; ++num) {
        if (row[i][num - 1] || col[j][num - 1] || block[(i / 3) + 3 * (j / 3)][num - 1])
            continue;
        board[i][j] = num + '0';
        row[i][num - 1] = col[j][num - 1] = block[(i / 3) + 3 * (j / 3)][num - 1] = true;
        if (dfs(board, pos + 1))
            return true;
        row[i][num - 1] = col[j][num - 1] = block[(i / 3) + 3 * (j / 3)][num - 1] = false;
    }
    return false;
}

void solveSudoku(vector<vector<char>> &board) {
    memset(row, false, sizeof(row));
    memset(col, false, sizeof(col));
    memset(block, false, sizeof(block));
    spaces.clear();

    int m = (int) board.size();
    int n = (int) board[0].size();
    for (int i = 0; i < m; ++i) {
        for (int j = 0; j < n; ++j) {
            if (board[i][j] == '.') {
                spaces.emplace_back(i, j);
            } else {
                int num = board[i][j] - '0';
                row[i][num - 1] = col[j][num - 1] = block[(i / 3) + 3 * (j / 3)][num - 1] = true;
            }
        }
    }
    dfs(board, 0);
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);


    return 0;
}

