package main

func maxCount(m int, n int, ops [][]int) int {
	num_op := len(ops)
	if num_op == 0 {
		return m * n
	}
	min_a, min_b := m+1, n+1
	for i := 0; i < num_op; i++ {
		if ops[i][0] < min_a {
			min_a = ops[i][0]
		}
		if ops[i][1] < min_b {
			min_b = ops[i][1]
		}
	}
	return min_a * min_b
}
