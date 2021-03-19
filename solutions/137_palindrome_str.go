package main

func equal(c1, c2 byte) bool {
	// Upper to lower
	if c1 < 'a' {
		c1 += 32
	}
	if c2 < 'a' {
		c2 += 32
	}
	return c1 == c2
}

func isInvalid(c byte) bool {
	if ('0' <= c && c <= '9') || ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z') {
		return false
	}
	return true
}

func isPalindrome(s string) bool {
	n := len(s)
	head, tail := 0, n-1
	for {
		for head < n && isInvalid(s[head]) {
			head++
		}
		for tail >= 0 && isInvalid(s[tail]) {
			tail--
		}
		if head > tail {
			break
		}
		if !equal(s[head], s[tail]) {
			return false
		}
		head++
		tail--
	}
	return true
}
