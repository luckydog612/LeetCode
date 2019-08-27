package main

func findKthLargest(nums []int, k int) int {
	bubleSort(nums)
	return nums[len(nums)-k]
}

func bubleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
