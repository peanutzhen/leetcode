package main

// 使用哈希表 记录出现过的数字 这样能在O(1)时间复杂度内
// 判断出是否出现
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	hashmap := map[int]bool{}
	// 初始化 哈希表
	for i := 0; i < len(nums); i++ {
		hashmap[nums[i]] = true
	}
	rtv := 1 // 至少为1
	for i := 0; i < len(nums); i++ {
		// 该数字的上一个数字没有出现 我们现在开始计算它的最远连续长度
		if !hashmap[nums[i]-1] {
			tmp := 1
			value := nums[i] + 1
			for hashmap[value] {
				tmp++
				value++
			}
			if tmp > rtv {
				rtv = tmp
			}
		}
	}
	return rtv
}
