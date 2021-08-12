package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 皮皮怪写法
	//copy(nums1[m:], nums2)
	//sort.Ints(nums1)

	// O(m+n)写法
	nums1Copy := make([]int, m)
	copy(nums1Copy, nums1[:m])
	i, j := 0, 0
	k := 0
	for i < m && j < n {
		if nums1Copy[i] < nums2[j] {
			nums1[k] = nums1Copy[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}
	for i < m {
		nums1[k] = nums1Copy[i]
		i, k = i+1, k+1
	}
	for j < n {
		nums1[k] = nums2[j]
		j, k = j+1, k+1
	}
}
