package main

// O(n) O(1) 好难想到
func removeDuplicates(nums []int) int {
	j, count := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			count = 1
		}
		if count <= 2 {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
