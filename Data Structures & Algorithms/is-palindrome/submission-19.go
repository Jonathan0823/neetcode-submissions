func isPalindrome(s string) bool {
	l, r := 0, len(s) - 1
	for l < r { 
		for l < r && !isAlphaNumeric(rune(s[l])) { 
			l++
		}

		for l < r && !isAlphaNumeric(rune(s[r])) { 
			r--
		}

		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) { 
			return false
		}
		l++
		r--
	}

	return true
}

func isAlphaNumeric(c rune) bool { 
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}
