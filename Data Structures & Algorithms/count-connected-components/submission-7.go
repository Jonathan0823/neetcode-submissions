func countComponents(n int, edges [][]int) int {
	adj := make([][]int, n)
	visit := make([]bool, n)

	for _, e := range edges { 
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	var dfs func(int)
	dfs = func(cur int) { 
		for _, nei := range adj[cur] { 
			if !visit[nei] { 
				visit[nei] = true
				dfs(nei)
			}
		}
	}

	res := 0
	for i := 0; i < n; i++ { 
		if !visit[i] { 
			visit[i] = true
			res++
			dfs(i)
		}
	}

	return res
}
