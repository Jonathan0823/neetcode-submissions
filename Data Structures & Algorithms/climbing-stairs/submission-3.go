func climbStairs(n int) int {
	if n < 3 { 
		return n
	}
    first, second := 1, 2
	for i:= 3; i <= n; i++ { 
		next := first + second
		first, second = second, next
	}

	return second

}
