func rob(nums []int) int {
    return max(nums[0], helper(nums[1:]), helper(nums[:len(nums) - 1]))
}

func helper(nums []int) int { 
	first, second := 0, 0

	for _, n := range nums { 
		temp := max(first + n, second)
		first = second
		second = temp
	}

	return second
}
