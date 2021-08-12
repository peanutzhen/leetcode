package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	dist := make([]int, n+1)
	max_path_dist := 0
	sum := 0
	sum_dist := 0
	actual_dist := 0
	for n > 1 {
		var u, v, w int
		fmt.Scan(&u, &v, &w)
		dist[v] = dist[u] + w
		if dist[v] > max_path_dist {
			max_path_dist = dist[v]
		}
		sum_dist += dist[v]
		sum += w
		n--
	}
	actual_dist = sum*2 - max_path_dist
	fmt.Printf("%d %d\n", sum_dist, actual_dist)
}
