package LeetCode

func sortArrayByParity(A []int) []int {
	var odd []int
	var even []int
	for _, i := range A {
		if i % 2 == 0 {
			even = append(even, i)
		}else{
			odd = append(odd, i)
		}
	}
	even = append(even, odd...)
	return even
}
