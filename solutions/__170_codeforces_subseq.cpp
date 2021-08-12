//
// Created by peanut on 2021/8/12.
//

#include <bits/stdc++.h>

using namespace std;

int main() {
    map<unsigned long long int, pair<int, int> > m;
    for (int i = 2; i <= 40; ++i) {
        for (int count = 0; count < 10; ++count) {
            unsigned long long int sum = 1;
            for (int j = 0; j < count; ++j) {
                sum *= i;
            }
            for (int j = 0; j < 10-count; ++j) {
                sum *= i-1;
            }
            m[sum] = make_pair(i, count);
        }
    }
    unsigned long long int k;
    cin >> k;

    string s = "codeforces";
    int idx = 0;
    int num = m.lower_bound(k)->second.first;
    int count = m.lower_bound(k)->second.second;
    for (int i = 0; i < count; ++i) {
        for (int j = 0; j < num; ++j) {
            printf("%c", s[idx]);
        }
        idx++;
    }
    for (int i = 0; i < 10-count; ++i) {
        for (int j = 0; j < num-1; ++j) {
            printf("%c", s[idx]);
        }
        idx++;
    }
    printf("\n");
    return 0;
}