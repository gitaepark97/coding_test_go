package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Question:
// N개의 숫자가 공백 없이 쓰여있다. 이 숫자를 모두 합해서 출력하는 프로그램을 작성하시오.

// Input:
// 첫째 줄에 숫자의 개수 N (1 ≤ N ≤ 100)이 주어진다. 둘째 줄에 숫자 N개가 공백없이 주어진다.

// Output:
// 입력으로 주어진 숫자 N개의 합을 출력한다.

func main() {
	var r = bufio.NewReader(os.Stdin)
	var w = bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var count int
	var nums string
	var sum int

	fmt.Fscan(r, &count)
	fmt.Fscan(r, &nums)

	arrNums := strings.Split(nums, "")

	for _, num := range arrNums {
		n, _ := strconv.Atoi(num)
		sum += n
	}

	fmt.Fprint(w, sum)
}

