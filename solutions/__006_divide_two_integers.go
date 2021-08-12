package main

import "fmt"

func divide(dividend int, divisor int) int {
	// 特殊情形
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend > -2147483648 {
			return -dividend
		}
		return 2147483647
	}

	sign := (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0)
	// 正数改成负数避免了溢出讨论
	if dividend > 0 {
		dividend = -dividend
	}
	if divisor > 0 {
		divisor = -divisor
	}

	if sign {
		return div(dividend, divisor)
	}
	return -div(dividend, divisor)
}

// logN的获得除法解
// 请保证a与b都为负数
func div(a, b int) int {
	if a > b {
		return 0
	}
	times, tmpB := 1, b
	for tmpB+tmpB >= a {
		times, tmpB = times+times, tmpB+tmpB
	}
	return times + div(a-tmpB, b)
}

// test
func main() {
	a, b := -2147483647, -1
	fmt.Println("divide():", divide(a, b))
	fmt.Println("golang:", a/b)
}
