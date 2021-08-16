//
// Created by peanut on 2021/8/16.
//

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

int ans = INT_MIN;

int maxGain(TreeNode *root) {
    if (root == nullptr)
        return 0;
    int left_max_gain = maxGain(root->left);
    int right_max_gain = maxGain(root->right);
    int tmp = left_max_gain + right_max_gain + root->val;
    if (tmp > ans)
        ans = tmp;
    int rtv = max(left_max_gain, right_max_gain) + root->val;
    return max(0, rtv);
}

int maxPathSum(TreeNode *root) {
    if (root == nullptr)
        return 0;
    maxGain(root);
    return ans;
}
