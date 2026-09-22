func climbStairs(n int) int {
	if (n <= 2 && n >0) { 
		return n
	}
    
    return climbStairs(n-1) + climbStairs(n-2)
}
