func longestCommonPrefix(strs []string) string {
	commonPrefix := strs[0]

	for i := 1; i < len(strs); i++ {
		str := strs[i]
		for !strings.HasPrefix(str, commonPrefix) && len(commonPrefix) > 0 { 
			commonPrefix = commonPrefix[:len(commonPrefix) - 1]
		}
	}

	return commonPrefix
}
