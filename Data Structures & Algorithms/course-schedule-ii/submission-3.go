func findOrder(numCourses int, prerequisites [][]int) []int {
    prereq := make(map[int][]int)
    for i := 0; i < numCourses; i++ {
        prereq[i] = []int{}
    }
    for _, p := range prerequisites {
        crs, pre := p[0], p[1]
        prereq[crs] = append(prereq[crs], pre)
    }

    output := []int{}
    visiting := make(map[int]bool)
    visited := make(map[int]bool) 

    var dfs func(int) bool
    dfs = func(crs int) bool {
        if visiting[crs] {
            return false
        }
        if visited[crs] {
            return true
        }

        visiting[crs] = true
        for _, p := range prereq[crs] {
            if !dfs(p) {
                return false
            }
        }

        visiting[crs] = false
        visited[crs] = true
        output = append(output, crs)
        return true
    }

    for i := 0; i < numCourses; i++ {
        if !dfs(i) {
            return []int{}
        }
    }

    return output
}