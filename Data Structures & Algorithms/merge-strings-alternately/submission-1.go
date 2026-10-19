func mergeAlternately(word1 string, word2 string) string {
	var res strings.Builder
	
	l := max(len(word1), len(word2))
	for i := 0; i < l; i++ { 
		if i < len(word1) { 
			res.WriteByte(word1[i])
		}
		if i < len(word2) { 
			res.WriteByte(word2[i])
		}
	}

	return res.String()
}
