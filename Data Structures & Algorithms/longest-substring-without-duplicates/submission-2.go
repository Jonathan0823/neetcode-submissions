func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]int)
	maxLength := 0
	l, r := 0, 0

	for r < len(s) { 
		maxLength = max(maxLength, r-l)
		if idx, exist := seen[s[r]]; exist { 
			l = max(l, idx+1)
		}
		seen[s[r]] = r
		r++
	}

	return maxLength

}
