type MaxHeap []int
func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any) { 
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() any { 
	old := *h
	length := len(old)
	x := old[length - 1]
	*h = old[:length - 1]
	return x
}

func lastStoneWeight(stones []int) int {
	var maxHeap MaxHeap
	heap.Init(&maxHeap)
	for _, s := range stones { 
		heap.Push(&maxHeap, s)
	}

	for maxHeap.Len() > 1 { 
		first := heap.Pop(&maxHeap).(int)
		second := heap.Pop(&maxHeap).(int)
		heap.Push(&maxHeap, first - second)
	}

	return maxHeap[0]

}
