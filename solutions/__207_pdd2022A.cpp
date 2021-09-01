/*
 *      Author: peanutzhen
 *      Created time: 2021/8/31 19:02:13
 */

#include <bits/stdc++.h>

using namespace std;

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int n, m, k;
    int a[10005];
    int b[10005];

    int T;
    cin >> T;
    while (T--) {
        int ans = 0;
        cin >> n >> m >> k;
        for (int i = 0; i < n; ++i)
            cin >> a[i];
        for (int i = 0; i < m; ++i)
            cin >> b[i];
        sort(a, a + n);
        sort(b, b + m);
        int i = 0, j = 0;
        int diff;
        while (i != n && j != m) {
            diff = a[i] - b[j];
            if (abs(diff) <= k) {
                ++ans;
                ++i;
                ++j;
            } else if (diff < 0)
                ++i;
            else if (diff > 0)
                ++j;
        }
        cout << ans << "\n";
    }
    return 0;
}

