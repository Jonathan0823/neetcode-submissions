func longestPalindrome(s string) string {
	var start, end int
	var countFromMiddle func(l, r int)
	countFromMiddle = func(l, r int) { 
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r - l + 1 > end - start { 
				start = l
				end = r + 1
			}
			l--
			r++
		}
	}

	for i, _ := range s { 
		countFromMiddle(i, i)
		countFromMiddle(i, i + 1)
	}

	return s[start:end]
}
