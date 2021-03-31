def minEmployee(nums):
	n = len(nums)
	nums.sort()
	rtv = 0

	i, j = 0, 0
	while i < n:
		val = nums[i]
		while j < n and nums[j] == val and j-i-1 < val:
			j = j + 1
		rtv = rtv + val + 1
		i = j
	
	return rtv

if __name__ == "__main__":
	str = input()
	nums = eval(str)
	print(minEmployee(nums))