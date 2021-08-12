#include <stdio.h>
#include <stdlib.h>

struct ListNode
{
    int val;
    struct ListNode *next;
};

struct ListNode*
addTwoNumbers(struct ListNode* l1, struct ListNode* l2)
{
    struct ListNode *l3 = (struct ListNode*)malloc(sizeof(struct ListNode));
    l3->next = NULL;
    struct ListNode *rtv = l3;
    int cflag = 0;  // 进位标志

    // 个位单独处理
    l3->val = l1->val + l2->val;
    // 进位处理
    if(l3->val > 9)
    {
        l3->val -= 10;
        cflag = 1;
    }
    l1 = l1->next;
    l2 = l2->next;

    // 后面位处理
    while(l1 || l2 || cflag)
    {
        // 分配空间
        l3->next = (struct ListNode*)malloc(sizeof(struct ListNode));
        l3 = l3->next;
        l3->val = 0;
        l3->next = NULL;
        // 相加
        if(l1 && l2)
            l3->val = l1->val + l2->val;
        else if(!l1 && l2)
            l3->val = l2->val;
        else if(l1 && !l2)
            l3->val = l1->val;
        // 考虑进位
        if(cflag)
        {
            l3->val += 1;
            cflag = 0;
        }
        // 进位处理
        if(l3->val > 9)
        {
            l3->val -= 10;
            cflag = 1;
        }
        // 移动指针
        if(l1)
            l1 = l1->next;
        if(l2)
            l2 = l2->next;
    }
    return rtv;
}

int
main()
{
    struct ListNode *l1 = (struct ListNode*)malloc(sizeof(struct ListNode));
    struct ListNode *l2 = (struct ListNode*)malloc(sizeof(struct ListNode));

    l1->val = 8;
    l2->val = 9;
    l1->next = NULL;
    l2->next = NULL;
    struct ListNode *l3 = addTwoNumbers(l1,l2);
    while(l3)
    {
        printf("%d",l3->val);
        l3 = l3->next;
    }
    return 0;
}