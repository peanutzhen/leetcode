//
// Created by peanut on 2021/8/16.
//

#include <bits/stdc++.h>

using namespace std;

int integerBreak(int n) {
    vector<int> dp(n + 1);
    dp[2] = 1;
    for (int i = 3; i <= n; i++) {
        for (int j = 1; j < i - 1; j++) {
            dp[i] = max(dp[i], max((i - j) * j, dp[i - j] * j));
        }
    }
    return dp[n];
}

int main() {
    cout << integerBreak(10) << endl;
    return 0;
}
