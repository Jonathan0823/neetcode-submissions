func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i,j int) bool { 
		return nums[i] < nums[j]
	})

	var result [][]int

	for i := 0; i < len(nums) - 2; i++ { 
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums) - 1
		for l < r { 
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 { 
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for nums[l] == nums[l-1] && l < r {
					l++
				}
				for nums[r] == nums[r+1] && l < r {
					r--
				}

			} else if sum < 0 { 
				l++

			} else { 
				r--
				
			}
		}
	}

	return result
}
