package main

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	// m位*n位的长度至多为m+n位
	len1, len2 := len(num1), len(num2)
	ans := make([]byte, len1+len2)
	for i := 0; i < len(ans); i++ {
		ans[i] = '0' // 初始化
	}

	// num1[i] * num2[j] 影响 ans[i+j+1] 位
	for i := len1 - 1; i >= 0; i-- {
		for j := len2 - 1; j >= 0; j-- {
			tmp := (num1[i] - '0') * (num2[j] - '0')
			a, c := tmp%10, tmp/10
			// a位置加法 即i+j+1处
			if ans[i+j+1]+a > '9' {
				c++
				ans[i+j+1] += a - 10
			} else {
				ans[i+j+1] += a
			}
			// c位置加法 即i+j处
			if ans[i+j]+c > '9' {
				ans[i+j-1]++
				ans[i+j] += c - 10
			} else {
				ans[i+j] += c
			}
		}
	}
	if ans[0] == '0' {
		return string(ans[1:])
	}
	return string(ans)
}
