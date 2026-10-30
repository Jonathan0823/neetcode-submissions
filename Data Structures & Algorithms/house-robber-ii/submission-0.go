func rob(nums []int) int {

	return max(nums[0], helper(nums[1:]), helper(nums[:len(nums) - 1]))
}

func helper(nums []int) int { 
	rob1, rob2 := 0,0
	for _, n := range nums { 
		temp := max(n + rob1, rob2)
		rob1 = rob2
		rob2 = temp
	}

	return rob2
}
