/*
 *      Author: peanutzhen
 *      Created time: 2021/8/18 18:43:02
 */

#include <bits/stdc++.h>

using namespace std;

int n, m, k;
int step; // dfs应保留空闲连通点的个数
vector<vector<char>> maze;

void dfs(int i, int j) {
    if (i == -1 || j == -1 || i == n || j == m || step == 0) {
        return;
    }
    if (maze[i][j] == '.') {
        maze[i][j] = 'Y';
        step--;
        dfs(i + 1, j);
        dfs(i - 1, j);
        dfs(i, j + 1);
        dfs(i, j - 1);
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    cin >> n >> m >> k;
    int s = 0;  // 空地数量
    int ei, ej; // 保存一个入口处 从入口处dfs
    for (int i = 0; i < n; ++i) {
        maze.emplace_back(m);
        for (int j = 0; j < m; ++j) {
            while ((maze[i][j] = cin.get()) == '\n');
            if (maze[i][j] == '.') {
                s++;
                ei = i;
                ej = j;
            }
        }
    }
    step = s - k;
    dfs(ei, ej);
    for (int i = 0; i < n; ++i) {
        for (int j = 0; j < m; ++j) {
            if (maze[i][j] == 'Y')
                cout << '.';
            else if (maze[i][j] == '.')
                cout << 'X';
            else
                cout << maze[i][j];
        }
        cout << "\n";
    }
    return 0;
}

