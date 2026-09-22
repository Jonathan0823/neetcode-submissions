func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers) - 1

	for l < r { 
		complement = l + r
		if complement == target { 
			return []int{l, r}
		} else if complement > 0 { 
			r--
		} else { 
			l++
		}
	}

	return []int{}

}
