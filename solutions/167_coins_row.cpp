#include <bits/stdc++.h>

using namespace std;

// O(n) 更新bob路径长
int main() {
    int a[2][100003];
    int t;
    scanf("%d", &t);
    while (t > 0) {
        int m;
        scanf("%d", &m);
        int up = 0; // bob 到 m-1 才down的路径和
        int down = 0; // bob 在 0 就down的路径和
        for (int i = 0; i < 2; ++i) {
            for (int j = 0; j < m; ++j) {
                scanf("%d", &a[i][j]);
                if (i == 0) {
                    up += a[i][j];
                }
            }
        }
        up -= a[0][0]; // alice拿走了coins
        int score = max(up, down);
        for (int i = 1; i < m; ++i) {
            down += a[1][i-1];
            up -= a[0][i];
            score = min(score,  max(up, down));
        }
        printf("%d\n", score);
        t--;
    }
    return 0;
}
