/*
 *      Author: peanutzhen
 *      Created time: 2021/8/19 01:06:42
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

int ans = 0;
int tmp = 0;

void travesal(TreeNode* root) {
    if (root == nullptr)
        return;
    tmp = (tmp*10) + root->val;
    if (root->left == nullptr && root->right == nullptr)
        ans += tmp;
    else {
        travesal(root->left);
        travesal(root->right);
    }
    tmp -= root->val;
    tmp /= 10;
}

int sumNumbers(TreeNode *root) {
    ans = 0;
    travesal(root);
    return ans;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    cout << sumNumbers(new TreeNode(1, new TreeNode(2), new TreeNode(3)));
    return 0;
}

