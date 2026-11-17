func validTree(n int, edges [][]int) bool {
    if len(edges) > n -1 { 
		return false
	}

	adj := make([][]int, n)
	for _, e := range edges { 
		u, v := e[0], e[1]
		adj[u]  = append(adj[u], v)
		adj[v]  = append(adj[v], u)
	}

	visit := make(map[int]bool)
	var dfs func(cur, prev int) bool
	dfs = func(cur, prev int) bool { 
		if visit[cur] { 
			return false
		}
		visit[cur] = true
		for _, n := range adj[cur] { 
			if n == prev { 
				continue
			}

			if !dfs(n, cur) { 
				return false
			}
		}
		return true
	}

	return dfs(0, -1) && len(visit) == n
}
