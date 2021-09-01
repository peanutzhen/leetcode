/*
 *      Author: peanutzhen
 *      Created time: 2021/8/31 19:26:45
 */

#include <bits/stdc++.h>

using namespace std;

// 获得棍数量为n时，返回围成最大的大正方形时，小正方形的个数
// 事实上 就是二分查找插入点的前一个点
unsigned long long binary_search(unsigned long long num) {
    // 以大正方形边长二分查找
    unsigned long long low = 1, high = 1e8;
    while (low <= high) {
        unsigned long long mid = (low + high) / 2;
        unsigned long long cnt = 2 * mid * (mid + 1);
        if (num <= cnt) {
            high = mid - 1;
        } else
            low = mid + 1;
    }
    return high * high;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    unsigned long long n;

    int T;
    cin >> T;
    while (T--) {
        cin >> n;
        auto ans = binary_search(n);
        // x为ans的边长
        auto x = (unsigned long long) sqrt(ans);
        // 更新剩余木棍数量
        n = (n >= 2 * x * (x + 1)) ? n - 2 * x * (x + 1) : 0;
        // 若剩余木棍至少构造x个正方形
        if (n > 2 * x + 1) {
            ans += x;
            n -= 2 * x + 1; // 3 2 2 2... -> 2n + 1
        }
        if (n > 0)
            ans += (n - 1) / 2;
        cout << ans << "\n";
    }
    return 0;
}

