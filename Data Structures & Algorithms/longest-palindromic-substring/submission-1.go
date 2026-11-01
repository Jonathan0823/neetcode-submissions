func longestPalindrome(s string) string {
	var res string
	var countFromMiddle func(l, r int)
	countFromMiddle = func(l, r int) { 
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r - l + 1 > len(res) { 
				res = s[l:r+1]
			}
			l--
			r++
		}
	}

	for i, _ := range s { 
		countFromMiddle(i, i)
		countFromMiddle(i, i + 1)
	}

	return res
}
