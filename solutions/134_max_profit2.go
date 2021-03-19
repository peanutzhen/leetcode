package main

// 我们只需把所有上升的差值累加起来即可
func maxProfit(prices []int) int {
	in, prev := 0, 0
	prices = append(prices, -1) // 避免prices是递增 然后没法触发line 13的代码
	n := len(prices)

	profit := 0
	for i := 1; i < n; i++ {
		if prices[prev] > prices[i] {
			profit += prices[prev] - prices[in]
			in = i
		}
		prev = i
	}
	return profit
}
