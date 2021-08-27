/*
 *      Author: peanutzhen
 *      Created time: 2021/8/25 20:00:28
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

// 递归删除
TreeNode *deleteNode(TreeNode *root, int key) {
    // 待删除节点不存在
    if (root == nullptr)
        return nullptr;
    if (key < root->val)
        root->left = deleteNode(root->left, key);
    else if (key > root->val) {
        root->right = deleteNode(root->right, key);;
    }
    else {
        // key == root->val
        // 三种情况 左右均非空 左非空(前继) 右非空(后继)
        if (root->left && root->right) {
            // 默认取后继
            TreeNode *sur_node = root->right;
            while (sur_node->left != nullptr)
                sur_node = sur_node->left;
            root->val = sur_node->val;  // 右继懒替换待删除节点
            // 删右继
            root->right = deleteNode(root->right, sur_node->val);
        }
        else if (root->left)
            return root->left;  // 让调用者继承左子树
        else if (root->right)
            return root->right; // 让调用者继承右子树
        else
            return nullptr;     // 叶子结点 直接让调用者指向null
    }
    return root;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    return 0;
}

