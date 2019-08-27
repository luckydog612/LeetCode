package main

func countBits(num int) []int {
	b := make([]int, 0)
	i := 0
	for i <= num {
		if i == 0 {
			b = append(b, 0)
		} else if i <= 2 {
			b = append(b, 1)
		} else {
			count := 0
			by := dtob(i)
			for _, j := range by {
				if j == 1 {
					count++
				}
			}
			b = append(b, count)
		}
		i++
	}
	return b
}

func dtob(num int) []int {
	a := make([]int, 0)
	for num != 0 {
		i := num % 2
		a = append(a, i)
		num = num / 2
	}
	return a
}
