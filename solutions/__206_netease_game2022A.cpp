/*
 *      Author: peanutzhen
 *      Created time: 2021/8/28 15:06:46
 */

#include <bits/stdc++.h>

using namespace std;

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int T;
    cin >> T;

    int n, A, B;                // 怪物数量/英雄属性
    int a[105], b[105], h[105]; // 怪物属性
    int y[105];                 // 怪物造成总伤害
    int z[105];                 // 杀死该怪物回的血量
    int attack, round;          // 英雄对怪物造成伤害/怪物攻击轮数
    while (T--) {
        int x = -1;             // 答案
        bool flag = true;       // 通关标志
        cin >> n >> A >> B;
        for (int i = 0; i < n; ++i)
            cin >> a[i] >> b[i] >> h[i];
        for (int i = 0; i < n; ++i) {
            attack = (A - b[i]) < 0 ? 0 : (A - b[i]);
            if (attack == 0) {
                flag = false;   // 对怪物造成伤害为0 无法通关
                break;
            }
            round = h[i] / attack;
            if (h[i] % attack == 0)
                --round;
            y[i] = round * max(a[i] - B, 0);
            z[i] = (h[i] % attack == 0) ? 0 : attack - (h[i] % attack);
        }
        // 计算最小所需生命值x
        if (flag) {
            x = y[n - 1] + 1;
            for (int i = n - 2; i >= 0; --i) {
                x = x + y[i];
                x = (x - z[i] >= y[i] + 1) ? x - z[i] : y[i] + 1;
            }
        }
        cout << x << "\n";
    }
    return 0;
}

