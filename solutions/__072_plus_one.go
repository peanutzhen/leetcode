package main

func plusOne(digits []int) []int {
	// 精巧的代码设计
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		digits[i]++
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}
	return append([]int{1}, digits...)

	// 之前写的屎代码
	/*
		n := len(digits)
		digits[n-1]++

		i, c := n-2, 1
		if digits[n-1] == 10 {
			digits[n-1] = 0
			for i >= 0 && c == 1 {
				digits[i]++
				if digits[i] != 10 {
					break
				}
				digits[i] = 0
				i--
			}
			if i == -1 {
				digits[0] = 0
				return append([]int{1}, digits...)
			}
		}

		return digits*/
}
