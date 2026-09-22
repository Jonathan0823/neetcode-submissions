func checkInclusion(s1 string, s2 string) bool {
	var s1ArrayFreq, windowFreq [26]int
	for i := 0; i< len(s1);i++ { 
		s1ArrayFreq[s1[i] - 'a']++
		windowFreq[s2[i] - 'a']++
	}

	if s1ArrayFreq == windowFreq { 
			return true
		}


	for r:= len(s1); r < len(s2); r++ { 
		windowFreq[s2[r]-'a']++
		windowFreq[s2[r-len(s1)] - 'a']--	
		

		if s1ArrayFreq == windowFreq { 
			return true
		}

	}

	return false

}
