/*
 *      Author: peanutzhen
 *      Created time: 2021/8/21 15:02:44
 */

#include <bits/stdc++.h>

using namespace std;

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);


    string line;
    getline(cin, line);
    istringstream is(line);

    vector<int> a;
    int tmp;
    while (is >> tmp)
        a.push_back(tmp);
    int m;
    cin >> m;

    int ans = 0;
    int n = (int) a.size();
    for (int i = 0; i < n - 1; ++i) {
        for (int j = i + 1; j < n; ++j) {
            if (a[i] + a[j] <= m)
                ++ans;
        }
    }
    cout << ans << "\n";
    return 0;
}

