func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums { 
		 freqMap[num]++
	}

	arr := make([][2]int, 0, len(freqMap))
	for num, freq := range freqMap { 
		arr = append(arr, [2]int{num, freq})
	}

	sort.Slice(arr, func(i,j int) bool { 
		return arr[i][1] > arr[j][1]
	})

	result := make([]int, k)
	for i := 0; i < k; i++ { 
		result[i] = arr[i][0]
	}

	return result

}
