package main

import "sort"

// Combination 是状态节点
type Combination struct {
	Record []int // 一种可行的组合
	Ans    int   // 组合累加和
}

// 和044一样 改一点代码就行啦 关键在于避免重复答案
// 总的来说，先对candidates排序，然后重复的数字在根节点只尝试最开头数字一次
// 和三数之和那波有点像
func combinationSum2(candidates []int, target int) [][]int {
	rtv := [][]int{} // 返回值
	combination := Combination{[]int{}, 0}
	sort.Ints(candidates)        // 先排序，这样能跳过重复的数字
	var backtrace func(int)      // 回朔算法
	backtrace = func(from int) { // 算法定义 i从from开始避免重复答案
		for i := from; i < len(candidates); {
			value := candidates[i]
			if combination.Ans+value == target { // 保存该组合
				combination.Record = append(combination.Record, value)
				rtv = append(rtv, append([]int{}, combination.Record...))
				combination.Record = combination.Record[:len(combination.Record)-1]
			} else if combination.Ans+value < target {
				combination.Record = append(combination.Record, value)
				combination.Ans += value
				backtrace(i + 1) // 数字只能使用一次 所以为 i + 1
				combination.Record = combination.Record[:len(combination.Record)-1]
				combination.Ans -= value
			}
			for i < len(candidates) && candidates[i] == value {
				i++ // 这里我们要跳过重复的数字 这样就能避免重复答案
			}
		}
	}
	backtrace(0)
	return rtv
}
