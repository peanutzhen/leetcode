# -*- coding: utf-8 -*-
class Solution:
    def nextPermutation(self, nums):
        """
        Do not return anything, modify nums in-place instead.

        我的做法是从倒数第二个元素N开始，找该元素右边比他大的元素K
        若这样的元素K存在多个，选取最小的那个
        若不存在这样的元素K，则N向左滑动，继续判断。

        找到了这样的N与K后，交换他们，之后对N右边的所有元素进行排序
        即得到正确答案
        """
        lastIndex = len(nums) - 1
        i = len(nums) - 2
        while i >= 0:
            N = nums[i]
            min_candidate = 101
            candidate_swap = -1
            for j in range(i+1, lastIndex+1):
                if N < nums[j]:
                    if nums[j] < min_candidate:
                        min_candidate = nums[j]
                        candidate_swap = j
                
            if candidate_swap != -1:
                # swap them
                tmp = nums[i]
                nums[i] = nums[candidate_swap]
                nums[candidate_swap] = tmp
                # sort
                nums[i+1:] = sorted(nums[i+1:])
                return
            else:
                i = i - 1
        nums[:] = sorted(nums)

if __name__ == "__main__":    
    test = Solution()
    nums = [2,3,1]
    test.nextPermutation(nums)
    print(nums)