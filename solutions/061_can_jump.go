package main

func canJump(nums []int) bool {
	idx := len(nums) - 1
	for i := idx - 1; i >= 0; i-- {
		if i+nums[i] >= idx {
			idx = i // 只要能跳 就等价于能否跳到idx这个点
		}
	}
	return idx == 0
}
