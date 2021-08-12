package main

import "fmt"

// 代码相当易懂，不多说
func algo(s string) bool {
	length := len(s)
	for i := length - 1; i >= 0; i-- {
		expected_letter := byte('a' + i)
		if s[i] == expected_letter {
			s = s[:i]
		} else if s[0] == expected_letter {
			s = s[1:]
		} else {
			return false
		}
	}
	return true
}

func input() []string {
	num_case := 0
	fmt.Scan(&num_case)
	testcases := make([]string, num_case)

	for i := 0; i < num_case; i++ {
		fmt.Scan(&testcases[i])
	}

	return testcases
}

func main() {
	strs := input()

	for _, s := range strs {
		if algo(s) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
