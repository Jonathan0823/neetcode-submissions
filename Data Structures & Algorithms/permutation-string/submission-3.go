func checkInclusion(s1 string, s2 string) bool {
	var firstCharCount [26]int
	var secondCharCount [26]int

	if len(s2) < len(s1) { 
		return false
	}

	for i := 0; i < len(s1); i++  { 
		firstCharCount[s1[i] - 'a']++
		secondCharCount[s2[i] - 'a']++
	}

	if firstCharCount == secondCharCount { 
		return true
	}

	for i := len(s1); i < len(s2); i++ { 
		secondCharCount[s2[i] - 'a']++
		secondCharCount[s2[i - len(s1)] - 'a']--

		if firstCharCount == secondCharCount { 
			return true
		}
	}

	return false

}
