func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int
	for i, n := range nums { 
		if i > 0 && nums[i-1] == n { 
			continue
		}

		l, r := i+1, len(nums) - 1
		for l < r { 
			sum := n + nums[l] + nums[r]
			if sum == 0 { 
				res = append(res, []int{n, nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] { 
					l++
				}
				for l < r && nums[r] == nums[r+1] { 
					r--
				}
				
			} else if sum > 0 { 
				r--
			} else { 
				l++
			}

		}
	}

	return res

}
