func isAnagram(s string, t string) bool {
    if len(s) != len(t) { 
        return false
    }

    count := make(map[rune]int)

    for i, char := range s { 
        count[char - 'a']++
        count[rune(t[i]) - 'a']--
    }

    for _, val := range count { 
        if val != 0 { 
            return false
        }
    }

    return true
}
