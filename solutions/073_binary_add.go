package main

func addBinary(a string, b string) string {
	i, j := len(a), len(b)
	n, gap := 0, 0
	if i > j {
		n, gap = i, i-j
		a, b = b, a // 保证 a 小于 b 的长度
	} else {
		n, gap = j, j-i
	}

	ans := make([]byte, n+1)
	var c byte = 0
	for k := n - 1; k >= 0; k-- {
		if k-gap >= 0 { // a+b
			v1, v2 := a[k-gap]-'0', b[k]-'0'
			ans[k+1] = (v1+v2+c)%2 + '0'
			c = (v1 + v2 + c) / 2
		} else { // 补a的空位
			v := b[k] - '0'
			ans[k+1] = (v+c)%2 + '0'
			c = (v + c) / 2
		}
	}
	if c == 1 {
		ans[0] = '1'
		return string(ans)
	}
	return string(ans[1:])
}
