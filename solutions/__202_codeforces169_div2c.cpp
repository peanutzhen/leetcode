/*
 *      Author: peanutzhen
 *      Created time: 2021/8/24 22:03:59
 */

#include <bits/stdc++.h>

using namespace std;

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int n, q, l, r;
    cin >> n >> q;
    // ai <= 2 * 1e5 cnti <= 2 * 1e5
    // 不能用int去存 否则 ai * cnti overflowed.
    vector<unsigned long long> a(n + 1), cnt(n + 2, 0);
    for (int i = 1; i <= n; ++i)
        cin >> a[i];

    // 核心代码 以O(n)计算每位查询次数
    while (q--) {
        cin >> l >> r;
        ++cnt[l];
        --cnt[r + 1];
    }
    for (int i = 1; i <= n; ++i)
        cnt[i] += cnt[i - 1];

    // 贪心 查询次数最多的应该放置较大的元素 使得和最大
    sort(a.rbegin(), a.rend() - 1);
    sort(cnt.rbegin() + 1, cnt.rend() - 1);

    unsigned long long ans = 0LL;
    for (int i = 1; i <= n; ++i) {
        ans += cnt[i] * a[i];
    }
    cout << ans << "\n";
    return 0;
}

