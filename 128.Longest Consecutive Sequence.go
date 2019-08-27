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

func main() {
	var nums = []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}
	length := longestConsecutive(nums)
	fmt.Println(length)
}
