/*
 *      Author: peanutzhen
 *      Created time: 2021/8/20 11:25:23
 */

#include <bits/stdc++.h>

using namespace std;

#define point int

int n, m;
vector<unordered_set<int>> graph;  // 题目说有重复边 所以用了set
long long ans = 0;

int timestamp = 0;
vector<int> dfn, low;
stack<point> s;
vector<bool> in_s;

// Tarjan算法 dfn记录节点u的访问次序 low记录节点u能与最低次序i连通
// 当 dfn[u] == low[u] 时，说明这是连通分支里的root节点
void tarjan(point u) {
    dfn[u] = low[u] = ++timestamp;
    s.push(u);
    in_s[u] = true;
    for (auto v: graph[u]) {
        if (!dfn[v]) {
            tarjan(v);
            low[u] = min(low[u], low[v]);
        }
        else if (in_s[v]) {
            low[u] = min(low[u], dfn[v]);
        }
    }
    if (dfn[u] == low[u]) {
        // 找到一个连通分支的根节点
        int cnt = 0; // 统计该连通分支点个数 即栈中在节点u之后进入的节点都是连通的
        point t;
        do {
            t = s.top();
            s.pop();
            in_s[t] = false;
            ++cnt;
        } while (t != u);
        ans += (long long) (cnt * (cnt - 1) / 2); // 组合计算公式
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    cin >> n >> m;
    graph = vector<unordered_set<int>>(n + 1);
    dfn = vector<int>(n + 1);
    low = vector<int>(n + 1);
    in_s = vector<bool>(n+1, false);
    int u, v;
    for (int i = 0; i < m; ++i) {
        cin >> u >> v;
        graph[u].insert(v);
    }
    for (int p = 1; p <= n; ++p) {
        if (!dfn[p])
            tarjan(p);
    }
    cout << ans << "\n";
    return 0;
}

