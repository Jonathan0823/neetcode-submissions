type MaxHeap []int

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j]}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i]}
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

type pendingTask struct {
    freq        int
    availableAt int
}

func leastInterval(tasks []byte, n int) int {
	taskFreq := make(map[byte]int)
	for _, t := range tasks { 
		taskFreq[t]++
	}

	var h MaxHeap
	heap.Init(&h)
	for _, v := range taskFreq { 
		heap.Push(&h, v)
	}

	time := 0
	var queue []pendingTask

	for h.Len() > 0 || len(queue) > 0 { 
		time++

		if h.Len() > 0 { 
			freq := heap.Pop(&h).(int) - 1
			if freq > 0 { 
				queue = append(queue, pendingTask{
					freq: freq,
					availableAt: time+n,
				})
			}
		}

		if len(queue) > 0 && queue[0].availableAt == time { 
				heap.Push(&h, queue[0].freq)
				queue = queue[1:]
			}
	}

	return time
}
