package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var r = bufio.NewReader(os.Stdin)
	var w = bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m int
	var sumDatas [1000001]int
	var remains [1000]int
	var answer int

	fmt.Fscan(r, &n, &m)

	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &sumDatas[i])

		sumDatas[i] += sumDatas[i - 1]

		remain := sumDatas[i] % m
		if remain == 0 {
			answer++
		}

		remains[remain]++
	}

	for i := 0; i < m; i++ {
		if remains[i] > 1 {
			answer = answer + (remains[i] * (remains[i] - 1) / 2)
		}
	}

	fmt.Fprint(w, answer)
}