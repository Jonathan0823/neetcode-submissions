func validTree(n int, edges [][]int) bool {
    if len(edges) > n - 1 {
		return false
	}

	adj := make([][]int, n)
	for _, e := range edges { 
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	visited := make(map[int]bool)
	var dfs func(i, prev int) bool
	dfs = func(i, prev int) bool { 
		if visited[i] { 
			return false
		}

		visited[i] = true
		for _, n := range adj[i] { 
			if n == prev { 
				continue
			}

			if !dfs(n, i) { 
				return false
			}
		}

		return true
	}

	return dfs(0, -1) && len(visited) == n
}
