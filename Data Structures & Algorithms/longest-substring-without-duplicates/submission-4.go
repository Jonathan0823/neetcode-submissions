func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]int)
	longestSubstring := 0

	l := 0
	for r := 0; r < len(s); r++ { 
		if idx, exist := seen[s[r]]; exist { 
			l = max(l, idx + 1)
		}

		longestSubstring = max(longestSubstring, r-l+1)

		seen[s[r]] = r
	} 

	return longestSubstring

}
