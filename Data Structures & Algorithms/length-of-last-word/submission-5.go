func lengthOfLastWord(s string) int {
	r := len(s) - 1
	count := 0

	for r >= 0 && !isValidChar(rune(s[r])){
		r--
	} 

	for r >= 0 && isValidChar(rune(s[r])){
		count++
		r--
	}

	return count
}

func isValidChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
