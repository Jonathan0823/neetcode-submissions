func countSubstrings(s string) int {
	res := 0

	for i, _ := range s { 
		res += expandFromMiddle(s, i, i)
		if i < len(s) - 1 {
			res += expandFromMiddle(s, i, i + 1)
		}
	}

	return res
}

func expandFromMiddle(s string, l, r int) int { 
	cnt := 0

	for l >= 0 && r < len(s) && s[l] == s[r] { 
		cnt++
		l--
		r++
	}

	return cnt
}
