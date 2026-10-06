func longestCommonPrefix(strs []string) string {
	commonPrefix := strs[0]

	for _, s := range strs { 
		for !strings.HasPrefix(s, commonPrefix) { 
			commonPrefix = commonPrefix[:len(commonPrefix) - 1]
		}
	}

	return commonPrefix
}
