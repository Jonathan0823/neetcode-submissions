func characterReplacement(s string, k int) int {
	var charCount [26]int

	l, r, maxF := 0, 0, 0
	res := 0
	for r < len(s) {
		charCount[s[r] - 'A']++
		maxF = max(maxF, charCount[s[r] - 'A'])

		count := r - l + 1
		if (count - maxF) <= k { 
			res = max(res, count)
		} else { 
			charCount[s[l] - 'A']--
			l++
		}
		r++
	}

	return res

}
