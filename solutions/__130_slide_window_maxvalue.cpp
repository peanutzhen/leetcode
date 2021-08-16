//
// Created by peanut on 2021/8/16.
//

#include <bits/stdc++.h>

using namespace std;

vector<int> maxSlidingWindow(vector<int>& nums, int k) {
    priority_queue<pair<int,int>> pq;
    for (int i = 0; i < k; ++i) {
        pq.emplace(nums[i], i);
    }
    vector<int> max_window = {pq.top().first};
    for (int i = k; i < nums.size(); ++i) {
        pq.emplace(nums[i], i);
        while (pq.top().second < i-k+1) {
            pq.pop();
        }
        max_window.push_back(pq.top().first);
    }
    return max_window;
}
