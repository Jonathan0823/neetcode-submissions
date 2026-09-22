func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    var stringCount [26]int

    for i:=0; i < len(s); i++ { 
        stringCount[s[i] - 'a']++
         stringCount[t[i] - 'a']--
    }

    for _, val := range stringCount { 
        if val != 0 { 
            return false
        }
    }

    return true

}
