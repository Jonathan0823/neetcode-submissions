func hammingWeight(n int) int {
	res := 0
	for n != 0 { 
		if n&1 != 0 { 
			res++
		}
		n >>= 1
	}

	return res
}
