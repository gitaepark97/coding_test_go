package main

import (
	"bufio"
	"fmt"
	"os"
)

// Question:
// 수 N개가 주어졌을 때, i번째 수부터 j번째 수까지 합을 구하는 프로그램을 작성하시오.

// Input:
// 첫째 줄에 수의 개수 N과 합을 구해야 하는 횟수 M이 주어진다. 둘째 줄에는 N개의 수가 주어진다. 수는 1,000보다 작거나 같은 자연수이다. 셋째 줄부터 M개의 줄에는 합을 구해야 하는 구간 i와 j가 주어진다.

// Output:
// 총 M개의 줄에 입력으로 주어진 i번째 수부터 j번째 수까지 합을 출력한다.

func main() {
	var r = bufio.NewReader(os.Stdin)
	var w = bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m int
	var start, end int
	var sumDatas [1025]int

	fmt.Fscan(r, &n, &m)

	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &sumDatas[i])

		sumDatas[i] += sumDatas[i - 1]
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(r, &start, &end)

		fmt.Fprintln(w, sumDatas[end] - sumDatas[start - 1])
	}
}