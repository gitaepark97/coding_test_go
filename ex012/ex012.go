package main

import (
	"bufio"
	"fmt"
	"os"
)

// Question:
// 크기가 N인 수열 A = A1, A2, ..., AN이 있다. 수열의 각 원소 Ai에 대해서 오큰수 NGE(i)를 구하려고 한다. Ai의 오큰수는 오른쪽에 있으면서 Ai보다 큰 수 중에서 가장 왼쪽에 있는 수를 의미한다. 그러한 수가 없는 경우에 오큰수는 -1이다.
// 예를 들어, A = [3, 5, 2, 7]인 경우 NGE(1) = 5, NGE(2) = 7, NGE(3) = 7, NGE(4) = -1이다. A = [9, 5, 4, 8]인 경우에는 NGE(1) = -1, NGE(2) = 8, NGE(3) = 8, NGE(4) = -1이다.

// Input:
// 첫째 줄에 수열 A의 크기 N (1 ≤ N ≤ 1,000,000)이 주어진다. 둘째 줄에 수열 A의 원소 A1, A2, ..., AN (1 ≤ Ai ≤ 1,000,000)이 주어진다.

// Output:
// 총 N개의 수 NGE(1), NGE(2), ..., NGE(N)을 공백으로 구분해 출력한다.

func main() {
	var r = bufio.NewReader(os.Stdin)
	var w = bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int

	fmt.Fscan(r, &n)

	nums := make([]int, n)
	result := make([]int, n)
	stack := make([]int, 0, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &nums[i])
	}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack) - 1]] {
			lastIdx := len(stack) - 1
			result[stack[lastIdx]] = nums[i]
			stack = stack[:lastIdx]
		}
		stack = append(stack, i)
	}

	for len(stack) > 0 {
		lastIdx := len(stack) - 1
		result[stack[lastIdx]] = -1
		stack = stack[:lastIdx]
	}

	for _, val := range result {
		fmt.Fprintf(w, "%d ", val)
	}
}