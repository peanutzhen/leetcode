#include <stdbool.h>
#include <stdlib.h>

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

/*
 * 这道题，一开始想用递归做，但由于被给定参数（只能有一个参数root）
 * 限制了我的思维。导致没有成功的拆分子问题，哎一道简单题都做不出来。。
 */

bool check(struct TreeNode* p, struct TreeNode* q)
{
    // 他这个写法好巧妙
    // 第二个if只会是 1 0 或 0 1，而 1 1 这种情况被第一个if控制
    if(!p && !q)
        return true;
    if(!p || !q)
        return false;
    return p->val == q->val && check(p->left,q->right) && check(p->right,q->left);
}

bool isSymmetric(struct TreeNode* root){
    return check(root,root);
}