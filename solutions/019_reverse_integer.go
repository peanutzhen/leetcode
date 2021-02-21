package main

import "fmt"

// 很暴力 谢谢
func reverse(x int) int {
	rtv := 0
	i := x
	if x < 0 {
		i = -x
	}
	for i > 0 {
		rtv = rtv*10 + i%10
		i /= 10
	}
	if x < 0 {
		rtv = -rtv
	}
	if rtv > 2147483647 || rtv < -2147483648 {
		return 0
	}
	return rtv
}

func main() {
	fmt.Printf("%d", reverse(2147483647))
}
