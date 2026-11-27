func partitionLabels(s string) []int {
    lastIndex := make(map[rune]int) // O(1)
	for i, r := range s { 
		lastIndex[r] = i
	}

	var res []int
	size, end := 0, 0
	for i, r := range s { 
		size++
		end = max(end, lastIndex[r])
		if i == end { 
			res = append(res, size)
			size = 0
		}
	}

	return res
}
