type Point struct { 
	x, y int
	dist int
}

type MaxHeap []Point

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].dist > h[j].dist }
func (h MaxHeap) Swap(i, j int)  { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any) { 
	*h = append(*h, x.(Point))
}
func (h *MaxHeap) Pop() any { 
	old := *h
	length := len(old)
	x := old[length-1]
	*h = old[:length-1]

	return x
}

func kClosest(points [][]int, k int) [][]int {
	h := &MaxHeap{}
	heap.Init(h)
	for _, p := range points { 
		dist := p[0] * p[0] + p[1] * p[1]
		pt := Point{x:p[0], y:p[1], dist: dist}
		if h.Len() < k { 
			heap.Push(h, pt)
		} else if dist < (*h)[0].dist { 
			heap.Pop(h)
			heap.Push(h, pt)
		}
	}

	result := make([][]int, k)
	for i, pt := range *h { 
		result[i] = []int{pt.x, pt.y}
	}

	return result
}
