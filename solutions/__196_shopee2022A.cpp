/*
 *      Author: peanutzhen
 *      Created time: 2021/8/23 15:15:05
 */

#include <bits/stdc++.h>

using namespace std;

int search_start(vector<int> &a, int target) {
    int n = (int) a.size();
    int low = 0;
    int high = n - 1;
    while (low <= high) {
        int mid = (low + high) >> 1;
        if (a[mid] < target)
            low = mid + 1;
        else
            high = mid - 1;
    }
    if (low < n && a[low] == target)
        return low;
    return -1;
}

int search_end(vector<int> &a, int target) {
    int n = (int) a.size();
    int low = 0;
    int high = n - 1;
    while (low <= high) {
        int mid = (low + high) >> 1;
        if (a[mid] <= target)
            low = mid + 1;
        else
            high = mid - 1;
    }
    if (high >= 0 && a[high] == target)
        return high;
    return -1;
}

int main() {
    vector<int> nums = {1, 2, 3, 3, 3, 3, 4};
    int target;
    cin >> target;
    printf("%d,%d\n", search_start(nums, target), search_end(nums, target));
    return 0;
}

