package main

// 用 0 补长较小数字
func fillZero(s string, n int) string {
	tmp := make([]byte, n)
	for i := 0; i < n; i++ {
		tmp[i] = '0'
	}
	tmp = append(tmp, []byte(s)...)
	return string(tmp)
}

func addStrings(num1 string, num2 string) string {
	len1, len2 := len(num1), len(num2)
	var rtvLen int
	var idx int
	if len1 > len2 {
		rtvLen, idx = len1+1, len1-1
		num2 = fillZero(num2, len1-len2)
	} else {
		rtvLen, idx = len2+1, len2-1
		num1 = fillZero(num1, len2-len1)
	}

	rtv := make([]byte, rtvLen)

	carry := false
	for idx >= 0 {
		tmp := num1[idx] + num2[idx] - '0' - '0'
		if carry {
			tmp++
			carry = false
		}
		if tmp > 9 {
			carry = true
			tmp -= 10
		}
		rtvLen--
		rtv[rtvLen] = tmp + '0'
		idx--
	}
	if carry {
		rtv[rtvLen-1] = '1'
		return string(rtv)
	}
	return string(rtv[1:])
}
