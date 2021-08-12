package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	var value, expense int
	for n > 0 {
		var x, y int
		fmt.Scan(&x, &y)
		if x > y {
			value += x
			expense += x - y
		} else {
			value += y
		}
		n--
	}
	fmt.Printf("%d %d\n", value, expense)
}
