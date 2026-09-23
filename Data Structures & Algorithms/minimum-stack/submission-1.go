type MinStack struct {
	Val []int
	Min []int

}

func Constructor() MinStack {
	return MinStack{
		Val: []int{},
		Min: []int{},
	}

}

func (this *MinStack) Push(val int) {
	this.Val = append(this.Val, val)
	topMin := val
	if len(this.Min) > 0 { 
		topMin = this.Min[len(this.Min) - 1]
	}
	this.Min = append(this.Min, min(topMin, val))
}

func (this *MinStack) Pop() {
	this.Val = this.Val[:len(this.Val) - 1]
	this.Min = this.Min[:len(this.Min) - 1]
}

func (this *MinStack) Top() int {
	top := 0
	if len(this.Val) > 0 { 
		top = this.Val[len(this.Val)-1]
	}
	return top
}

func (this *MinStack) GetMin() int {
	top := 0
	if len(this.Min) > 0 { 
		top = this.Min[len(this.Min)-1]
	}
	return top
}
