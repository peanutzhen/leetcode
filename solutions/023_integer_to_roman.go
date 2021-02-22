package main

// 非常辣眼睛 啊哈哈哈哈哈哈哈哈
func intToRoman(num int) string {
	table := [][]string{
		{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"M", "MM", "MMM"},
	}

	var tmp [4]int
	for i := 3; num > 0; i-- {
		tmp[i] = num % 10
		num /= 10
	}

	var rtv string
	for i := 0; i < 4; i++ {
		if tmp[i] != 0 {
			rtv += table[3-i][tmp[i]-1]
		}
	}

	return rtv
}
