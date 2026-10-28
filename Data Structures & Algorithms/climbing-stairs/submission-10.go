func climbStairs(n int) int {
	one, two := 1, 2
	if n < 2 { 
		return n
	}
	for i := 3; i <= n; i++ { 
		temp := two
		two += one
		one = temp
	}

	return two
}
