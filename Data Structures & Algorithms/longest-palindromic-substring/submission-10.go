func longestPalindrome(s string) string {
	var count int
	var res string
    
	var checkFromMiddle func(l, r int) 
	checkFromMiddle = func(l, r int) { 
		for l >= 0 && r < len(s) && s[l] == s[r] { 
			length := r - l + 1
			if length > count { 
				res = s[l:r + 1]
				count = length
			}
			l--
			r++
		}

		return
	}

	for i := range s { 
		checkFromMiddle(i, i)
		checkFromMiddle(i, i + 1)
	}

	return res
}
