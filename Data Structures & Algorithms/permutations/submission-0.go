func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	perms := permute(nums[1:])
	var res [][]int
	for _, p := range perms { 
		for i := 0; i <= len(p); i++ { 
			temp := make([]int, len(p))
			copy(temp, p)
			temp = append(temp[:i], append([]int{nums[0]}, temp[i:]...)...)
			res = append(res, temp)
		}
	}


	return res
}
