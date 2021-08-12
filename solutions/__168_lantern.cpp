//
// Created by peanut on 2021/8/12.
//
#include <bits/stdc++.h>

using namespace std;

int main() {
    int n, l;
    scanf("%d %d", &n, &l);
    set<int> points;
    for (int i = 0; i < n; ++i) {
        int point;
        scanf("%d", &point);
        points.insert(point);
    }
    double max_gap = 0;
    double left = *points.begin();
    double right = l - *(--points.end());
    for (auto it = points.begin(); it != --points.end(); ++it) {
        double gap = *next(it) - *it;
        max_gap = max(gap, max_gap);
    }
    printf("%.10f\n", max(max(max_gap/2, left), right));
    return 0;
}
