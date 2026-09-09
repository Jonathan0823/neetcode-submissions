func characterReplacement(s string, k int) int {
	charFreq := make(map[byte]int)
	l, maxF, res := 0, 0, 0

	for r:= 0; r < len(s); r++ { 
		charFreq[s[r]]++

		maxF = max(maxF, charFreq[s[r]])

		for (r-l + 1) - maxF > k { 
			charFreq[s[l]]--
			l++
		}

		res = max(res, r-l + 1)
	}

	return res

}
