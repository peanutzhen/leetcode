package main

func searchRange(nums []int, target int) []int {
	rtv := []int{binarySearch(nums, target), binarySearch(nums, target)}
	if rtv[0] == -1 {
		return rtv
	}
	// 一直找 直到找到左边那个和target不等那个
	for rtv[0] > 0 && nums[rtv[0]-1] == nums[rtv[0]] {
		rtv[0] = binarySearch(nums[0:rtv[0]], target)
	}
	// 一直找 直到找到右边那个和target不等那个
	for rtv[1] < len(nums)-1 && nums[rtv[1]+1] == nums[rtv[1]] {
		rtv[1] += binarySearch(nums[rtv[1]+1:], target) + 1
	}
	return rtv
}

func binarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high + low) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1 // -1表示没找着
}
