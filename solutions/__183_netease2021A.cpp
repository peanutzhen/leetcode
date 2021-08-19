/*
 *      Author: peanutzhen
 *      Created time: 2021/8/19 20:32:54
 */

#include <bits/stdc++.h>

using namespace std;

int a[20];
int n;
int sum; // a[n]总和
int ans;

void backtrace(int left, int right, int i) {
    if (i == n) {
        if (left == right) {
            ans = min(ans, sum - left - right);
        }
        return;
    }
    // 剪枝
    if (left > sum / 2 || right > sum / 2) {
        return;
    }
    // 三种情况: (1)给left (2)给right (3)丢弃
    backtrace(left + a[i], right, i + 1);
    backtrace(left, right + a[i], i + 1);
    backtrace(left, right, i + 1);
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt;
    cin >> tt;
    while (tt--) {
        sum = 0;
        ans = INT_MAX;
        cin >> n;
        for (int i = 0; i < n; ++i) {
            cin >> a[i];
            sum += a[i];
        }
        backtrace(0, 0, 0);
        cout << ans << "\n";
    }
    return 0;
}

