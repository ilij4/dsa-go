package leetcode

import "github.com/ilij4/go-demo/helpers"

type Element struct {
	Val int
	Min int
}

// 155. Min Stack
type MinStack struct {
	items []*Element
}

func Constructor() MinStack {
	return MinStack{
		items: make([]*Element, 0),
	}
}

func (this *MinStack) Push(val int) {
	min := val

	if !this.IsEmpty() {
		min = helpers.Min(this.items[len(this.items)-1].Min, val)
	}

	this.items = append(this.items, &Element{Val: val, Min: min})
}

func (this *MinStack) Pop() {
	n := len(this.items)
	if n == 0 {
		return
	}

	this.items[n-1] = nil
	this.items = this.items[:n-1]
}

func (this *MinStack) Top() int {
	if this.IsEmpty() {
		return 0
	}

	return this.items[len(this.items)-1].Val
}

func (this *MinStack) GetMin() int {
	if this.IsEmpty() {
		return 0
	}

	return this.items[len(this.items)-1].Min
}

func (this *MinStack) IsEmpty() bool {
	return len(this.items) == 0
}

func (this *MinStack) Peek() (int, bool) {
	if this.IsEmpty() {
		return 0, false
	}

	return this.items[len(this.items)-1].Val, true
}
