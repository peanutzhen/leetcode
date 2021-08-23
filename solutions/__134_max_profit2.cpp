/*
 *      Author: peanutzhen
 *      Created time: 2021/8/22 13:54:07
 */

#include <bits/stdc++.h>

using namespace std;

// 把 prices 中的所有递增区间累加即可
int maxProfit(vector<int>& prices) {
    int ans = 0;
    int n = (int) prices.size();
    if (n < 2)
        return ans;
    for (int i = 1; i < n; ++i) {
        if (prices[i] > prices[i-1])
            ans += prices[i] - prices[i-1];
    }
    return ans;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);


    return 0;
}

