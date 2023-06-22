package main

import (
	"bufio"
	"fmt"
	"os"
)

// Question:
// N×N개의 수가 N×N 크기의 표에 채워져 있다. (x1, y1)부터 (x2, y2)까지 합을 구하는 프로그램을 작성하시오. (x, y)는 x행 y열을 의미한다.
// 표에 채워져 있는 수와 합을 구하는 연산이 주어졌을 때, 이를 처리하는 프로그램을 작성하시오.

// Input:
// 첫째 줄에 표의 크기 N과 합을 구해야 하는 횟수 M이 주어진다. (1 ≤ N ≤ 1024, 1 ≤ M ≤ 100,000) 둘째 줄부터 N개의 줄에는 표에 채워져 있는 수가 1행부터 차례대로 주어진다. 다음 M개의 줄에는 네 개의 정수 x1, y1, x2, y2 가 주어지며, (x1, y1)부터 (x2, y2)의 합을 구해 출력해야 한다. 표에 채워져 있는 수는 1,000보다 작거나 같은 자연수이다. (x1 ≤ x2, y1 ≤ y2)

// Output:
// 총 M줄에 걸쳐 (x1, y1)부터 (x2, y2)까지 합을 구해 출력한다.

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	var n, m int
	var startX, startY, endX, endY int
	var sumDatas [1025][1025]int
	
	fmt.Fscan(r, &n, &m)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Fscan(r, &sumDatas[i][j])

			sumDatas[i][j] += sumDatas[i - 1][j] + sumDatas[i][j - 1] - sumDatas[i - 1][j - 1]
		}
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(r, &startX, &startY, &endX, &endY)

		fmt.Fprintln(w, sumDatas[endX][endY] - sumDatas[startX - 1][endY] - sumDatas[endX][startY - 1] + sumDatas[startX - 1][startY - 1])
	}
}