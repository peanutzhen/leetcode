//
// Created by peanut on 2021/8/17.
//

#include <bits/stdc++.h>

using namespace std;

void nextPermutation(vector<int> &nums) {
    int i = nums.size() - 2;
    while (i >= 0 && nums[i] >= nums[i + 1]) {
        i--;
    }
    // 如果 i == -1 说明当前排列为最后一个排列 直接反转即可
    // 否则找第一个严格大于nums[i]的数
    if (i >= 0) {
        int j;
        for (j = nums.size() - 1; nums[j] <= nums[i]; j--) {}
        swap(nums[i], nums[j]);
    }
    reverse(nums.begin() + i + 1, nums.end());
}
