func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	freq := make([][]int, len(nums) + 1)

	for _, c := range nums { 
		count[c]++
	}
	for n, c := range count { 
		freq[c] = append(freq[c], n)
	}

	var res []int
	for i := len(freq) - 1; i > 0; i-- { 
			for _, v := range freq[i] { 
				res = append(res, v)
				if len(res) == k { 
					return res
				}
			}

	}

	return res
}
