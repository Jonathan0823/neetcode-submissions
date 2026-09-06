func lengthOfLongestSubstring(s string) int {
	substringLength:= 0
	seen := make(map[byte]int) 
	left := 0 
	for right:= 0; right < len(s); right++ { 
		if num, ok := seen[s[right]]; ok { 
			left = max(left, num + 1)
		}
		substringLength = max(substringLength, (right - left + 1))

		seen[s[right]] = right 
	}

	return substringLength
}
