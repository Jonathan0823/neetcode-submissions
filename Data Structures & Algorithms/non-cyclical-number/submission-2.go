func isHappy(n int) bool {
    seen := make(map[int]bool)
	for n > 1 {
		n = sumSquare(n) 
		if seen[n] { 
			return false
		}
		seen[n] = true
	}
	return true
}

func sumSquare(n int) int { 
	sum := 0
	for n > 0 { 
		digit := n % 10
		sum += digit * digit
		n = n/10
	}
	return sum
}
