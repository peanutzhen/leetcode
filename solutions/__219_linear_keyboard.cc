/*
 *      Copyright 2021 peanutzhen
 *      Created time: 2021/11/14 12:17:59
 */

#include <bits/stdc++.h>

using namespace std;

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    map<char, int> pos;
    string kb;
    string word;

    int t;
    cin >> t;
    while (t--) {
        int ans = 0;
        cin >> kb >> word;
        for (int i = 0; i < (int) kb.length(); i++)
            pos[kb[i]] = i;
        for (int i = 1; i < (int) word.length(); i++) {
            ans += abs(pos[word[i]] - pos[word[i-1]]);
        }
        cout << ans << "\n";
    }

    return 0;
}
