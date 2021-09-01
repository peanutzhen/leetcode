/*
 *      Author: peanutzhen
 *      Created time: 2021/9/1 21:19:55
 */

#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
    vector<vector<int>> ans;        // 保存答案
    vector<bool>        visited;    // 该数字是否被填入
    vector<int>         permute;    // 一种排列的可能
    int                 n;          // 数字数量

    void backtrace(vector<int> &nums, int depth) {
        if (depth == n) {
            ans.push_back(permute);
            return;
        }
        for (int i = 0; i < n; ++i) {
            // 去重
            if (!visited[i] && (i == 0 || nums[i] != nums[i-1] || visited[i-1])) {
                visited[i] = true;
                permute[depth] = nums[i];
                backtrace(nums, depth+1);
                visited[i] = false;
            }
        }
    }

    vector<vector<int>> permuteUnique(vector<int>& nums) {
        ans.clear();
        n = (int) nums.size();
        sort(nums.begin(), nums.end());
        visited = vector<bool>(n, false);
        permute = vector<int>(n);
        backtrace(nums, 0);
        return ans;
    }
};

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    auto nums = vector<int>{1,1,2};
    auto s = new Solution();
    s->permuteUnique(nums);
    return 0;
}

