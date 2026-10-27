func letterCombinations(digits string) []string {
	res := []string{}
	combination := []rune{}

	if len(digits) == 0 { 
			return []string{}
	}

	phoneMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var dfs func(i int) 
	dfs = func(i int) { 
		if i == len(digits) { 
			res = append(res, string(combination))
			return
		}

		letters := phoneMap[digits[i]]
		for _, l := range letters { 
			combination = append(combination, l)
			dfs(i + 1)
			combination = combination[:len(combination) - 1]
		}
	} 

	dfs(0)

	return res
}
