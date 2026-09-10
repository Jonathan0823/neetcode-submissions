func threeSum(nums []int) [][]int {
	var result [][]int

	// sort the slice
	sort.Slice(nums, func(i, j int) bool { 
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums)- 2; i++ { 
		if nums[i] > 0 { 
			break
		}

		// skip duplicate first value
		if i > 0 && nums[i] == nums[i - 1] { 
			continue
		}

		l, r := i+1, len(nums) - 1
		for l < r {		
			target := nums[l] + nums[r] + nums[i]
			if target == 0 { 
				result = append(result, []int{nums[i], nums[l], nums[r]})
				
				for l < r && nums[l] == nums[l+1]  { 
					l++
				}
				for l < r && nums[r] == nums[r -1 ] { 
					r--
				}
				l++
				r--
			} else if target > 0 { 
				r--
			} else { 
				l++
			}
		}
	}

	return result

}
