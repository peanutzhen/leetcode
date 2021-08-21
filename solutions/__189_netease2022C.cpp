/*
 *      Author: peanutzhen
 *      Created time: 2021/8/21 15:14:18
 */

#include <bits/stdc++.h>

using namespace std;

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    vector<int> a;
    int tmp;
    string line;
    getline(cin, line);
    istringstream is(line);
    while (is >> tmp)
        a.push_back(tmp);

    vector<int> ans(a.size(), 1);
    int n = (int) a.size();
    for (int i = 1; i < n; ++i) {
        if (a[i] > a[i - 1])
            ans[i] = ans[i - 1] + 1;
    }
    for (int i = n - 1; i > 0; --i) {
        if (a[i - 1] > a[i] && ans[i - 1] <= ans[i])
            ans[i-1] = ans[i] + 1;
    }
    // 站成一环 所以最后一个小孩要和第一个小孩比较
    if (a[n-1] > a[0] && ans[n - 1] <= ans[0]) {
        ans[n-1] = ans[0] + 1;
    }
    else if (a[0] > a[n-1] && ans[0] <= ans[n-1]) {
        ans[0] = ans[n-1] + 1;
    }
    cout << accumulate(ans.begin(), ans.end(), 0) << "\n";
    return 0;
}

