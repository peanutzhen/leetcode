/*
 *      Author: peanutzhen
 *      Created time: 2021/8/20 16:41:51
 */

#include <bits/stdc++.h>

using namespace std;

#define point int
#define color int

vector<pair<int, int>> graph[105];
bool visited_point[105];
bool visited_color[105];

bool dfs(point u, point v, color c) {
    visited_point[u] = true;
    if (u == v)
        return true;
    for (auto e: graph[u]) {
        if (!visited_point[e.first] && e.second == c)
            if (dfs(e.first, v, c))
                return true;
    }
    return false;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    int n, m;
    cin >> n >> m;

    point a, b;
    color c;
    while (m--) {
        cin >> a >> b >> c;
        graph[a].emplace_back(b, c);
        graph[b].emplace_back(a, c);
    }

    int q;
    cin >> q;
    point u, v;
    while (q--) {
        cin >> u >> v;
        int ans = 0;
        memset(visited_color, false, sizeof(visited_color));
        for (int i = 0; i < graph[u].size(); ++i) {
            c = graph[u][i].second;
            // 若节点u仍未尝试过此颜色c
            if (!visited_color[c]) {
                memset(visited_point, false, sizeof(visited_point));
                if (dfs(u, v, c))
                    ++ans;
                visited_color[c] = true;
            }
        }
        cout << ans << "\n";
    }
    return 0;
}

