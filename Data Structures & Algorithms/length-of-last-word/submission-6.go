func lengthOfLastWord(s string) int {
	count := 0
	for r := len(s) - 1; r >= 0; r-- { 
		if s[r] != ' ' { 
			count++
		} else if count > 0 {
			break
		} 
	}

	return count
}

