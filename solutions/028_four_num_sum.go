package main

import "sort"

var rtv [][]int

// 稍微将三数之和算法变形就可用了
func threeSum(nums []int, target, w int) {
	length := len(nums)

	if length > 2 {
		target = target - w
		i := 0
		for i < length {
			j, k := i+1, length-1
			for j < k {
				if nums[j]+nums[k]+nums[i] == target {
					rtv = append(rtv, []int{w, nums[i], nums[j], nums[k]})
					tmp := nums[j]
					for j < length && nums[j] == tmp {
						j++
					}
				} else if nums[j]+nums[k]+nums[i] < target {
					j++
				} else if nums[j]+nums[k]+nums[i] > target {
					k--
				}
			}
			// 跳转至下一数
			tmp := nums[i]
			for i < length && nums[i] == tmp {
				i++
			}
		}
	}
}

func fourSum(nums []int, target int) [][]int {
	rtv = [][]int{} //清空缓存
	length := len(nums)
	if length > 3 {
		sort.Ints(nums)
		i := 0
		for i < length {
			threeSum(nums[i+1:], target, nums[i]) // x+y+z == target-w
			tmp := nums[i]
			for i < length && nums[i] == tmp {
				i++
			}
		}
	}
	return rtv
}

func main() {
	a := []int{5, 5, 3, 5, 1, -5, 1, -2}
	b := 4
	fourSum(a, b)
}
