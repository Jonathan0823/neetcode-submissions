func lengthOfLongestSubstring(s string) int {
	substringLength:= 0
	seen := make(map[rune]int) 
	left := 0 
	for right:= 0; right < len(s); right++ { 
		if num, ok := seen[rune(s[right])]; ok { 
			left = max(left, num + 1)
		}
		substringLength = max(substringLength, (right - left + 1))

		seen[rune(s[right])] = right 
	}

	return substringLength
}
