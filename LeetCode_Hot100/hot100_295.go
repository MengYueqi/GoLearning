package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type MedianFinder struct {
	arrMin hp // 小根堆
	arrMax hp // 大根堆
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	if this.arrMax.Len() == 0 || this.arrMin.Len() == 0 {
		heap.Push(&this.arrMax, -num)
		if this.arrMax.Len() > this.arrMin.Len()+1 {
			heap.Push(&this.arrMin, -heap.Pop(&this.arrMax).(int))
		}
	} else {
		if -this.arrMax.IntSlice[0] < num {
			heap.Push(&this.arrMin, num)
			fmt.Println(this.arrMin, this.arrMax)
			if this.arrMin.Len() > this.arrMax.Len()+1 {
				heap.Push(&this.arrMax, -heap.Pop(&this.arrMin).(int))
			}
		} else {
			heap.Push(&this.arrMax, -num)
			fmt.Println(this.arrMin, this.arrMax)
			if this.arrMax.Len() > this.arrMin.Len()+1 {
				heap.Push(&this.arrMin, -heap.Pop(&this.arrMax).(int))
			}
		}
	}
	fmt.Println(this.arrMin, this.arrMax)
}

func (this *MedianFinder) FindMedian() float64 {
	if (this.arrMax.Len()+this.arrMin.Len())%2 == 1 {
		if this.arrMax.Len() > this.arrMin.Len() {
			return float64(-this.arrMax.IntSlice[0])
		} else {
			return float64(this.arrMin.IntSlice[0])
		}
	} else {
		return float64(-this.arrMax.IntSlice[0]+this.arrMin.IntSlice[0]) / 2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func main() {
	m := Constructor()
	m.AddNum(-1)
	m.AddNum(-2)
	m.AddNum(-3)
	m.AddNum(-4)
	m.AddNum(-6)
	m.AddNum(-5)
	fmt.Println(m.FindMedian())
}
