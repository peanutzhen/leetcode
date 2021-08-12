package main

// 此题可以用堆或者双端队列来做
// 双端队列更快

// 无论用heap还是queue，检查max值是否在窗口中是关键
func maxSlidingWindow(nums []int, k int) []int {
	// 双向队列法
	// 队列存的是数组下标，下标递增
	// 下标对应元素递减
	queue := []int{}
	push := func(idx int) {
		// 保证元素递减
		for len(queue) != 0 && nums[queue[len(queue)-1]] < nums[idx] {
			queue = queue[:len(queue)-1] // 队尾出队
		}
		queue = append(queue, idx) // 队尾入队
	}

	// 初始化队列
	for i := 0; i < k; i++ {
		push(i)
	}
	rtv := []int{nums[queue[0]]}

	for i := k; i < len(nums); i++ {
		push(i)
		// 这个最大的元素并不在窗口内 移除
		for queue[0] < i-k+1 {
			queue = queue[1:] // 队首出队
		}
		rtv = append(rtv, nums[queue[0]])
	}
	return rtv
}
