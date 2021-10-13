/*
 *      Author: peanutzhen
 *      Created time: 2021/10/13 20:29:40
 */

#include <bits/stdc++.h>

using namespace std;

// Definition for a Node.
class Node {
public:
    int val;
    Node *left;
    Node *right;
    Node *next;

    Node() : val(0), left(NULL), right(NULL), next(NULL) {}

    Node(int _val) : val(_val), left(NULL), right(NULL), next(NULL) {}

    Node(int _val, Node *_left, Node *_right, Node *_next)
            : val(_val), left(_left), right(_right), next(_next) {}
};

// 引入头节点简化代码设计典范
class Solution {
public:
    Node *connect(Node *root) {
        Node *node = root;
        while (node != nullptr) {
            Node *dummy = new Node();
            Node *pre = dummy;
            while (node != nullptr) {
                if (node->left != nullptr) {
                    pre->next = node->left;
                    pre = pre->next;
                }
                if (node->right != nullptr) {
                    pre->next = node->right;
                    pre = pre->next;
                }
                node = node->next;
            }
            node = dummy->next;
        }

        return root;
    }
};

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);


    return 0;
}