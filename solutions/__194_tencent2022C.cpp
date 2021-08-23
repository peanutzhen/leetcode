/*
 *      Author: peanutzhen
 *      Created time: 2021/8/22 20:15:32
 */

#include <bits/stdc++.h>

using namespace std;

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt;
    int n;
    int w;
    int a[1005];


    cin >> tt;
    while (tt--) {
        cin >> n >> w;
        for (int i = 0; i < n; ++i) {
            cin >> a[i];
        }
        // 倒序排序
        sort(a, a+n, [](int a, int b){
            return a > b;
        });

        stack<int> odd;
        stack<int> even;
        int ans = 0;
        // 双栈贪心上船
        for (int i = 0; i < n; ++i) {
            if (a[i] % 2 == 1) {
                if (!odd.empty() && odd.top() + a[i] <= w) {
                    odd.pop();
                    ++ans;
                }
                else
                    odd.push(a[i]);
            }
            else {
                if (!even.empty() && even.top() + a[i] <= w) {
                    even.pop();
                    ++ans;
                }
                else
                    even.push(a[i]);
            }
        }
        cout << ans + odd.size() + even.size() << "\n";
    }
    return 0;
}

