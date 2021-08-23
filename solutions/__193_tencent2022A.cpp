/*
 *      Author: peanutzhen
 *      Created time: 2021/8/22 21:46:32
 */

#include <bits/stdc++.h>

using namespace std;

struct ListNode {
    int val;
    struct ListNode *next;

    ListNode(int x) :
            val(x), next(NULL) {
    }
};

class Solution {
public:
    /**
     * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
     *
     * @param m int整型
     * @param a ListNode类 指向彩带的起点，val表示当前节点的val，next指向下一个节点
     * @return ListNode类vector
     */
    vector<ListNode *> solve(int m, ListNode *a) {
        // write code here
        vector<ListNode *> rtv(m);
        vector<ListNode *> endp(m);
        for (int i = 0; i < m; i++) {
            rtv[i] = new ListNode(-1);
            endp[i] = rtv[i];
        }
        ListNode* ptr = a;
        int color;
        while (ptr != nullptr) {
            color = ptr->val % m;
            endp[color]->next = ptr;
            endp[color] = endp[color]->next;
            ptr = ptr->next;
            endp[color]->next = nullptr;
        }
        for (int i = 0; i < m; i++) {
            rtv[i] = rtv[i]->next;
        }
        return rtv;
    }
};

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);


    return 0;
}

