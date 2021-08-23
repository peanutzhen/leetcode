/*
 *      Author: peanutzhen
 *      Created time: 2021/8/23 18:11:04
 */

#include <bits/stdc++.h>

using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;

    TreeNode() : val(0), left(nullptr), right(nullptr) {}

    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}

    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

void dfs(vector<int> &ans, TreeNode *r, int depth) {
    if (r == nullptr)
        return;
    if (depth == (int) ans.size())
        ans.push_back(r->val);
    dfs(ans, r->right, depth + 1);
    dfs(ans, r->left, depth + 1);
}

vector<int> rightSideView(TreeNode *root) {
    vector<int> ans;
    dfs(ans, root, 0);
    return ans;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);


    return 0;
}

