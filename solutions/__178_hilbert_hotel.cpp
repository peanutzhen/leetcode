/*
 *      Author: peanutzhen
 *      Created time: 2021/8/18 20:34:41
 */

#include <bits/stdc++.h>

using namespace std;

#define MAX_LEN ((int)2e5+3)
bool cnt[MAX_LEN];

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int tt;
    cin >> tt;
    int n;
    while (tt--) {
        bool f = true;
        cin >> n;
        int ai;
        int shift;
        for (int i = 0; i < n; ++i) {
            cin >> ai;
            shift = (ai >= 0) ? (i + ai) % n : (n + ((i + ai) % n)) % n;
            if (cnt[shift])
                f = false;  // 只要多名客人在一个房间 就是NO
            else
                cnt[shift] = true;
        }
        memset(cnt, false, sizeof(cnt));
        if (f)
            cout << "YES" << "\n";
        else
            cout << "NO" << "\n";
    }
    return 0;
}

