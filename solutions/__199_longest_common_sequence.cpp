/*
 *      Author: peanutzhen
 *      Created time: 2021/8/23 18:27:47
 */

#include <bits/stdc++.h>

using namespace std;

int longestCommonSubsequence(string text1, string text2) {
    int n = (int) text1.size();
    int m = (int) text2.size();
    if (n == 0 || m == 0)
        return 0;
    vector<vector<int>> dp(n + 1, vector<int>(m + 1, 0));
    for (int i = 1; i <= n; ++i) {
        for (int j = 1; j <= m; ++j) {
            if (text1[i - 1] == text2[j - 1])
                dp[i][j] = dp[i - 1][j - 1] + 1;
            else
                dp[i][j] = max(dp[i-1][j], dp[i][j-1]);
        }
    }
    return dp[n][m];
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    cout << longestCommonSubsequence("abcde", "ace");
    return 0;
}

