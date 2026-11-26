type MaxHeap [][]int
func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool { 
	distI := h[i][0] * h[i][0] + h[i][1] * h[i][1]
	distJ := h[j][0] * h[j][0] + h[j][1] * h[j][1]
	return distI > distJ
	}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any) { *h = append(*h, x.([]int))}
func (h *MaxHeap) Pop() any { 
	old := *h
	x := old[len(old) - 1]
	*h = old[:len(old) - 1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	h := &MaxHeap{}
	heap.Init(h)
	for _, p := range points { 
		heap.Push(h, p)
		if h.Len() > k { 
			heap.Pop(h)
		}
	}

	return [][]int(*h)
}
