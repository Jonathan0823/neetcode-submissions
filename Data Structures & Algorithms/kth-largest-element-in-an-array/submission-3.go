type MinHeap []int

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i]}
func (h *MinHeap) Push(x any) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any { 
	old := *h
	x := old[len(old) - 1]
	*h = old[:len(old) - 1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := MinHeap{}
	minHeap := &h
	heap.Init(minHeap)
	for _, n := range nums { 
		heap.Push(minHeap, n)
		if minHeap.Len() > k { 
			heap.Pop(minHeap)
		}
	}

	return (*minHeap)[0]
}
