#include <stdio.h>

void
swap(int *nums, int i, int j)
{
	int tmp = nums[i];
	nums[i] = nums[j];
	nums[j] = tmp;
}

int
median(int *nums, int start, int end)
{
	int mid = (start + end) >> 1;
	if (nums[start] < nums[mid]) {
		swap(nums, start, mid);
	}
	if (nums[mid] < nums[end]) {
		swap(nums, mid, end);
	}
	if (nums[start] < nums[mid]) {
		swap(nums, start, mid);
	}
	swap(nums, mid, end-1);
	return nums[end-1];
}

int
partition(int *nums, int start, int end)
{
	if (start == end) {
		return start;
	}
	int pivot = median(nums, start, end);
	int i = start;
	int j = end - 1;
	while (i < j) {
		while (nums[++i] > pivot) {}
		while (nums[--j] < pivot) {}
		if (i < j) {
			swap(nums, i, j);
		}
	}
	swap(nums, i, end-1);
	return i;
}

int
findKthLargest(int* nums, int numsSize, int k)
{
	int start = 0, end = numsSize-1;
	while (1) {
		int cur = partition(nums, start, end);
		if (cur+1 == k) {
			return nums[cur];
		}
		else if (cur+1 > k) {
			end = cur - 1;
		}
		else {
			start = cur + 1;
		}
	}
}


// 测试
int
main()
{
	int nums[] = {1};
	printf("%d\n", findKthLargest(nums,sizeof(nums)/sizeof(int),1));
	return 0;
}
