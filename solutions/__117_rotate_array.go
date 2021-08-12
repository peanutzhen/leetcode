package main

// T:O(n) S:O(n)
func method1(nums []int, k int) {
	n := len(nums)
	tmp := append([]int{}, nums[n-k:]...)
	copy(nums[k:], nums[:n-k])
	copy(nums[:k], tmp)
}

// T:O(kn) S:O(1)
func method2(nums []int, k int) {
	n := len(nums)
	for i := 0; i < k; i++ {
		lastOne := nums[n-1]
		for j := n - 2; j >= 0; j-- {
			nums[j+1] = nums[j]
		}
		nums[0] = lastOne
	}
}

// 还有一个翻转法 类似于
/*
	nums = "----->-->"; k =3
	result = "-->----->";

	reverse "----->-->" we can get "<--<-----"
	reverse "<--" we can get "--><-----"
	reverse "<-----" we can get "-->----->"
	this visualization help me figure it out :)
*/
func rotate(nums []int, k int) {
	method1(nums, k%len(nums))
}
