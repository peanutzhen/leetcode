package main

import "fmt"

// 冒泡排序

// 选择排序

// 插入排序

// 归并排序

// 快速排序
func quickSort(nums []int) []int {
	// 三元素取中值法
	var median func(low, high int) int
	median = func(low, high int) int {
		mid := (low + high) >> 1
		if nums[low] > nums[mid] {
			nums[low], nums[mid] = nums[mid], nums[low]
		}
		if nums[mid] > nums[high] {
			nums[mid], nums[high] = nums[high], nums[mid]
		}
		if nums[low] > nums[mid] {
			nums[low], nums[mid] = nums[mid], nums[low]
		}
		nums[mid], nums[high-1] = nums[high-1], nums[mid]
		return nums[high-1]
	}

	// 划分
	var partition func(low, high int) int
	partition = func(low, high int) int {
		pivot := median(low, high)
		i, j := low, high-1
		for i < j {
			for {
				i++
				if nums[i] >= pivot {
					break
				}
			}
			for {
				j--
				if nums[j] <= pivot {
					break
				}
			}
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		nums[i], nums[high-1] = nums[high-1], nums[i]
		return i
	}

	// 快速排序
	var qsort func(i, j int)
	qsort = func(i, j int) {
		if i < j {
			p := partition(i, j)
			qsort(i, p-1)
			qsort(p+1, j)
		}
	}

	qsort(0, len(nums)-1) // in-place
	return nums
}

func sortArray(nums []int) []int {
	return quickSort(nums)
}

func main() {
	nums := []int{5, 1, 1, 2, 0, 0, 2, 3, 10, 2, 1}
	sortArray(nums)
	fmt.Println(nums)
}
