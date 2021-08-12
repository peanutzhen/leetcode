package main

// 如果在遍历的过程中 有个更小的点 我们记录下来
// 因为如果后面有更大的点 则与这个点相减会比之前的差值更大
func maxProfit(prices []int) int {
	min := prices[0]
	rtv := 0
	n := len(prices)
	for i := 1; i < n; i++ {
		if prices[i]-min > rtv {
			rtv = prices[i] - min // 保存最大差值
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return rtv
}
