func findRedundantConnection(edges [][]int) []int {
    n := len(edges)
	par := make([]int, n + 1)
	rank := make([]int, n + 1)

	for i := 0; i <= n; i++ { 
		par[i] = i
		rank[i] = 1
	}

	var find func(int) int
	find = func(x int) int { 
		if par[x] != x { 
			par[x] = find(par[x])
		}

		return par[x]
	}

	union := func(x, y int) bool { 
		rootX, rootY := find(x), find(y)
		if rootX == rootY { 
			return false
		}

		if rank[rootX] < rank[rootY] { 
			rootX, rootY = rootY, rootX
		}

		par[rootY] = rootX
		rank[rootX] += rank[rootY]
		return true
	}

	for _, edge := range edges { 
		if !union(edge[0], edge[1]) { 
			return edge
		}
	}
	return []int{}
}
