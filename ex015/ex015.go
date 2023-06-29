package main

import (
	"bufio"
	"fmt"
	"os"
)

// Question:
// N개의 수가 주어졌을 때, 이를 오름차순으로 정렬하는 프로그램을 작성하시오.

// Input:
// 첫째 줄에 수의 개수 N(1 ≤ N ≤ 1,000)이 주어진다. 둘째 줄부터 N개의 줄에는 수가 주어진다. 이 수는 절댓값이 1,000보다 작거나 같은 정수이다. 수는 중복되지 않는다.

// Output:
// 첫째 줄부터 N개의 줄에 오름차순으로 정렬한 결과를 한 줄에 하나씩 출력한다.

func main() {
	var r = bufio.NewReader(os.Stdin)
	var w = bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int

	fmt.Fscan(r, &n)

	datas := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &datas[i])
	}

	for i := 0; i < n - 1; i++ {
		for j := 0; j < n - 1 - i; j++ {
			if datas[j] > datas[j + 1] {
				tmp := datas[j]
				datas[j] = datas[j + 1]
				datas[j + 1] = tmp
			}
		}
	}

	for _, data := range datas {
		fmt.Fprintln(w, data)
	}
}