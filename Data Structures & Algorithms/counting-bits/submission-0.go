func countBits(n int) []int {
	res := make([]int, n + 1)
	for i := 1; i <= n; i++ { 
		num := i
		for num != 0 { 
			res[i]++
			num &= (num - 1)
		}
	}

	return res
}
