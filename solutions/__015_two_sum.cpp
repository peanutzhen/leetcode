//
// Created by peanut on 2021/8/17.
//

#include <bits/stdc++.h>

using namespace std;

vector<int> twoSum(vector<int>& nums, int target) {
    unordered_map<int,int> map = {{nums[0], 0}};
    for (int i = 1; i < nums.size(); ++i) {
        auto it = map.find(target-nums[i]);
        if (it != map.end()) {
            return {it->second, i};
        }
        map[nums[i]] = i;
    }
    return {};
}
