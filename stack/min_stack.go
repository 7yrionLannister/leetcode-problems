package main

type MinStack struct {
	min     *int
	array   []int
	ordered []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.min == nil || val <= *this.min {
		this.min = &val
		this.ordered = append(this.ordered, val)
	}
	this.array = append(this.array, val)
}

func (this *MinStack) Pop() {
	top := this.Top()
	if top == *this.min {
		nO := len(this.ordered)
		var newMin *int
		if nO >= 2 {
			newMin = &this.ordered[nO-2]
		}
		this.ordered = this.ordered[:nO-1]
		this.min = newMin
	}
	this.array = this.array[:len(this.array)-1]
}

func (this *MinStack) Top() int {
	return this.array[len(this.array)-1]
}

func (this *MinStack) GetMin() int {
	if this.min == nil {
		return 0
	}
	return *this.min
}
