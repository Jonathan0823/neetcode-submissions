func checkInclusion(s1 string, s2 string) bool {
	var s1Key, windowKey [26]int

	if len(s2) < len(s1) { 
		return false
	}
	
	for i := 0; i < len(s1); i++ { 
		s1Key[s1[i] - 'a']++
		windowKey[s2[i] - 'a']++
	}

	if s1Key == windowKey { 
		return true
	}

	for i := len(s1); i < len(s2); i++ { 
		windowKey[s2[i] - 'a']++
		windowKey[s2[i - len(s1)] - 'a']--

		if s1Key == windowKey { 
			return true
		}
	}

	return false
}
