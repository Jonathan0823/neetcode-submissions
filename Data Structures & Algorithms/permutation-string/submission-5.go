func checkInclusion(s1 string, s2 string) bool {
	var l1, l2 [26]int

	if len(s2) < len(s1) { 
		return false
	}

	for i, s := range s1 { 
		 l1[s - 'a']++
		 l2[s2[i] - 'a']++ 
	}

	if l1 == l2 { 
		return true
	}

	for i := len(s1); i < len(s2); i++ { 	
		l2[s2[i - len(s1)] - 'a']--
		l2[s2[i] - 'a']++
		if l1 == l2 { 
			return true
		}
	}

	return false
}
