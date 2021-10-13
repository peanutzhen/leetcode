/*
 *      Author: peanutzhen
 *      Created time: 2021/10/13 20:47:30
 */

#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
    vector<vector<int>> generate(int numRows) {
        vector<vector<int>> ans(numRows);
        for (int i = 1; i <= numRows; ++i) {
            if (i == 1)
                ans[0] = vector<int>{1};
            if (i == 2)
                ans[1] = vector<int>{1, 1};
            ans[i-1] = vector<int>(i);
            ans[i-1][0] = ans[i-1][i-1] = 1;
            for (int j = 1; j < i-1; ++j) {
                ans[i-1][j] = ans[i-2][j] + ans[i-2][j-1];
            }
        }
        return ans;
    }
};

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);


    return 0;
}