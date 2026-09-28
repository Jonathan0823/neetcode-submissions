func groupAnagrams(strs []string) [][]string {
	anagramsMap := make(map[[26]int][]string)
	for _, str := range strs { 
		var count [26]int
		for _, s := range str { 
			count[s - 'a']++
		}
		anagramsMap[count] = append(anagramsMap[count], str)
	}

	var res [][]string
	for _, val := range anagramsMap {
		res = append(res, val)
	}

	return res
}
