package main

// 如果从某起点a出发 到某个点b走不动了 就尝试b的下一个节点c开始尝试
// 如果c没有下一节点了 就返回-1 否则返回那个可行的起点
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)

	var recursion func(size, idx, count int) int
	recursion = func(size, idx, count int) int {
		if count != n {
			size += gas[idx]
			if size-cost[idx] >= 0 {
				return recursion(size-cost[idx], (idx+1)%n, count+1)
			}
		}
		return count // 不够油到 返回走的步数 或 已经到了
	}

	rtv := -1
	start := 0
	for start < n {
		step := recursion(0, start, 0)
		if step == n {
			rtv = start
			break
		}
		start += step + 1 // 到不了 尝试下一节点开始
	}
	return rtv
}
