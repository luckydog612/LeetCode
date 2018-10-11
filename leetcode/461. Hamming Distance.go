package LeetCode

func hammingDistance(x int, y int) int {
	var count int
	i := dtb(x)
	j := dtb(y)
	d := len(i) - len(j)
	if d >= 0 {
		c := make([]int, d)
		j = append(j, c...)
	} else {
		c := make([]int, -d)
		i = append(i, c...)
	}
	for m := 0; m < len(i); m++ {
		if i[m] != j[m] {
			count++
		}
	}
	return count
}

func dtb(a int) []int {
	var b []int
	var i int
	for {
		i = a % 2
		a = a / 2
		b = append(b, i)
		if a == 0 {
			break
		}
	}
	return b
}
