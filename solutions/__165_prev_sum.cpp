#include <bits/stdc++.h>

using namespace std;

// 做法与Java一样 set + lower_bound获取 prev
int main() {
    int n;
    cin >> n;
    vector<int> v(n);
    for (int i = 0; i < n; ++i) {
        cin >> v[i];
    }

    set<int> s{v[0]};
    int ans = 0;
    for (int i = 1; i < n; ++i) {
        auto it = s.lower_bound(v[i]);
        int prev = *(--it);
        if (prev < v[i]) {
            ans += prev * (i+1);
        }
        s.insert(v[i]);
    }
    cout << ans << endl;
    return 0;
}
