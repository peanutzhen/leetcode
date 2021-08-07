package main

import "fmt"

func main() {
	var s1, s2, s3, s4, s5 int
	fmt.Scan(&s1, &s2, &s3, &s4, &s5)
	total := s1 + s2 + s3 + s4 + s5
	if total == 0 {
		fmt.Println("0.0")
		return
	}
	score := float64(s1*1+s2*2+s3*3+s4*4+s5*5) / float64(total)
	// 利用整型截断float小数位
	fmt.Printf("%.1f\n", float64(int(score*10))/10)
}
