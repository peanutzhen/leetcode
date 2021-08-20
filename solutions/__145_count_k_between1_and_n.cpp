/*
 *      Author: peanutzhen
 *      Created time: 2021/8/20 16:09:05
 */

#include <bits/stdc++.h>

using namespace std;

// 核心思路：分别统计每一位出现1的次数
int countDigitOne(int n) {
    int cnt = 0;
    int cur = 1;    // cur代表现在统计的位 1(个位) 10(十位) 100(百位) ...
    int right = 0;  // cur右边的数字 如 n = 2593 cur =3 right = 93
    while (n > 0) {
        int r = n % 10;
        n /= 10;
        cnt += n * cur;
        if (r == 1)
            cnt += right + 1;
        else if (r > 1)
            cnt += cur;
        right += cur * r;
        cur *= 10;
    }
    return cnt;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    cout << countDigitOne(13);
    return 0;
}

