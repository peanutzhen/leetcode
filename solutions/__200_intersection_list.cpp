/*
 *      Author: peanutzhen
 *      Created time: 2021/8/24 10:17:49
 */

#include <bits/stdc++.h>

using namespace std;

struct ListNode {
    int val;
    ListNode *next;

    ListNode(int x) : val(x), next(NULL) {}
};

ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
    // 双指针做法
    if (!headA || !headB)
        return nullptr;
    ListNode *ptrA = headA, *ptrB = headB;
    while (ptrA != ptrB) {
        ptrA = !ptrA ? headB : ptrA->next;
        ptrB = !ptrB ? headA : ptrB->next;
    }
    return ptrA;

    // 哈希表做法
//    unordered_set<ListNode *> set;
//
//    while (headA != nullptr) {
//        set.insert(headA);
//        headA = headA->next;
//    }
//    while (headB != nullptr) {
//        auto target = set.find(headB);
//        if (target != set.end())
//            return *target;
//        headB = headB->next;
//    }
//    return nullptr;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);


    return 0;
}

