/*
 *      Author: peanutzhen
 *      Created time: 2021/9/1 21:19:00
 */

#include <bits/stdc++.h>

using namespace std;

vector<vector<int>> ans;
int n;

void backtrace(vector<int> &nums, int depth) {
    if (depth == n) {
        ans.push_back(nums);
        return;
    }
    for (int i = depth; i < n; ++i) {
        swap(nums[depth], nums[i]);
        backtrace(nums, depth + 1);
        swap(nums[depth], nums[i]);
    }
}

vector<vector<int>> permute(vector<int> &nums) {
    ans.clear();
    n = (int) nums.size();
    backtrace(nums, 0);
    return ans;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);


    return 0;
}

