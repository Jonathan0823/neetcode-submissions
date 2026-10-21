func lengthOfLongestSubstring(s string) int {
	res := 0
	sMap := make(map[byte]int)
	l := 0
	for r := 0; r < len(s); r++ {
		if idx, exist := sMap[s[r]]; exist { 
			l = max(l, idx + 1)
		}

		sMap[s[r]] = r
		res = max(res, r - l + 1)
	}

	return res

}
