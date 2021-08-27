/*
 *      Author: peanutzhen
 *      Created time: 2021/8/25 19:56:54
 */

#include <bits/stdc++.h>

using namespace std;

int T;
int n;
int a[1003];

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    cin >> T;
    while (T--) {
        cin >> n;
        for (int i = 1; i <= n; ++i) {
            cin >> a[i];
        }
        int ans = 0;
        while (!is_sorted(a+1, a+n+1)) {
            ++ans;
            int i = (ans % 2 == 0) ? 2 : 1;
            while (i < n) {
                if (a[i] > a[i + 1])
                    swap(a[i], a[i + 1]);
                i += 2;
            }
        }
        cout << ans << "\n";
    }
    return 0;
}

