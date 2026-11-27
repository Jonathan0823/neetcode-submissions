func longestPalindrome(s string) string {
    var res string
	var maxLen int

	var checkFromMiddle func(l, r int)
	checkFromMiddle = func(l, r int) { 
		for l >= 0 && r < len(s) && s[l] == s[r] { 
			length := r - l + 1
			if length > maxLen { 
				res = s[l:r+1]
				maxLen = length
			}
			l--
			r++
		}
		return
	}

	for i := range s { 
	 	checkFromMiddle(i, i)
	 	checkFromMiddle(i, i+1)
	}

	return res
}
