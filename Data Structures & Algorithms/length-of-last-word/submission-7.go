func lengthOfLastWord(s string) int {
	count := 0
	for i := len(s) - 1; i >=0; i-- {
		if count != 0 && s[i] == ' '{ 
			return count
		}
		if s[i] != ' ' { 
			count++
		}
	}

	return count
}
