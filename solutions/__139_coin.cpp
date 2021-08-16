//
// Created by peanut on 2021/8/16.
//

#include <bits/stdc++.h>

using namespace std;

int coinChange(vector<int> &coins, int amount) {
    vector<int> dp(amount + 1, INT_MAX);
    dp[0] = 0;
    for (int i = 1; i <= amount; ++i) {
        for (int coin: coins) {
            if (i - coin >= 0 && dp[i - coin] != INT_MAX) {
                dp[i] = min(dp[i], dp[i - coin] + 1);
            }
        }
    }
    if (dp[amount] == INT_MAX) {
        return -1;
    }
    return dp[amount];
}
