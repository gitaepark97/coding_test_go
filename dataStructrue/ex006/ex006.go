package main

import (
	"bufio"
	"fmt"
	"os"
)

// Question:
// 어떠한 자연수 N은, 몇 개의 연속된 자연수의 합으로 나타낼 수 있다. 당신은 어떤 자연수 N(1 ≤ N ≤ 10,000,000)에 대해서, 이 N을 몇 개의 연속된 자연수의 합으로 나타내는 가지수를 알고 싶어한다. 이때, 사용하는 자연수는 N이하여야 한다.
// 예를 들어, 15를 나타내는 방법은 15, 7+8, 4+5+6, 1+2+3+4+5의 4가지가 있다. 반면에 10을 나타내는 방법은 10, 1+2+3+4의 2가지가 있다.
// N을 입력받아 가지수를 출력하는 프로그램을 작성하시오.

// Input:
// 첫 줄에 정수 N이 주어진다.

// Output:
// 입력된 자연수 N을 몇 개의 연속된 자연수의 합으로 나타내는 가지수를 출력하시오

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	var n int

	fmt.Fscan(r, &n)
	
	count := 1
	startIdx := 1
	endIdx := 1
	sum := 1

	for endIdx != n {
		if sum == n {
			count++
			endIdx++
			sum += endIdx
		} else if sum > n {
			sum -= startIdx
			startIdx++
		} else {
			endIdx++
			sum += endIdx
		}
	}

	fmt.Fprint(w, count)
}