#include <stdio.h>

int search(int* nums, int numsSize, int target){
    /* O(n)时间复杂度
    for(int i=0;i<numsSize;i++)
    {
        if(nums[i]==target)
            return i;
    }
    return -1;
    */
    // O(logn)解法

    int front = 0;
    int rear = numsSize - 1;
    int mid = (front + rear) / 2;
    
    while(front <= rear)
    {
        if(nums[mid]==target)
            return mid;
        else if(nums[mid] < target)
        {
            // 右边如果没有断点且rear小于target
            // 请去左边找而不是右边
            if(nums[mid+1] <= nums[rear] && target > nums[rear])
            {
                rear = mid - 1;
                mid = (front + rear) / 2;
            }
            else
            {
                front = mid + 1;
                mid = (front + rear) / 2;
            }
            
        }
        else if(target < nums[mid])
        {
            // 左边如果没有断点且front大于target
            // 请去右边找而不是左边
            if(nums[front] <= nums[mid - 1] && target < nums[front])
            {
                front = mid + 1;
                mid = (front + rear) / 2;
            }
            else
            {
                rear = mid - 1;
                mid = (front + rear) / 2;
            }
        }
    }
    return -1;
}

int
main()
{
    int nums[] = {3,5,7,8,10,11,12,-10,-9,-7};
    int numsSize = sizeof(nums)/sizeof(int);
    int target;
    for(int i=0;i<numsSize;i++)
        printf("%d\n", search(nums, numsSize, nums[i]));
    
    printf("%d\n", search(nums, numsSize, 0));
    return 0;
}