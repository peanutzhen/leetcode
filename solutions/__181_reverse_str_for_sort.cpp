/*
 *      Author: peanutzhen
 *      Created time: 2021/8/19 11:16:14
 */

#include <bits/stdc++.h>

using namespace std;

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int n;
    cin >> n;
    vector<int> c(n + 1);
    // 一定要是long long 否则会溢出
    vector<vector<long long>> dp(n + 1, vector<long long>(2, LONG_LONG_MAX));

    for (int i = 1; i <= n; ++i) {
        cin >> c[i];
    }
    string prev, cur;
    string reversed_prev, reversed_cur;
    cin >> prev;
    dp[1][0] = 0;
    dp[1][1] = c[1];
    for (int i = 2; i <= n; ++i) {
        cin >> cur;
        reversed_prev = string(prev.rbegin(), prev.rend());
        reversed_cur = string(cur.rbegin(), cur.rend());
        if (dp[i - 1][0] != LONG_LONG_MAX) {
            // 如果第i-1个字符不反转可行
            if (prev <= cur)
                dp[i][0] = dp[i - 1][0];
            if (prev <= reversed_cur)
                dp[i][1] = dp[i - 1][0] + c[i];
        }
        if (dp[i - 1][1] != LONG_LONG_MAX) {
            // 如果第i-1个字符反转可行
            if (reversed_prev <= cur)
                dp[i][0] = min(dp[i][0], dp[i - 1][1]);
            if (reversed_prev <= reversed_cur)
                dp[i][1] = min(dp[i][1], dp[i - 1][1] + c[i]);
        }
        // 如果第i个字符反不反转都不可行 则return
        if (dp[i][0] == LONG_LONG_MAX && dp[i][1] == LONG_LONG_MAX) {
            cout << -1 << "\n";
            return 0;
        }
        prev = cur;
    }
    cout << min(dp[n][0], dp[n][1]) << "\n";
    return 0;
}

