/*
 *      Author: peanutzhen
 *      Created time: 2021/9/21 16:09:00
 */

#include <bits/stdc++.h>

using namespace std;

// 题目描述：
// 给定N栋房子，每栋房子有自己的能量需求，为正整数
// 第K栋房子能量需求记为A[K]
// 现有两种太阳能
// 一种花费为X，恰好满足房子需求
// 一种花费为Y，在满足房子需求的同时还能向其他房子提供能量，这个能量大小恰好为房子的能量需求
// 在满足所有房子的能量需求下，求最小花费代价
// 注意：可以过剩满足房子能量需求，只要满足就行了
int solution(vector<int> &A, int X, int Y) {
    double rate = double(Y) / X;
    int ans = 0;

    sort(A.rbegin(), A.rend());
    int n = (int) A.size();
    int left = 0;
    int right = n - 1;
    while (left <= right) {
        double cnt = 1;
        int power = A[left];
        int ori_right = right;
        while (right > left && power >= A[right]) {
            power -= A[right];
            ++cnt;
            --right;
        }
        if (cnt >= rate) {
            ans += Y;
            A[right] -= power;
        } else {
            ans += X;
            right = ori_right;
        }
        ++left;
    }

    return ans;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    vector<int> A = {3, 8, 3, 5, 2};
    int X = 2, Y = 5;
    cout << solution(A, X, Y);
    return 0;
}

