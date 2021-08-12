package main

import "fmt"

// 冒泡排序

// 选择排序

// 插入排序

// 归并排序

// 堆排序
func heapSort(nums []int) {
	// 自 from 向 end 的 调整节点使得成为从from开始成为最大堆
	var maxHeapify func(from, end int)
	maxHeapify = func(from, end int) {
		left := 2*from + 1
		if left > end {
			return
		}
		right := left + 1
		idx := left
		if right <= end && nums[left] < nums[right] {
			idx = right
		}
		if nums[from] > nums[idx] {
			return
		}
		nums[from], nums[idx] = nums[idx], nums[from] // swap
		maxHeapify(idx, end)
	}

	// 最大堆化
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		maxHeapify(i, n-1)
	}

	// 不断pop 丢弃至nums末尾
	for i := n - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		maxHeapify(0, i-1)
	}
}

// 快速排序
func quickSort(nums []int) {
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
}

func sortArray(nums []int) []int {
	heapSort(nums)
	return nums
}

func main() {
	nums := []int{5, 1, 1, 2, 0, 0, 2, 3, 10, 2, 1}
	sortArray(nums)
	fmt.Println(nums)
}
