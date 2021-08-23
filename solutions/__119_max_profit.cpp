/*
 *      Author: peanutzhen
 *      Created time: 2021/8/22 13:50:21
 */

#include <bits/stdc++.h>

using namespace std;

int maxProfit(vector<int> &prices) {
    int n = (int) prices.size();
    int min_price = INT_MAX;
    int ans = 0;
    for (int i = 0; i < n; ++i) {
        if (prices[i] < min_price)
            min_price = prices[i];
        int tmp = prices[i] - min_price;
        if (tmp > ans)
            ans = tmp;
    }
    return ans;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);


    return 0;
}

