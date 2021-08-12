package main

import "fmt"

func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := (low + high) >> 1
		if nums[mid] < nums[high] {
			high = mid
		} else if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high--
		}
	}
	return nums[low]
}

func main() {
	fmt.Println(findMin([]int{1, 3, 5, 0}))
}
