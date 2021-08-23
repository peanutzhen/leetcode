/*
 *      Author: peanutzhen
 *      Created time: 2021/8/22 13:58:54
 */

#include <bits/stdc++.h>

using namespace std;

int maxProfit(int k, vector<int> &prices) {
    int n = (int) prices.size();
    if (k == 0 || n < 2)
        return 0;
    int dp[n][k + 1][2];
    memset(dp, 0, sizeof(dp));
    for (int cnt = 0; cnt <= k; ++cnt)
        dp[0][cnt][1] = -prices[0];
    // Todo: 滚动数组优化
    for (int i = 1; i < n; ++i) {
        dp[i][0][1] = max(dp[i - 1][0][1], dp[i - 1][0][0] - prices[i]);
        for (int cnt = 1; cnt <= k; ++cnt) {
            dp[i][cnt][0] = max(dp[i - 1][cnt][0], dp[i - 1][cnt - 1][1] + prices[i]);
            dp[i][cnt][1] = max(dp[i - 1][cnt][1], dp[i - 1][cnt][0] - prices[i]);
        }
    }
    return dp[n - 1][k][0];
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    vector<int> prices{1, 2, 3, 4, 5};
    cout << maxProfit(5, prices);
    return 0;
}

