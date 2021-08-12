package main

// 注意 其他元素都是2次 只有一个元素1次
// 相同元素异或等于0 异或具有交换律
func singleNumber(nums []int) int {
	rtv := nums[0]
	for i := 1; i < len(nums); i++ {
		rtv = rtv ^ nums[i]
	}
	return rtv
}
