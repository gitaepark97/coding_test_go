package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// Question:
// 절댓값 힙은 다음과 같은 연산을 지원하는 자료구조이다.
// 1. 배열에 정수 x (x ≠ 0)를 넣는다.
// 2. 배열에서 절댓값이 가장 작은 값을 출력하고, 그 값을 배열에서 제거한다. 절댓값이 가장 작은 값이 여러개일 때는, 가장 작은 수를 출력하고, 그 값을 배열에서 제거한다.
// 프로그램은 처음에 비어있는 배열에서 시작하게 된다.

// Input:
// 첫째 줄에 연산의 개수 N(1≤N≤100,000)이 주어진다. 다음 N개의 줄에는 연산에 대한 정보를 나타내는 정수 x가 주어진다. 만약 x가 0이 아니라면 배열에 x라는 값을 넣는(추가하는) 연산이고, x가 0이라면 배열에서 절댓값이 가장 작은 값을 출력하고 그 값을 배열에서 제거하는 경우이다. 입력되는 정수는 -231보다 크고, 231보다 작다.

// Output:
// 입력에서 0이 주어진 회수만큼 답을 출력한다. 만약 배열이 비어 있는 경우인데 절댓값이 가장 작은 값을 출력하라고 한 경우에는 0을 출력하면 된다.

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
	var r = bufio.NewReader(os.Stdin)
	var w = bufio.NewWriter(os.Stdout)
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