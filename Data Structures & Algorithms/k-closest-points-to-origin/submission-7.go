type MaxHeap [][]int

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i][2] > h[j][2] }
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any) { *h = append(*h, x.([]int))}
func (h *MaxHeap) Pop() any { 
	old := *h
	x := old[len(old) - 1]
	*h = old[:len(old) - 1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	for _, p := range points { 
		x, y := p[0], p[1]
		dist := x*x + y*y
		heap.Push(maxHeap, []int{x, y, dist})
		if maxHeap.Len() > k { 
			heap.Pop(maxHeap)
		}
	}

	result := make([][]int, k)
	for i := k-1; i >= 0; i-- { 
		point := heap.Pop(maxHeap).([]int)
		result[i] = []int{point[0], point[1]}
	}

	return result
}
