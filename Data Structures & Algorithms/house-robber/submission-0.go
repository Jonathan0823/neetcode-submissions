func rob(nums []int) int {
	first, second := 0, 0

	for _, n := range nums { 
		temp := max(n + first, second)
		first = second
		second = temp
	}
	return second
}
