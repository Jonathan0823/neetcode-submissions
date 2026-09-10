func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers) - 1

	for l < r { 
		complement := numbers[l] + numbers[r]
		if complement == target { 
			return []int{l+1, r+1}
		} else if complement > target { 
			r--
		} else { 
			l++
		}
	}

	return []int{}

}
