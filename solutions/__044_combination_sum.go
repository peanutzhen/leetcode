package main

// Combination 是状态节点
type Combination struct {
	Record []int // 一种可行的组合
	Ans    int   // 组合累加和
}

func combinationSum(candidates []int, target int) [][]int {
	rtv := [][]int{} // 返回值
	combination := Combination{[]int{}, 0}
	var backtrace func(int)      // 回朔算法
	backtrace = func(from int) { // 算法定义 i从from开始避免重复答案
		for i := from; i < len(candidates); i++ {
			value := candidates[i]
			if combination.Ans+value == target { // 保存该组合
				combination.Record = append(combination.Record, value)
				rtv = append(rtv, append([]int{}, combination.Record...))
				combination.Record = combination.Record[:len(combination.Record)-1]
			} else if combination.Ans+value < target {
				combination.Record = append(combination.Record, value)
				combination.Ans += value
				backtrace(i)
				combination.Record = combination.Record[:len(combination.Record)-1]
				combination.Ans -= value
			}
		}
	}
	backtrace(0)
	return rtv
}
