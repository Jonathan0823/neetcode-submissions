func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]int)
	maxLength := 0
	l, r := 0, 0

	for r < len(s) { 		
		if idx, exist := seen[s[r]]; exist { 
			l = max(l, idx+1)
		}
		seen[s[r]] = r
		maxLength = max(maxLength, r-l+1)
		r++
	}

	return maxLength

}
