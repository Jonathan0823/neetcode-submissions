func climbStairs(n int) int {
	if n < 3 { 
		return n
	}
    first, second := 1, 2 
	for i := 3; i <= n; i++ { 
		temp := first + second
		first, second = second, temp
	}
	return second
}
