/*
 *      Author: peanutzhen
 *      Created time: 2021/8/23 15:25:46
 */

#include <bits/stdc++.h>

using namespace std;


class Solution {
public:
    vector<vector<int> > board;
    int n;
    int m;

    /**
     * Note: 类名、方法名、参数名已经指定，请勿修改
     *
     *
     *
     * @param grid int整型vector<vector<>>
     * @return int整型
     */
    int numberOfIslands(vector<vector<int> > &grid) {
        // write code here
        int ans = 0;
        board = grid;
        n = (int) board.size();
        if (n == 0)
            return 0;
        m = (int) board[0].size();
        for (int i = 0; i < n; ++i) {
            for (int j = 0; j < m; ++j) {
                if (board[i][j] == 1) {
                    dfs(i, j);
                    ++ans;
                }
            }
        }
        return ans;
    }

    void dfs(int i, int j) {
        if (i == -1 || j == -1 || i == n || j == m || board[i][j] == 0)
            return;
        board[i][j] = 0;
        dfs(i + 1, j);
        dfs(i - 1, j);
        dfs(i, j + 1);
        dfs(i, j - 1);
    }
};


int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    auto solution = new Solution();
    vector<vector<int>>
            grid = {
            {1, 1, 0, 0, 0},
            {1, 1, 0, 0, 0},
            {0, 0, 1, 0, 0},
            {0, 0, 0, 1, 1}
    };
    cout << solution->numberOfIslands(grid);
    return 0;
}

