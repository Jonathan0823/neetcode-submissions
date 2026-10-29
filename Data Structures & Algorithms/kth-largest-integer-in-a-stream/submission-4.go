type MinHeap []int

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)  { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any { 
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[:n - 1]
	return x
}

type KthLargest struct {
	heap *MinHeap
	k int
}


func Constructor(k int, nums []int) KthLargest {
	h := MinHeap(append([]int{}, nums...))
	heap.Init(&h)
	for h.Len() > k { 
		heap.Pop(&h)
	}

	return KthLargest{ 
		heap: &h,
		k: k,
	}
}


func (this *KthLargest) Add(val int) int {
	heap.Push(this.heap, val)
	if this.heap.Len() > this.k { 
		heap.Pop(this.heap)
	}

	return (*this.heap)[0]
}
