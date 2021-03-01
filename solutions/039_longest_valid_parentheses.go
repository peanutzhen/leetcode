package main

// 看官方题解才会做，吗的
func longestValidParentheses(s string) int {
	stack := []int{-1}
	rtv := 0
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[0 : len(stack)-1] // pop
			if len(stack) != 0 {
				length := i - stack[len(stack)-1]
				if length > rtv {
					rtv = length
				}
			} else {
				stack = append(stack, i)
			}
		}
	}
	return rtv
}
