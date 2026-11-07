func mergeAlternately(word1 string, word2 string) string {
	var sb strings.Builder
	sb.Grow(len(word1) + len(word2))
	i, j := 0, 0
	for i < len(word1) && j < len(word2) { 
		sb.WriteByte(word1[i])
		sb.WriteByte(word2[j])
		i++
		j++
	}

	sb.WriteString(word1[i:])
	sb.WriteString(word2[j:])

	return sb.String()
}
