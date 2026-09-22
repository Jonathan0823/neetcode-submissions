func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
    for idx, num := range nums { 
        completion := target - num
        if val, exist := seen[completion]; exist{ 
            return []int{val, idx}
        }

        seen[num] = idx
    }

    return []int{}
    
}
