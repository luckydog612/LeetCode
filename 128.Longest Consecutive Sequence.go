package main

import (
	"fmt"
	"math"
	"sort"
)

/*
	Given an unsorted array of integers, find the length of the longest consecutive elements sequence.

	Your algorithm should run in O(n) complexity.

	Example:

		Input: [100, 4, 200, 1, 3, 2]
		Output: 4
	Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
*/

// time:12ms cache:3.1MB
func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	sort.Ints(nums)
	length := 1
	leng := 1
	first := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == first {
			continue
		}
		if nums[i] == first+1 {
			first = nums[i]
			leng++
		} else {
			if leng > length {
				length = leng
			}
			first = nums[i]
			leng = 1
		}
	}
	return int(math.Max(float64(length), float64(leng)))
}

// time:8ms cache:5MB
func longestConsecutive2(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	var length = 1
	var count int
	var onlyOne = make([]int, 0)
	var numMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, exist := numMap[nums[i]]; !exist {
			numMap[nums[i]] = 1
			onlyOne = append(onlyOne, nums[i])
		}
	}
	sort.Ints(onlyOne)
	var num = onlyOne[0]
	count++
	fmt.Println(onlyOne)
	for j := 1; j < len(onlyOne); j++ {
		if onlyOne[j] == num+1 {
			count++
		} else {
			if length < count {
				length = count
			}
			count = 1
		}
		num = onlyOne[j]
	}
	return int(math.Max(float64(length), float64(count)))
}

func main() {
	var nums = []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}
	length := longestConsecutive(nums)
	fmt.Println(length)
}
