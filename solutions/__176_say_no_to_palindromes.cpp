/*
 *      Author: peanutzhen
 *      Created time: 2021/8/18 17:22:08
 */

#include <bits/stdc++.h>

using namespace std;

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int n, m;
    cin >> n >> m;
    string s;
    cin >> s;

    string t = "abc";
    // dp[cur][j] 表示使用abc的第cur个排列时 s的前j个字符所需替换次数
    vector<vector<int>> dp(6, vector<int>(n + 1));
    int cur = 0;
    do {
        for (int j = 1; j <= n; ++j)
            dp[cur][j] = dp[cur][j - 1] + (s[j - 1] != t[(j - 1) % 3]);
        cur++;
        // 使用next_permutation帮助我们构造排列
    } while (next_permutation(t.begin(), t.end()));

    while (m--) {
        int l, r;
        cin >> l >> r;
        int ans = n;
        for (int i = 0; i < 6; ++i)
            ans = min(ans, dp[i][r] - dp[i][l-1]);
        cout << ans << "\n";
    }
    return 0;
}

