func lengthOfLongestSubstring(s string) int {
	seen := make(map[rune]int)

	var res, l int
	for r, val := range s { 
		if val, ok := seen[val]; ok { 
			l = max(l, val + 1)
		}
		seen[val] = r
		res = max(res, r - l + 1)
	}

	return res
}
