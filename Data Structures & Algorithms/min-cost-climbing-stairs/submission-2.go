func minCostClimbingStairs(cost []int) int {
    first, second := 0, 0

    for i := 2; i <= len(cost); i++ {
        current := min(first+cost[i-2], second+cost[i-1])
        first = second 
        second = current
    }

    return second
}