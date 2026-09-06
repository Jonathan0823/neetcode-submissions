func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

    for idx, num := range nums { 
        completion := target - num;

        if v, ok := seen[completion]; ok { 
            return []int{v, idx};
        }

        seen[num] = idx
    }

    return []int{}
    
}
