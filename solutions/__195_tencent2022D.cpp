/*
 *      Author: peanutzhen
 *      Created time: 2021/8/22 21:08:16
 */

#include <bits/stdc++.h>

using namespace std;

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int n, k;
    string s;
    cin >> n >> k >> s;

    int ptr = 0;
    // 贪心获取字典序最优字符
    // 存在多个最优字符时，选择首个位置
    while (k > 0) {
        pair<char, int> best_char = make_pair('\0', -1);
        for (int i = ptr; i <= n - k; i++)
            if (s[i] > best_char.first)
                best_char = make_pair(s[i], i);
        ptr = best_char.second + 1;
        --k;
        cout << best_char.first;
    }
    return 0;
}

