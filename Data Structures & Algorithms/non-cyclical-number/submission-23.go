func isHappy(n int) bool {
	slow := countSquares(n)
	fast := countSquares(countSquares(n))
	
	for slow != fast && fast != 1 { 
		slow = countSquares(slow)
		fast = countSquares(countSquares(fast))
	}

	return fast == 1
}

func countSquares(n int) int { 
	res := 0
	for n > 0 { 
		res += (n % 10) * (n % 10)
		n = n/10
	}

	return res
}