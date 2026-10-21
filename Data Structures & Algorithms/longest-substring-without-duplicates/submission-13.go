func lengthOfLongestSubstring(s string) int {
	res := 0
	sMap := make(map[rune]int)
	l := 0
	for idx, val := range s { 
		if i, exist := sMap[val]; exist { 
			l = max(l, i + 1)
		}

		res = max(res, idx - l + 1)
		
		sMap[val] = idx
	}

	return res

}
