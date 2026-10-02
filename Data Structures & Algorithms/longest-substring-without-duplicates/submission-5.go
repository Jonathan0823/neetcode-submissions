func lengthOfLongestSubstring(s string) int {
	seen := make(map[rune]int)
	l := 0

	res := 0
	for idx, val := range s { 
		if val, exist := seen[val]; exist { 
			l = max(l, val + 1)
		}

		res = max(res, idx - l + 1)
		seen[val] = idx
	}

	return res
}
