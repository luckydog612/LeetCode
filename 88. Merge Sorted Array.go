package main

func merge(nums1 []int, m int, nums2 []int, n int)  {
	if m == 0 {
		copy(nums1, nums2)
	}
	if n == 0 {
		return
	}
	copy(nums1[m:], nums2)
	for i := 0; i < m+n; i++ {
		for j := 0; j < m+n-1-i; j++ {
			if nums1[j] > nums1[j+1] {
				nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
			}
		}
	}
}
