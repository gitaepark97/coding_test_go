package main

import (
	"bufio"
	"fmt"
	"os"
)

// Question:
// N개의 수 A1, A2, ..., AN과 L이 주어진다.
// Di = Ai-L+1 ~ Ai 중의 최솟값이라고 할 때, D에 저장된 수를 출력하는 프로그램을 작성하시오. 이때, i ≤ 0 인 Ai는 무시하고 D를 구해야 한다.

// Input:
// 첫째 줄에 N과 L이 주어진다. (1 ≤ L ≤ N ≤ 5,000,000)
// 둘째 줄에는 N개의 수 Ai가 주어진다. (-109 ≤ Ai ≤ 109)

// Output:
// 첫째 줄에 Di를 공백으로 구분하여 순서대로 출력한다.

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

type Node struct {
	idx, val int
}

type Deque []Node

func (d Deque) isEmpty() bool {
	return len(d) == 0
}

func (d *Deque) push(node Node) {
	*d = append(*d, node)
}

func (d Deque) getFirst() Node {
	return d[0]
}

func (d *Deque) popFirst() {
	*d = (*d)[1:]
}

func (d Deque) getLast() Node {
	return d[len(d)-1]
}

func (d *Deque) popLast() {
	*d = (*d)[:len(*d)-1]
}

func main() {
	defer w.Flush()

	var n, l int
	var d Deque

	fmt.Fscan(r, &n, &l)

	for i := 0; i < n; i++ {
		var now int

		fmt.Fscan(r, &now)

		if !d.isEmpty() && d.getFirst().idx <= i-l {
			d.popFirst()
		}

		for !d.isEmpty() && d.getLast().val > now {
			d.popLast()
		}

		d.push(Node{i, now})

		fmt.Fprintf(w, "%d ", d.getFirst().val)
	}
}
