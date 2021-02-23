package main

func isValid(s string) bool {
	var stack []rune
	push := func(word rune) {
		stack = append(stack, word)
	}
	pop := func() rune {
		lastIndex := len(stack) - 1
		word := stack[lastIndex]
		stack = stack[:lastIndex]
		return word
	}
	for _, word := range s {
		switch word {
		case ')':
			if len(stack) == 0 || pop() != '(' {
				return false
			}
		case '}':
			if len(stack) == 0 || pop() != '{' {
				return false
			}
		case ']':
			if len(stack) == 0 || pop() != '[' {
				return false
			}
		default:
			push(word)
		}
	}
	return len(stack) == 0
}

func main() {
	isValid(")")
}
