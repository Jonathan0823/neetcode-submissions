func climbStairs(n int) int {
	first, second := 1, 2
	if n < 3 { 
		return n
	}

	for i := 3; i <= n; i++ { 
		sum := first + second
		first = second
		second = sum
	}

	return second
}
