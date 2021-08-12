package main

import "fmt"

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// 当F在A与B中间时，答案加2即可
// Tips: 使用异或判断F是否处于AB之间
func algo(testcases [][3][2]int) {
	for _, test_case := range testcases {
		row_diff := abs(test_case[0][0] - test_case[1][0])
		col_diff := abs(test_case[0][1] - test_case[1][1])
		ans := row_diff + col_diff

		if row_diff == 0 && test_case[2][0] == test_case[0][0] {
			a := test_case[0][1] - test_case[2][1]
			b := test_case[1][1] - test_case[2][1]
			if a^b < 0 {
				ans += 2
			}
		} else if col_diff == 0 && test_case[2][1] == test_case[0][1] {
			a := test_case[0][0] - test_case[2][0]
			b := test_case[1][0] - test_case[2][0]
			if a^b < 0 {
				ans += 2
			}
		}
		fmt.Println(ans)
	}
}

func input() [][3][2]int {
	num_case := 0
	fmt.Scanln(&num_case)
	testcases := make([][3][2]int, num_case)

	for i := 0; i < num_case; i++ {
		fmt.Scanf("\n")
		for j := 0; j < 3; j++ {
			fmt.Scanf("%d %d\n", &testcases[i][j][0], &testcases[i][j][1])
		}
	}

	return testcases
}

func main() {
	testcases := input()
	algo(testcases)
}
