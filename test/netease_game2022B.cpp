/*
 *      Author: peanutzhen
 *      Created time: 2021/8/28 15:06:52
 */

#include <bits/stdc++.h>

using namespace std;

void topsort(vector<vector<pair<int, int>>> &graph, vector<int> &in_deg, vector<int> &dp) {
    queue<int> q;
    // 找出没有前置任务的任务
    for (int i = 0; i < (int) in_deg.size(); ++i) {
        if (in_deg[i] == 0)
            q.push(i);
    }
    while (!q.empty()) {
        int job_id = q.front();
        q.pop();
        for (int i = 0; i < (int) graph[job_id].size(); ++i) {
            int v = graph[job_id][i].first;
            int c = graph[job_id][i].second;
            --in_deg[v];
            if (in_deg[v] == 0)
                q.push(v);
            dp[v] = max(dp[v], dp[job_id] + c);
        }
        graph[job_id].clear();
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int T;
    int n;  // 任务个数
    int m;  // 任务前置任务数
    vector<pair<int, vector<int>>>  info;
    vector<vector<pair<int, int>>>  graph;
    vector<int>                     in_deg;
    vector<int>                     out_deg;
    vector<int>                     dp;

    cin >> T;
    while (T--) {
        cin >> n;
        info    = vector<pair<int, vector<int>>>(n);
        graph   = vector<vector<pair<int, int>>>(n + 1);
        in_deg  = vector<int>(n + 1, 0);
        out_deg = vector<int>(n + 1, 0);
        dp      = vector<int>(n + 1, 0);
        for (int i = 0; i < n; ++i) {
            info[i] = make_pair(0, vector<int>());
            cin >> info[i].first >> m;
            info[i].second = vector<int>(m);
            for (int j = 0; j < m; ++j) {
                cin >> info[i].second[j];
                --info[i].second[j];
            }
        }
        // info转邻接矩阵graph
        for (int i = 0; i < n; ++i) {
            for (int j = 0; j < (int) info[i].second.size(); ++j) {
                graph[info[i].second[j]].emplace_back(i, info[j].first);
                ++in_deg[i];
                ++out_deg[info[i].second[j]];;
            }
        }
        // 对于出度为0的任务 指向结束任务
        for (int i = 0; i < n; ++i) {
            if (out_deg[i] == 0) {
                graph[i].emplace_back(n, info[i].first);
                ++in_deg[n];
            }
        }
        topsort(graph, in_deg, dp);
        cout << *max_element(dp.begin(), dp.end()) << "\n";
    }
    return 0;
}

