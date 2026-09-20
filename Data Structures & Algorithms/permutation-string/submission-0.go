func checkInclusion(s1 string, s2 string) bool {
	var s1ArrayFreq [26]int
	for _, s := range s1 { 
		s1ArrayFreq[s - 'a']++
	}

	for r:= len(s1) - 1; r < len(s2); r++ { 
		var windowFreq [26]int
		for i:= r-(len(s1)-1); i<=r;i++ { 
			windowFreq[s2[i] - 'a']++
		}

		if s1ArrayFreq == windowFreq { 
			return true
		}

	}

	return false

}
