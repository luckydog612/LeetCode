package main

func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	flag := false
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
			flag = false
		}else {
			if !flag {
				i++
				nums[i] = nums[j]
				flag = true
			}
		}
	}
	return i + 1
}
