package main

func jump(nums []int) int {
	count := 0
	m := len(nums)
	// 倒遍历 找出能够到达lastIdx的最小idx
	idx := m - 1
	for idx > 0 {
		lastIdx := idx
		for i := idx - 1; i >= 0; i-- {
			if i+nums[i] >= lastIdx {
				idx = i
			}
		}
		count++
	}
	return count
}
