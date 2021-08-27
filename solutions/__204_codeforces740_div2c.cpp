/*
 *      Author: peanutzhen
 *      Created time: 2021/8/25 21:07:23
 */

#include <bits/stdc++.h>

using namespace std;

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int T;
    cin >> T;

    int n, k;
    while (T--) {
        cin >> n;
        // 用小根堆保存每个洞穴的最小通过power及其怪物数量
        priority_queue<pair<int, int>, vector<pair<int, int>>, greater<pair<int, int>>> priorityQueue;
        for (int i = 0; i < n; ++i) {
            cin >> k;
            int armor;          // 当前怪物armor
            int max_armor = 0;  // 当前洞穴里的极大armor
            int min_power = 0;  // 通过该洞穴的最小power
            for (int j = 0; j < k; ++j) {
                cin >> armor;
                if (armor > max_armor) {
                    max_armor = armor;
                    int need_power = max_armor + 1 - j;
                    if (need_power > min_power)
                        min_power = need_power;
                }
            }
            priorityQueue.emplace(min_power, k);
        }
        // 英雄最小生命值ans公式 bi为通过洞穴的最小能量 ki为洞穴怪物数
        // bi是从小到大排序的 我们贪心从最小所需能量的洞穴出发
        // 因为如果从一个更大的所需能量洞穴出发 则更小所需能量洞穴也能通过
        // 所以为什么不走所需更小的出发先呢？但是先走小洞穴，出来后未必能进入下一个最小洞穴
        // 因此，ans由max(b1, b2-k1)决定。那么，以此类推:
        // ans=max(b1, b2-k1, b3-(k1+k2), ..., bn-sum(k, 1, n-1))
        int ans = 0;
        int sum_k = 0;
        while (!priorityQueue.empty()) {
            auto p = priorityQueue.top();
            if (p.first - sum_k > ans)
                ans = p.first - sum_k;
            sum_k += p.second;
            priorityQueue.pop();
        }
        cout << ans << "\n";
    }
    return 0;
}

