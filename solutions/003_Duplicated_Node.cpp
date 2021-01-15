#include<iostream>

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
    ListNode* removeDuplicateNodes(ListNode* head) {
        int memo[20001]={0};
        ListNode *tmp=new ListNode(0);
        tmp->next=head;
        while(tmp->next)
        {
            if(memo[tmp->next->val]==0){
                memo[tmp->next->val]=1;
                tmp=tmp->next;
            }
            else
                tmp->next=tmp->next->next;
        }
        return head;
    }
};

int
main()
{
    Solution solution = Solution();
    ListNode list = ListNode(1);
    solution.removeDuplicateNodes(&list);
    return 0;
}