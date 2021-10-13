/*
 *      Author: peanutzhen
 *      Created time: 2021/10/12 22:19:48
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

class Solution {
public:
    Node *connect(Node *root) {
        if (root == nullptr || root->left == nullptr)
            return root;

        root->left->next = root->right;

        for (Node *begin = root->left; begin->left != nullptr; begin = begin->left) {
            Node *node = begin;
            while (node != nullptr) {
                node->left->next = node->right;
                if (node->next != nullptr)
                    node->right->next = node->next->left;
                node = node->next;
            }
        }

        return root;
    }
};

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);


    return 0;
}