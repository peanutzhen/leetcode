package main

import "fmt"

// 深度优先搜索？我也不知道
func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	start := 0
	end := len(nums) - 1

	// 注意判定条件 因为使用mid指针判断是否存在 则start==end也有可能
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return true
		}
		if nums[start] == nums[mid] { // 无法判断哪部份有序
			start++ // 跳过这个冲突项
		} else if nums[start] < nums[mid] { // 说明前半部分有序
			if nums[start] <= target && target < nums[mid] {
				end = mid - 1 // 去前半部分找
			} else {
				start = mid + 1 // 去后半部分找
			}
		} else { // 说明后半部分有序
			if nums[mid] < target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return false
	// 老代码
	// var recursion func(low, high int) bool
	// recursion = func(low, high int) bool {
	// 	if low > high {
	// 		return false
	// 	}
	// 	mid := (low + high) / 2
	// 	if nums[mid] == target {
	// 		return true
	// 	}
	// 	return recursion(low, mid-1) || recursion(mid+1, high)
	// }
	// return recursion(0, len(nums)-1)
}

func main() {
	fmt.Println(search([]int{1, 0, 1, 1, 1}, 0))
}
