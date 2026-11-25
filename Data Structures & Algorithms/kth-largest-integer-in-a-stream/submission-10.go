type MinHeap []int
func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any { 
	old := *h
	x := old[len(old) - 1]
	*h = old[:len(old) - 1]
	return x
}

type KthLargest struct {
    k int
	minHeap *MinHeap
}


func Constructor(k int, nums []int) KthLargest {
	h := MinHeap(append([]int{}, nums...))
	heap.Init(&h)
	for h.Len() > k { 
		heap.Pop(&h)
	}
    return KthLargest{
		k: k,
		minHeap: &h,
	}
}


func (this *KthLargest) Add(val int) int {
    heap.Push(this.minHeap, val)
	if this.minHeap.Len() > this.k { 
		heap.Pop(this.minHeap)
	}

	return (*this.minHeap)[0]
}
