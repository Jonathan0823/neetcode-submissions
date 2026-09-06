func isHappy(n int) bool {
	seen := make(map[int]struct{})
	for n > 1 {
		n = calculateSumSquare(n)
		if _, ok := seen[n]; ok { 
			return false
		}
		seen[n] = struct{}{}
	}
	return true
    
}

func calculateSumSquare(n int) int {
	var answer int
	for n > 0 {
		answer += int(math.Pow(float64(n%10), 2))
		n = n / 10
	}
	return answer
}