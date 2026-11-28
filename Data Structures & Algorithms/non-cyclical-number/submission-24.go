func isHappy(n int) bool {
	slow, fast := sumSquaresDigits(n), sumSquaresDigits(sumSquaresDigits(n))
	for slow != fast { 
		slow = sumSquaresDigits(slow)
		fast = sumSquaresDigits(sumSquaresDigits(fast))
	}

	return fast == 1
}

func sumSquaresDigits(n int) int { 
	res := 0
	for n > 0  {
		res += (n % 10) * (n % 10)
		n /= 10
	}
	return res
}
