package LeetCode

func flipAndInvertImage(A [][]int) [][]int {
	for _, j := range A {
		i := fzs(j)
		for m := range i {
			if i[m] == 0 {
				i[m] = 1
			} else {
				i[m] = 0
			}
		}
	}
	return A
}

func fzs(a []int) []int {
	start := 0
	end := len(a) - 1
	for start <= end {
		a[start], a[end] = a[end], a[start]
		start++
		end--
	}
	return a
}
