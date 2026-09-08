func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) < 3 {
		return result
	}
	sort.Slice(nums, func(i, j int) bool { 
		return nums[i] <  nums[j]	})

	
	for i := 0; i < len(nums) - 2;i++ { 
		if nums[i] > 0 { 
			break
		}

		if i > 0 && nums[i] == nums[i - 1] { 
			continue
		}

		seen := make(map[int]struct{})
		j := i + 1
		for j < len(nums) { 
			target := -1 * (nums[i] + nums[j])
			if _, exist := seen[target]; exist { 
				result = append(result, []int{nums[i], target, nums[j]})

				for j+1 < len(nums) && nums[j] == nums[j + 1] { 
					j++
				}
			}

			seen[nums[j]] = struct{}{}
			j++
		}
	}
	

	return result

}
