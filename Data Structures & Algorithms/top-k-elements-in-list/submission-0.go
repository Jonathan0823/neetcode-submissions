func topKFrequent(nums []int, k int) []int {
	seen := make(map[int]int)
for _, num := range nums { 
	seen[num] += 1;
}
var result []int

for key, val := range seen { 
	if val >= k { 
		result = append(result, key)
	}
}

return result
}
