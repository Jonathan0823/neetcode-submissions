type MinStack struct {
	Val []int
	Min []int

}

func Constructor() MinStack {
	return MinStack { 
		Val: []int{},
		Min: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.Val = append(this.Val, val)
	minVal := val
	if len(this.Min) > 0 && this.Min[len(this.Min) - 1] < minVal { 
		minVal = this.Min[len(this.Min) - 1]
	}
	this.Min = append(this.Min, minVal)
}

func (this *MinStack) Pop() {
	this.Val = this.Val[:len(this.Val) - 1]
	this.Min = this.Min[:len(this.Min) - 1]
}

func (this *MinStack) Top() int {
	return this.Val[len(this.Val) - 1]
}

func (this *MinStack) GetMin() int {
	return this.Min[len(this.Min) - 1]
}
