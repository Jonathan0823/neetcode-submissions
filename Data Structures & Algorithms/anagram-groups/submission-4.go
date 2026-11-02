func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[[26]int][]string)
	for _, s := range strs { 
		var key [26]int
		for _, c := range s { 
			key[c - 'a']++
		}

		anagramMap[key] = append(anagramMap[key], s)
	}

	var res [][]string
	for _, val := range anagramMap { 
		res = append(res, val)
	}

	return res
}
