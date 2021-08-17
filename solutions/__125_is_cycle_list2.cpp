//
// Created by peanut on 2021/8/17.
//

#include <bits/stdc++.h>

using namespace std;

struct ListNode {
    int val;
    ListNode *next;

    ListNode(int x) : val(x), next(NULL) {}
};

ListNode* detectCycle(ListNode *head) {
    ListNode *slow = head;
    ListNode *fast = slow;
    while (fast != nullptr) {
        slow = slow->next;
        if (fast->next == nullptr)
            return nullptr;
        fast = fast->next->next;
        if (slow == fast) {
            while (head != slow) {
                head = head->next;
                slow = slow->next;
            }
            return head;
        }
    }
    return nullptr;
}