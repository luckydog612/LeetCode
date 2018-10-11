package LeetCode

import "sort"

func thirdMax(nums []int) int {
	n := make([]int, 0)
	sort.Ints(nums)
	i := 0
	n = append(n, nums[0])
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			n = append(n, nums[j])
			i++
			nums[i] = nums[j]
		}
	}
	ln := len(n)
	if ln <= 2 {
		return n[ln-1]
	}
	return n[ln-3]
}
