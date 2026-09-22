func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	numFreq := make([][]int, len(nums)+1)

	for _, num := range nums {
		count[num]++
	}

	for num, freq := range count { 
		numFreq[freq] = append(numFreq[freq], num)
	}

	result := []int{}
	for i := len(numFreq) - 1; i >= 0; i-- { 
		for _, num := range numFreq[i] { 
			result = append(result, num)
			if len(result) == k { 
				return result
			}
		}

	}

	return result

}
