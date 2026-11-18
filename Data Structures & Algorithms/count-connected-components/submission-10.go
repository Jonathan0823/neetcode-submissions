func countComponents(n int, edges [][]int) int {
	parent := make([]int, n)
	size := make([]int, n)

	for i := 0; i < n; i++ { 
		parent[i] = i
		size[i] = 1
	}

	var find func(int) int
	find = func(x int) int { 
		if parent[x] != x { 
			parent[x] = find(parent[x]) 
		}
		return parent[x]
	}

	components := n
	for _, e := range edges { 
		rootU := find(e[0])
		rootV := find(e[1])

		if rootU != rootV { 
			if size[rootU] < size[rootV] { 
				rootU, rootV = rootV, rootU
			}
			parent[rootV] = rootU
			size[rootU] += size[rootV]
			components--
		}
	}

	return components
}
