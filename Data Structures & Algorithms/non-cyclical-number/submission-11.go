func isHappy(n int) bool {
	slow, fast := n, countSumSquare(n)

	for slow != fast && fast != 1 {
		slow = countSumSquare(slow)
		fast = countSumSquare(countSumSquare(fast))
	}

	return fast == 1
}

func countSumSquare(n int) int { 
	res := 0
	for n > 0 { 
		digit := n%10
		res += digit * digit
		n /= 10
	}

	return res

}