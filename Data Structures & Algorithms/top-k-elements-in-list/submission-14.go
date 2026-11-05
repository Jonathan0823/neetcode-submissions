func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int) // number:count
	for _, n := range nums { 
		countMap[n]++
	}

	buckets := make([][]int, len(nums) + 1) // count:[]int{numbers...}
	for n, cnt := range countMap { 
		buckets[cnt] = append(buckets[cnt], n)
	}

	var res []int
	for i := len(nums); i >= 0; i-- { 
		for j := 0; j < len(buckets[i]); j++ { 
			res = append(res, buckets[i][j])
			if len(res) == k { 
				return res
			}
		}
	}

	return res
}
