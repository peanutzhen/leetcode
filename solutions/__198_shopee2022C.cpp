/*
 *      Author: peanutzhen
 *      Created time: 2021/8/23 15:26:09
 */

#include <bits/stdc++.h>

using namespace std;


struct ListNode {
    int val;
    struct ListNode *next;

    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *nxt) : val(x), next(nxt) {}
};

class Solution {
public:
    /**
     * Note: 类名、方法名、参数名已经指定，请勿修改
     *
     *
     * k个节点一组，逆转单链表
     * @param head ListNode类 头节点
     * @param k int整型 组大小
     * @return ListNode类
     */
    ListNode *revertLinkList(ListNode *head, int k) {
        // write code here
        auto *dummy = new ListNode(-1);
        dummy->next = head;
        ListNode *prev = dummy;

        while (prev != nullptr) {
            ListNode *end = prev;
            ListNode *nxt;
            pair<ListNode *, ListNode *> rtv = make_pair(nullptr, nullptr);
            for (int i = 0; i < k && end != nullptr; ++i)
                end = end->next;

            if (end != nullptr) {
                nxt = end->next;
                rtv = reverseSubList(prev->next, end);
                prev->next = rtv.first;
                rtv.second->next = nxt;
            }
            prev = rtv.second;
        }
        return dummy->next;
    }

    pair<ListNode *, ListNode *> reverseSubList(ListNode *start, ListNode *end) {
        if (start == end)
            return {end, start};
        auto *head = new ListNode(-1);
        head->next = start;
        ListNode *prev = head;
        ListNode *cur = head->next;

        while (prev != end) {
            ListNode *nxt = cur->next;
            cur->next = prev;
            prev = cur;
            cur = nxt;
        }
        // 返回 {新链头, 链尾}
        return {prev, head->next};
    }
};


int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    auto *s = new Solution();
    auto *head = new ListNode(1, new ListNode(2, new ListNode(3)));
    s->revertLinkList(head, 2);
    return 0;
}

