/*
 *      Author: peanutzhen
 *      Created time: 2021/8/20 23:56:23
 */

#include <bits/stdc++.h>

using namespace std;

// Input: 给定长度为n的数组，数组元素为0/1
// n
// a1 a2 ... an
// Output:
// ans
int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);


    int n;
    cin >> n;
    vector<int> b;
    // 当b[i] == b[j]时 说明a[i+1, ..., j]为符合期望的字符列
    unordered_map<int, int> map;
    // 使用map记录b[i]的值 这样遇到相等的值时 可以O(1)查询得到下标
    int cnt[2] = {0, 0};
    int ans = 0;
    for (int i = 0; i < n; ++i) {
        int ai;
        cin >> ai;
        cnt[ai]++;
        b.push_back(cnt[0] - cnt[1]);
        if (map.find(b.back()) != map.end()) {
            int tmp_len = i - map[b.back()];
            if (tmp_len > ans)
                ans = tmp_len;
        }
        else
            map[b.back()] = i;
    }
    cout << ans;
    return 0;
}

