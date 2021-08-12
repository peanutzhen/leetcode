package main

// 非常蛋疼的代码
// 据说解法是DFA(有穷自动机)
// 以后再改善
func isInteger(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	if n == 1 {
		return s[0] >= '0' && s[0] <= '9'
	}
	if s[0] != '+' && s[0] != '-' && (s[0] < '0' || s[0] > '9') {
		return false
	}
	for i := 1; i < n; i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func isFloat(s string) bool {
	pIdx, n := -1, len(s)
	for i := 0; i < n; i++ {
		if s[i] == '.' {
			pIdx = i
		}
	}
	if pIdx == -1 {
		return false
	}
	if pIdx == 0 {
		return s[1:] != "" && (s[1] != '-' && s[1] != '+') && isInteger(s[1:])
	}
	if pIdx == 1 {
		if s[0] == '+' || s[0] == '-' {
			return isInteger(s[pIdx+1:])
		}
		return (s[0] >= '0' && s[0] <= '9') && (s[pIdx+1:] == "" || isInteger(s[pIdx+1:]))
	}
	return isInteger(s[:pIdx]) && (s[pIdx+1:] == "" || (s[pIdx+1] != '-' && s[pIdx+1] != '+' && isInteger(s[pIdx+1:])))
}

func isNumber(s string) bool {
	eIdx, n := -1, len(s)
	for i := 0; i < n; i++ {
		if s[i] == 'e' || s[i] == 'E' {
			eIdx = i
		}
	}
	if eIdx == -1 {
		return isFloat(s) || isInteger(s)
	}
	return (isFloat(s[:eIdx]) || isInteger(s[:eIdx])) && isInteger(s[eIdx+1:])
}
