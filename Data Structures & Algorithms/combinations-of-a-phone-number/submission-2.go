func letterCombinations(digits string) []string {
	if len(digits) == 0 { 
		return []string{}
	}
	res := []string{}
	stack := []rune{}

	phoneMap := map[byte]string { 
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}	

	var dfs func(int)
	dfs = func(i int) { 
		if i == len(digits) { 
			res = append(res, string(stack))
			return
		}

		letter := phoneMap[digits[i]]
		for _, l := range letter { 
			stack = append(stack, l)
			dfs(i + 1)
			stack = stack[:len(stack) - 1]
		}
	}

	dfs(0)
	return res
}
