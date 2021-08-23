/*
 *      Author: peanutzhen
 *      Created time: 2021/8/22 13:58:30
 */

#include <bits/stdc++.h>

using namespace std;

int maxProfit(vector<int>& prices) {
    int ans = 0;
    int n = (int) prices.size();
    if (n < 2)
        return ans;
    // dp[i][cnt][status] 代表第i天时已经完成了cnt次交易 status 0/1代表不持有/持有股票中
    int dp[n][3][2];
    memset(dp, 0, sizeof(dp));
    dp[0][0][1] -= prices[0];   // 第一天购入股票
    dp[0][1][0] = 0;            // 第一天购入后当天又卖掉了
    dp[0][1][1] -= prices[0];   // 第一天购入后当天又卖掉了 然后又买了
    dp[0][2][0] = 0;            // 第一天买卖了两次 股票鬼才了属于是
    for (int i = 1; i < n; ++i) {
        // Todo: 这里可以用滚动数组优化
        dp[i][0][1] = max(dp[i-1][0][1], dp[i-1][0][0]-prices[i]);
        dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][0][1]+prices[i]);
        dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][1][0]-prices[i]);
        dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][1][1]+prices[i]);
//        cout << "-------------------------------------------\n";
//        cout << "dp[" << i << "][0][0] = " << dp[i][0][0] << "\n";
//        cout << "dp[" << i << "][0][1] = " << dp[i][0][1] << "\n";
//        cout << "dp[" << i << "][1][0] = " << dp[i][1][0] << "\n";
//        cout << "dp[" << i << "][1][1] = " << dp[i][1][1] << "\n";
//        cout << "dp[" << i << "][2][0] = " << dp[i][2][0] << "\n";
    }
    return max(dp[n-1][1][0], dp[n-1][2][0]);
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    vector<int> prices = {1,2,3,4,5};
    cout << maxProfit(prices);
    return 0;
}

