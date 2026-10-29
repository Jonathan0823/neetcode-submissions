type MaxHeap []int

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any) { *h = append(*h , x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	x := old[len(old) - 1]
	*h = old[:len(old) - 1]
	return x
}

func lastStoneWeight(stones []int) int {
	h := MaxHeap(append([]int{}, stones...))
	maxHeap := &h
	heap.Init(maxHeap)

	for maxHeap.Len() > 1 { 
		first := heap.Pop(maxHeap).(int)
		second := heap.Pop(maxHeap).(int)
		sub := first - second
		if sub != 0 { 
			heap.Push(maxHeap, sub)
		}
		
	}

	if maxHeap.Len() == 0 { 
		return 0
	}

	return (*maxHeap)[0]
}
