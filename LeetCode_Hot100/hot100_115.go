package main

type MinStack struct {
	data    []int
	minData []int
	length  int
}

func Constructor() MinStack {
	return MinStack{
		data:    make([]int, 0),
		minData: make([]int, 0),
		length:  0,
	}
}

func (this *MinStack) Push(val int) {
	if this.length != 0 {
		if val < this.minData[this.length-1] {
			this.minData = append(this.minData, val)
		} else {
			this.minData = append(this.minData, this.minData[this.length-1])
		}
	} else {
		this.minData = append(this.minData, val)
	}
	this.data = append(this.data, val)
	this.length += 1
}

func (this *MinStack) Pop() {
	this.data = this.data[0 : this.length-1]
	this.minData = this.minData[0 : this.length-1]
	this.length -= 1
}

func (this *MinStack) Top() int {
	return this.data[this.length-1]
}

func (this *MinStack) GetMin() int {
	return this.minData[this.length-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
