/*
 *      Author: peanutzhen
 *      Created time: 2021/9/3 10:54:36
 */

#include <bits/stdc++.h>

using namespace std;

bool canPartition(vector<int> &nums) {
    int n = (int) nums.size();
    int sum = accumulate(nums.begin(), nums.end(), 0);
    int target = sum / 2;
    int max_num = *max_element(nums.begin(), nums.end());
    // 特殊情况
    if (n <= 1 || sum % 2 == 1 || max_num > target)
        return false;
    if (target == max_num)
        return true;

    vector<vector<bool>> dp(n, vector<bool>(target+1, false));
    dp[0][0] = true;
    dp[0][nums[0]] = true;
    for (int i = 1; i < n; ++i) {
        for (int j = 0; j <= target; ++j) {
            if (j == 0)
                dp[i][j] = true;
            else if (nums[i] <= j)
                dp[i][j] = dp[i - 1][j] || dp[i - 1][j - nums[i]];
            else
                dp[i][j] = dp[i - 1][j];
        }
    }
    return dp[n - 1][target];
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);


    return 0;
}

