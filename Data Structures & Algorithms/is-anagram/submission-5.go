func isAnagram(s string, t string) bool {
    if len(s) != len(t) { 
        return false
    }

    var wCount [26]int
    for i:= 0; i < len(s); i++ { 
        wCount[s[i] - 'a']++
        wCount[t[i] - 'a']--
    }

    for _, v := range wCount { 
        if v != 0 { 
            return false
        }
    }

    return true

}
