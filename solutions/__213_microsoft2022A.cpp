/*
 *      Author: peanutzhen
 *      Created time: 2021/9/21 15:16:06
 */

#include <bits/stdc++.h>

using namespace std;

// 题目描述：
// 给定三位数N(100~999)和整数K(1~30)
// 你可以在N的任意位上加K次1
// 使得得到的N尽可能大
// IO example:
// N = 285 K = 20 -> 999
// N = 123 K = 4 -> 523
int solution(int N, int K) {
    int ar[3];
    for (int i = 2; i >= 0; --i) {
        int rem = N % 10;
        N /= 10;
        ar[i] = rem;
    }
    for (int i = 0; i < 3 && K > 0; ++i) {
        int inc = min(K, 9 - ar[i]);
        ar[i] += inc;
        K -= inc;
    }
    return ar[0] * 100 + ar[1] * 10 + ar[2];
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    cout << solution(285, 20);
    return 0;
}

