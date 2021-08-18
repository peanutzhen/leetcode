/*
 *      Author: peanutzhen
 *      Created time: 2021/8/18 23:47:25
 */

#include <bits/stdc++.h>

using namespace std;

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    unordered_map<char, long long> hashmap;
    hashmap['B'] = hashmap['S'] = hashmap['C'] = 0;
    long long nb, ns, nc;
    long long pb, ps, pc;
    long long money;
    long long ans = 0;
    char ch;
    while ((ch = cin.get()) != '\n') {
        hashmap[ch]++;
    }
    cin >> nb >> ns >> nc;
    cin >> pb >> ps >> pc;
    cin >> money;
    long long low = 0;
    long long high = 1e14;
    while (low <= high) {
        long long mid = (low + high) >> 1;
        long long cost = max(0LL, hashmap['B'] * mid - nb) * pb + max(0LL, hashmap['S'] * mid - ns) * ps +
                         max(0LL, hashmap['C'] * mid - nc) * pc;
        if (cost == money) {
            ans = mid; // 找到恰好的解 退出
            break;
        }
        else if (cost < money) {
            ans = mid; // 保存当前最优解
            low = mid + 1;
        }
        else
            high = mid - 1;
    }
    cout << ans << "\n";
    return 0;
}

