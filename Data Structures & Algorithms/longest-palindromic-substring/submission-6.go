func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, maxLen := 0, 1

	countFromMiddle := func(l, r int) { 
		for l >= 0 && r < len(s) && s[l] == s[r] {
			currentLen := r - l + 1
			if currentLen > maxLen { 
				start = l
				maxLen = currentLen
			}
			l--
			r++
		}
	}

	for i := range s { 
		countFromMiddle(i, i)
		countFromMiddle(i, i + 1)
	}

	return s[start : start+maxLen]
}
