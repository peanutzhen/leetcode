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

bool hasCycle(ListNode *head) {
    ListNode *slow = head;
    ListNode *fast = slow;
    while (fast != nullptr) {
        slow = slow->next;
        if (fast->next == nullptr)
            return false;
        fast = fast->next->next;
        if (slow == fast)
            return true;
    }
    return false;
}