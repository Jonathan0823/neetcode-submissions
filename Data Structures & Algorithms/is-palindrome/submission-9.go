func isPalindrome(s string) bool {
	l, r := 0, len(s) - 1
	for l < r { 
		for l < r && !isAlphanumeric(rune(s[l])) { 
			l++
		}
		for l < r && !isAlphanumeric(rune(s[r])) { 
			r--
		}
		
		if unicode.ToLower(rune(s[l]))!= unicode.ToLower(rune(s[r])) { 
			return false
		} 			
		l++
			r--
		
	}

	return true

}

func isAlphanumeric(s rune) bool { 
	return unicode.IsLetter(s) || unicode.IsDigit(s)
}
