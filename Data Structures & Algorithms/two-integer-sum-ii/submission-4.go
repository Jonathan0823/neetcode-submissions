func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers) - 1

	for l < r { 
		complement := numbers[l] + numbers[r]
		if complement == target { 
			return []int{numbers[l], numbers[r]}
		} else if complement > target { 
			r--
		} else { 
			l++
		}
	}

	return []int{}

}
