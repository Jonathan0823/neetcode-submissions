func isAnagram(s string, t string) bool {
    if len(s) != len(t) { 
        return false
    }

    var count [26]int

    for i, char := range s { 
        count[char - 'a']++
        count[t[i] - 'a']--
    }

    for _, val := range count { 
        if val != 0 { 
            return false
        }
    }

    return true
}
