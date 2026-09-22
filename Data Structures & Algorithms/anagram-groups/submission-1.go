func groupAnagrams(strs []string) [][]string {
	anagramsMap := make(map[[26]int][]string)
	for _, str := range strs { 
		var key [26]int
		for _, s := range str { 
			key[s - 'a']++
		}
		anagramsMap[key] = append(anagramsMap[key], str)
	}

	var result [][]string
	for _, val := range anagramsMap { 
		result = append(result, val)
	}

	return result

}
