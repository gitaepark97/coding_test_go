package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

type Heap struct {
	datas   []int
	length int
}

func (heap *Heap) insertHeap(input int) {
	heap.length++
	var idx = heap.length
	absInput := math.Abs(float64(input))

	for idx != 1 {
		mark := heap.datas[idx / 2]
		absMark := math.Abs(float64(mark)) 

		if absInput < absMark || absInput == absMark && input <= mark {
			heap.datas[idx] = mark
			idx = idx / 2
		} else {
			break
		}
	}
	heap.datas[idx] = input
}

func (heap *Heap) getPriorityChildIdx(parentIdx int) int {
	if parentIdx * 2 > heap.length {
		return -1
	} else if parentIdx * 2 == heap.length {
		return parentIdx * 2
	} else {
		mark1 := heap.datas[parentIdx * 2]
		mark2 := heap.datas[parentIdx * 2 + 1]
		absMark1 := math.Abs(float64(mark1))
		absMark2 := math.Abs(float64(mark2))

		if absMark1 > absMark2 {
			return parentIdx * 2 + 1
		} else if absMark1 == absMark2 {
			if mark1 >= mark2 {
				return parentIdx * 2 + 1
			}
		}
		return parentIdx * 2
	}
}

func (heap *Heap) deleteHeap() int {
	if heap.length == 0 {
		return 0
	}

	delVal := heap.datas[1]
	lastVal := heap.datas[heap.length]
	absLastVal := math.Abs(float64(lastVal))

	var parentIdx = 1
	var childIdx = heap.getPriorityChildIdx(parentIdx)
	for childIdx != -1 {
		mark := heap.datas[childIdx]
		absMark := math.Abs(float64(mark))

		if absLastVal < absMark || absLastVal == absMark && lastVal <= mark {
			break
		}
		heap.datas[parentIdx] = mark
		parentIdx = childIdx
		childIdx = heap.getPriorityChildIdx(parentIdx)
	}

	heap.datas[parentIdx] = lastVal
	heap.length--

	return delVal
}

func main() {
	defer w.Flush()

	var n int

	fmt.Fscan(r, &n)

	var heap Heap
	heap.datas = make([]int, n + 1)

	for i := 0; i < n; i++ {
		var input int

		fmt.Fscan(r, &input)

		if input == 0 {
			fmt.Fprintln(w, heap.deleteHeap())
		} else {
			heap.insertHeap(input)
		}
	}
}