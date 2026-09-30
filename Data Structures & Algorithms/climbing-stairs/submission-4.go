func climbStairs(n int) int {
	first, second := 1, 2
	if n < 3 { 
		return n
	}

	for i := 3; i <= n; i++ { 
		res := first + second
		first = second
		second = res
	}
	return second
}
