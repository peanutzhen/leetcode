//
// Created by peanut on 2021/8/14.
//

#include <bits/stdc++.h>

using namespace std;

// 自底向上的动态规划解法
string longestPalindrome(string s) {
    int n = s.length();
    int low = 0, high = 0;
    bool dp[n][n];
    for (int len = 1; len <= n; ++len) {
        for (int i = 0; i <= n-len; ++i) {
            if (len == 1) {
                dp[i][i+len-1] = true;
            }
            else if (len == 2) {
                dp[i][i+len-1] = s[i] == s[i+len-1];
            }
            else {
                dp[i][i+len-1] = s[i] == s[i+len-1] && dp[i+1][i+len-1-1];
            }
            if (dp[i][i+len-1] && len > high-low+1) {
                low = i;
                high = i+len-1;
            }
        }
    }

    return s.substr(low, high-low+1);
}

int main() {
    cout << longestPalindrome("cbbd") << endl;
    return 0;
}