/*
 *      Author: peanutzhen
 *      Created time: 2021/9/12 12:33:50
 */

#include <bits/stdc++.h>

using namespace std;

const int direction[4][2] = {
        {-1, 0},
        {1,  0},
        {0,  -1},
        {0,  1}
};

struct state {
    int x;
    int y;
    int z; // 还能破墙(走1)的次数

    state(int _x, int _y, int _z) : x(_x), y(_y), z(_z) {}
};

class Solution {
public:
    int shortestPath(vector<vector<int>> &grid, int k) {
        // 特殊情况
        int n = (int) grid.size();
        if (n == 0)
            return -1;
        int m = (int) grid[0].size();
        if (n == 1 && m == 1)
            return 0;
        // 最短路径上最多出现 m+n-3 个障碍物
        if (k >= m + n - 3)
            return m + n - 2;

        bool vis[n][m][k + 1];
        memset(vis, false, sizeof(vis));
        queue<state> q;
        q.emplace(0, 0, k);
        vis[0][0][k] = true;
        int ans = 1;

        while (!q.empty()) {
            int cnt = (int) q.size();
            while (cnt--) {
                auto s = q.front();
                q.pop();

                for (auto dir: direction) {
                    int new_x = s.x + dir[0];
                    int new_y = s.y + dir[1];
                    if (new_x == -1 || new_x == n || new_y == -1 || new_y == m)
                        continue;   // 越界
                    if (new_x == n - 1 && new_y == m - 1)
                        return ans; // 到达目的地
                    if (grid[new_x][new_y] == 1 && s.z > 0 && !vis[new_x][new_y][s.z - 1]) {
                        // 可破除障碍物
                        q.emplace(new_x, new_y, s.z - 1);
                        vis[new_x][new_y][s.z - 1] = true;
                    } else if (grid[new_x][new_y] == 0 && !vis[new_x][new_y][s.z]) {
                        // 无障碍物
                        q.emplace(new_x, new_y, s.z);
                        vis[new_x][new_y][s.z] = true;
                    }
                }
            }
            ++ans;
        }
        return -1;
    }
};

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    auto s = new Solution();
    vector<vector<int>> grid = {{0, 0, 0},
                                {1, 1, 0},
                                {0, 0, 0},
                                {0, 1, 1},
                                {0, 0, 0}};

    cout << s->shortestPath(grid, 1); // 6
    return 0;
}

