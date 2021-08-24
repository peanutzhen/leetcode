/*
 *      Author: peanutzhen
 *      Created time: 2021/8/24 10:51:44
 */

#include <bits/stdc++.h>

using namespace std;

int findPeakElement(vector<int> &nums) {
    // O(logN)
    int n = (int) nums.size();
    int low = 0;
    int high = n - 1;
    while (low < high) {
        int mid = (low + high) >> 1;
        if (nums[mid] > nums[mid + 1])
            high = mid;
        else
            low = mid + 1;
    }
    return low;
    // O(N)
//    int n = (int) nums.size();
//    if (n == 1 || nums[0] > nums[1])
//        return 0;
//    if (nums[n - 1] > nums[n - 2])
//        return n - 1;
//    for (int i = 1; i < n - 1; ++i) {
//        if (nums[i] > nums[i - 1] && nums[i] > nums[i + 1])
//            return i;
//    }
//    return -1;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
    vector<int> nums {2};
    cout << findPeakElement(nums);
    return 0;
}

