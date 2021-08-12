package main

import "fmt"

func lengthOfLastWord(s string) int {
	n := len(s)
	rtv := 0
	flag := false
	for i := n - 1; i >= 0; i-- {
		if s[i] != ' ' {
			flag = true
		}
		if flag {
			if s[i] == ' ' {
				break
			}
			rtv++
		}
	}
	return rtv
}

func main() {
	fmt.Println(lengthOfLastWord("asdjkas a "))
}
