class Solution:
    def removeDuplicates(self, nums):
        i = 0
        while i < len(nums):
            if i == 0:
                lastNum = nums[i]
            else:
                if lastNum != nums[i]:
                    lastNum = nums[i]
                else:
                    nums.pop(i)
                    continue
            i = i + 1
        return len(nums)


if __name__ == "__main__":    
    test = Solution()
    nums = [0,0,1,1,1,2,2,3,3,4]
    ans = test.removeDuplicates(nums)
    print(nums)
