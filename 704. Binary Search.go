package main

func search(nums []int, target int) int {
	var start = 0
	var end = len(nums) - 1
	for start <= end {
		mid := (start + end)/2
		if nums[mid] > target {
			end = mid -1
			continue
		}else if nums[mid]<target{
			start = mid +1
			continue
		}else{
			return mid
		}
	}
	return -1
}
