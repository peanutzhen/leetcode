package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minStep(s1, s2 string, start int) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	hashmap := map[byte][]int{}
	for i := 0; i < m; i++ {
		hashmap[s1[i]] = append(hashmap[s1[i]], i)
	}

	var recursion func(begin int, j int) int
	recursion = func(begin, j int) int {
		if j == n {
			return 0
		}
		if dp[begin][j] == 0 {
			minVal := 2147483647
			for i := 0; i < len(hashmap[s2[j]]); i++ {
				dist := hashmap[s2[j]][i] - begin
				if begin < hashmap[s2[j]][i] {
					dist = min(dist, m-dist)
				} else {
					dist = min(-dist, m+dist)
				}
				if dist > minVal { //剪枝
					continue
				}
				tmp := dist + recursion(hashmap[s2[j]][i], j+1)
				if tmp < minVal {
					minVal = tmp
				}
			}
			dp[begin][j] = minVal
		}
		return dp[begin][j]
	}
	return recursion(start, 0)
}

func main() {
	/*
		s1, s2 := "", ""
		start := 0
		fmt.Scan(&s1)
		fmt.Scan(&s2)
		fmt.Scan(&start)
		fmt.Println(minStep(s1, s2, start))
	*/
	fmt.Println(minStep("aemoyn", "amo", 0))
}
