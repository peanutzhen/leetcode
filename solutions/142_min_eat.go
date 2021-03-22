/*
	牛牛想要吃饱，吃饱至少要m美味值
	现在给n个菜，每道菜有一美味值a。
	问牛牛最少吃多少菜可以吃饱？
	返回该数并返回对应的选择方案
*/

package main

import (
	"fmt"
	"sort"
)

// 将美味值从大到小排好
// 并连同他们的下标也排好
// 贪心算法即可
func minEating(nums []int, m int) (int, []int) {
	n := len(nums)
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = i
	}
	tmp := make([]int, n)
	copy(tmp, nums)
	// 插入排序
	// 注意：这里不能
	// sort.Slice(idx, func(i, j int) {
	//	 return nums[i] > nums[j]
	// })
	// 没有用的这样，草泥马golang
	for i := 0; i < n-1; i++ {
		maxValue := -1
		maxIndex := -1
		for j := i + 1; j < n; j++ {
			if nums[j] > maxValue {
				maxValue = nums[j]
				maxIndex = j
			}
		}
		if nums[i] < nums[maxIndex] {
			idx[i], idx[maxIndex] = idx[maxIndex], idx[i]
			nums[i], nums[maxIndex] = nums[maxIndex], nums[i]
		}
	}

	count := 0
	choice := []int{}
	copy(nums, tmp)
	for count < n && m > 0 {
		m -= nums[idx[count]]
		choice = append(choice, idx[count]+1)
		count++
	}
	if m > 0 {
		count = -1
		choice = []int{}
	}
	sort.Ints(choice)
	return count, choice
}

func main() {
	loop := 0
	fmt.Scan(&loop)

	// 获取测试用例
	testcases := make([][]int, loop)
	for i := 0; i < loop; i++ {
		n, m := 0, 0
		fmt.Scan(&n, &m)
		testcases[i] = make([]int, n+2)
		testcases[i][0], testcases[i][1] = n, m
		for j := 2; j < n+2; j++ {
			fmt.Scan(&testcases[i][j])
		}
	}

	// 测试
	for i := 0; i < loop; i++ {
		count, choice := minEating(testcases[i][2:], testcases[i][1])
		fmt.Println(count)
		if count != -1 {
			for j := 0; j < count; j++ {
				fmt.Printf("%d ", choice[j])
			}
		}
		fmt.Printf("\n")
	}
}
