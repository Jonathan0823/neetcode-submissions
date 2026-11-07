func mergeAlternately(word1 string, word2 string) string {
	var sb strings.Builder
	i, j := 0, 0
	for i < len(word1) && j < len(word2) { 
		sb.WriteByte(word1[i])
		sb.WriteByte(word2[j])
		i++
		j++
	}

	res := word1[i:]
	if res == "" { 
		res = word2[j:]
	} 

	sb.WriteString(res)
	return sb.String()
}
