package main

import "fmt"

// 回朔算法 尝试所有可能即可
func checkValidString(s string) bool {
	stack := []rune{}
	push := func(word rune) {
		stack = append(stack, word)
	}
	pop := func() rune {
		lastIndex := len(stack) - 1
		if lastIndex == -1 {
			return 0
		}
		word := stack[lastIndex]
		stack = stack[:lastIndex]
		return word
	}

	var backtrace func(string) bool
	backtrace = func(s string) bool {
		if len(s) == 0 {
			return len(stack) == 0
		}
		word := s[0]
		if word == '(' {
			push('(')
			return backtrace(s[1:])
		} else if word == ')' {
			tmp := pop()
			if tmp != '(' {
				return false
			}
			return backtrace(s[1:])
		} else {
			curStack := append([]rune{}, stack...) // 保存当前堆栈
			// * 的 ( 尝试
			push('(')
			if backtrace(s[1:]) {
				return true
			}
			stack = append([]rune{}, curStack...) // 恢复堆栈
			// * 的 ) 尝试
			tmp := pop()
			if tmp == '(' && backtrace(s[1:]) {
				return true
			}
			stack = append([]rune{}, curStack...) // 恢复堆栈
			return backtrace(s[1:])
		}
	}
	return backtrace(s)
}

// test
func main() {
	fmt.Println(checkValidString("(((*())()*")) // ((()())()) true
	fmt.Println(checkValidString("(((**))"))    //(((-))) true
}
