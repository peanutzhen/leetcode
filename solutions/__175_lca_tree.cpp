//
// Created by peanut on 2021/8/16.
//

#include <bits/stdc++.h>

using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;

    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

unordered_map<int, TreeNode*> parent;
unordered_map<int, bool> vis;

void dfs(TreeNode* root) {
    if (root->left != nullptr) {
        parent[root->left->val] = root;
        dfs(root->left);
    }
    if (root->right != nullptr) {
        parent[root->right->val] = root;
        dfs(root->right);
    }
}

TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
    if (root == nullptr || p == nullptr || q == nullptr) {
        return nullptr;
    }
    parent[root->val] = nullptr;
    dfs(root);
    while (p != nullptr) {
        vis[p->val] = true;
        p = parent[p->val];
    }
    while (q != nullptr) {
        if (vis[q->val]) return q;
        q = parent[q->val];
    }
    return nullptr;
}