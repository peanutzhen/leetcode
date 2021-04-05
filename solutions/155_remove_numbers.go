package main

import "fmt"

func minLength(str string) int {
	stack := []byte{}
	n := len(str)
	for i := 0; i < n; i++ {
		if s := len(stack); s != 0 && str[i]+stack[s-1]-96 == 10 {
			stack = stack[:s-1]
		} else {
			stack = append(stack, str[i])
		}
	}
	return len(stack)
}

func main() {
	n := 0
	str := ""
	fmt.Scan(&n)
	fmt.Scan(&str)
	fmt.Println(minLength(str))
}
