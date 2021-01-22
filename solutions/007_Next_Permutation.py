class Solution:
    def nextPermutation(self, nums):
        """
        Do not return anything, modify nums in-place instead.
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