func groupAnagrams(strs []string) [][]string {
    anagramMap := make(map[[26]int][]string)
    var res [][]string

    for _, str := range strs { 
        var count [26]int
        for _, s := range str { 
            count[s - 'a']++
        }
        anagramMap[count] = append(anagramMap[count], str)
    }

    for _, val := range anagramMap { 
        res = append(res, val)
    }
    return res
}
