func characterReplacement(s string, k int) int {
	var charCount [26]int

	l, res, maxF := 0, 0, 0

	for r := 0; r < len(s); r++ {
		charCount[s[r] - 'A']++
		maxF = max(maxF, charCount[s[r] - 'A'])

		for ((r-l+1) - maxF) > k { 
			charCount[s[l] - 'A']--
			l++
		}
		res = max(res, r-l+1)
	}

	return res

}
