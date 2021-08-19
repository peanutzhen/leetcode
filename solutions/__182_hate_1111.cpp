/*
 *      Author: peanutzhen
 *      Created time: 2021/8/19 13:32:21
 */

#include <bits/stdc++.h>

using namespace std;

// 问题等价于x == 11 * a + 111 * b，其中a,b in [0,+inf]。
// 那么又等价于b = 11 * c + d,其中c in [0,+inf], d in [0,10] 。
// 即x == (111 * c + a) * 11+111 * d
int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt;
    cin >> tt;
    int x;

    while (tt--) {
        cin >> x;
        bool f = false;
        for (int i = 0; i <= 10; ++i) {
            if (x % 11 == 0) {
                f = true;
                break;
            }
            x -= 111;
            if (x < 0)
                break;
        }
        if (f)
            cout << "YES\n";
        else
            cout << "NO\n";
    }
    return 0;
}

