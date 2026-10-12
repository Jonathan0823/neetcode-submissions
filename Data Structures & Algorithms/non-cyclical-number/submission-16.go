func isHappy(n int) bool {
	slow := sumSquareDigits(n)
	fast := sumSquareDigits(sumSquareDigits(n))

	for fast != 1 && slow != fast { 
		slow = sumSquareDigits(slow)
		fast = sumSquareDigits(sumSquareDigits(fast))
	}

	return fast == 1
}

func sumSquareDigits(n int) int { 
	res := 0
	for n > 0 { 
		res += (n % 10) * (n % 10)
		n /= 10
	}
	fmt.Println(res)

	return res
}
