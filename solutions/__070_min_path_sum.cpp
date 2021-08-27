/*
 *      Author: peanutzhen
 *      Created time: 2021/8/27 10:55:28
 */

#include <bits/stdc++.h>

using namespace std;

int n;
int m;
int dp[205][205];
vector<vector<int>> board;
vector<int> path;
vector<int> min_path;

// 带最优路径的dfs算法
// dp带最优路径还没想好怎么做比较好
void dfs(int i, int j, int walk) {
    if (i == n || j == m)
        return;
    walk += board[i][j];
    if (dp[i][j] != -1 && walk >= dp[i][j])
        return;
    dp[i][j] = walk;
    path.push_back(board[i][j]);
    if (i == n-1 && j == m-1)
        min_path = path;
    dfs(i+1, j, walk);
    dfs(i, j+1, walk);
    path.pop_back();
}

int minPathSum(vector<vector<int>>& grid) {
    n = (int) grid.size();
    if (n == 0)
        return 0;
    m = (int) grid[0].size();
    memset(dp, -1, sizeof(dp));
    board = grid;
    path.clear();
    dfs(0, 0, 0);
    return dp[n-1][m-1];
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    vector<vector<int>> grid = {{1,3,1},{1,5,1},{4,2,1}};
    cout << minPathSum(grid) << "\n";
    for (auto v: min_path)
        cout << v << " ";
    return 0;
}

