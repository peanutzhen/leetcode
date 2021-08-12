package main

// ignore 为忽略nums[ignore]这项元素
// 用于查找nums是否存在target元素
func findNum(nums []int, target, ignore int) int {
	var rtv = -1
	for i := 0; i < len(nums); i++ {
		if i == ignore {
			continue
		}
		if nums[i] == target {
			rtv = i
			break
		}
	}
	return rtv
}

func twoSum(nums []int, target int) []int {
	rtv := []int{0, 0}
	for i := 0; i < len(nums); i++ {
		rtv[1] = findNum(nums, target-nums[i], i)
		if rtv[1] != -1 {
			rtv[0] = i
			break
		}
	}
	return rtv
}
